package repository

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

const (
	usersTable        = "users"
	dictionariesTable = "dictionaries"
	usersDictsTable   = "users_dictionaries"
	wordsTable        = "words"
	dictsWordsTable   = "dictionaries_words"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func PostgresConnect(cfg Config) (*sqlx.DB, error) {
	var db *sqlx.DB
	var connectCount time.Duration
	var err error
	connectionsLimit := 10
	for {
		connectCount++
		if int(connectCount) >= connectionsLimit {
			logrus.Fatalf("\n\nLangi_backend: Cannot connect to Postgres, Connections limit exceeded! Exit.\n\n")
			break
		}
		db, err = sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))

		if err != nil {
			return nil, err
		}
		err = db.Ping()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Langi_backend: Cannot connect to Postgres, try again...(%d) Error: %s", connectCount, err.Error())
			time.Sleep(connectCount * time.Second)
			continue
		} else {
			break
		}
	}

	return db, nil
}

func MakeMigrations(db *sqlx.DB) error {
	sqlDB := db.DB
	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance("file:///go/migrations", "postgres", driver)
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil {
		return err
	}

	return nil
}
