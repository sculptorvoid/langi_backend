package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sculptorvoid/langi_backend/pkg/entity"
	"github.com/sirupsen/logrus"
)

type DictionaryPsql struct {
	db *sqlx.DB
}

func NewDictionaryPsql(db *sqlx.DB) *DictionaryPsql {
	return &DictionaryPsql{db: db}
}

func (r *DictionaryPsql) CreateDictionary(userId int, dictionary entity.Dictionary) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createDictionaryQuery := fmt.Sprintf("INSERT INTO %s (title) VALUES ($1) RETURNING id", dictionariesTable)
	row := tx.QueryRow(createDictionaryQuery, dictionary.Title)
	if err := row.Scan(&id); err != nil {
		if err := tx.Rollback(); err != nil {
			logrus.Fatalf("cannot rollback transaction createDictionaryQuery: %s", err.Error())
		}
		return 0, err
	}

	createUsersDictionariesQuery := fmt.Sprintf("INSERT INTO %s (user_id, dict_id) VALUES ($1, $2)", usersDictsTable)
	_, err = tx.Exec(createUsersDictionariesQuery, userId, id)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			logrus.Fatalf("cannot rollback transaction createUsersDictionariesQuery: %s", err.Error())
		}
		return 0, err
	}

	return id, tx.Commit()
}
