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

func (r *DictionaryPsql) GetAllDictionaries(userId int) ([]entity.Dictionary, error) {
	var dictionaries []entity.Dictionary
	query := fmt.Sprintf(`SELECT dicts.id, dicts.title FROM %s dicts 
					      INNER JOIN %s userDicts ON dicts.id = userDicts.dict_id 
						  WHERE userDicts.user_id = $1`, dictionariesTable, usersDictsTable)
	err := r.db.Select(&dictionaries, query, userId)

	return dictionaries, err
}

func (r *DictionaryPsql) GetById(userId, dictId int) (entity.Dictionary, error) {
	var dictionary entity.Dictionary
	query := fmt.Sprintf(`SELECT dicts.id, dicts.title FROM %s dicts 
 						  INNER JOIN %s userDicts ON dicts.id = userDicts.dict_id 
						  WHERE userDicts.user_id = $1 
						  AND userDicts.dict_id = $2`, dictionariesTable, usersDictsTable)
	err := r.db.Get(&dictionary, query, userId, dictId)

	return dictionary, err
}

func (r *DictionaryPsql) Delete(userId, dictId int) error {
	query := fmt.Sprintf(`DELETE FROM %s dicts USING %s userDicts 
					      WHERE dicts.id = userDicts.dict_id 
					      AND userDicts.user_id = $1 
						  AND userDicts.dict_id = $2`, dictionariesTable, usersDictsTable)
	_, err := r.db.Exec(query, userId, dictId)

	return err
}

func (r *DictionaryPsql) Update(userId, dictId int, input entity.UpdateDictionaryInput) error {
	query := fmt.Sprintf(`UPDATE %s dicts
 						  SET title = $1
						  FROM %s userDicts 
						  WHERE dicts.id = userDicts.dict_id 
						  AND userDicts.dict_id=$2 
						  AND userDicts.user_id=$3`, dictionariesTable, usersDictsTable)

	_, err := r.db.Exec(query, input.Title, dictId, userId)

	return err
}
