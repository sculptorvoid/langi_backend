package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sculptorvoid/langi_backend/pkg/entity"
	"github.com/sirupsen/logrus"
	"strings"
)

type WordPsql struct {
	db *sqlx.DB
}

func NewWordPsql(db *sqlx.DB) *WordPsql {
	return &WordPsql{db: db}
}

func (r *WordPsql) Create(dictId int, word entity.Word) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var wordId int
	createWordQuery := fmt.Sprintf("INSERT INTO %s (word, translation) VALUES ($1, $2) RETURNING id", wordsTable)
	row := tx.QueryRow(createWordQuery, word.Word, word.Translation)
	err = row.Scan(&wordId)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			logrus.Fatalf("cannot rollback transaction createWordQuery: %s", err.Error())
		}
		return 0, err
	}

	createDictsWordsQuery := fmt.Sprintf("INSERT INTO %s (dict_id, word_id) VALUES ($1, $2)", dictsWordsTable)
	_, err = tx.Exec(createDictsWordsQuery, dictId, wordId)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			logrus.Fatalf("cannot rollback transaction createDictsWordsQuery: %s", err.Error())
		}
		return 0, err
	}

	return wordId, tx.Commit()
}

func (r *WordPsql) GetAll(userId, dictId int) ([]entity.Word, error) {
	var words []entity.Word
	query := fmt.Sprintf(`SELECT word.id, word.word, word.translation 
							FROM %s word 
							INNER JOIN %s dictsWords ON dictsWords.word_id = word.id
							INNER JOIN %s usersDicts ON usersDicts.dict_id = dictsWords.dict_id 
							WHERE dictsWords.dict_id = $1 AND usersDicts.user_id = $2`,
		wordsTable, dictsWordsTable, usersDictsTable)

	if err := r.db.Select(&words, query, dictId, userId); err != nil {
		return nil, err
	}

	return words, nil
}

func (r *WordPsql) GetById(userId, wordId int) (entity.Word, error) {
	var word entity.Word
	query := fmt.Sprintf(`SELECT word.id, word.word, word.translation 
							FROM %s word 
							INNER JOIN %s dictsWords ON dictsWords.word_id = word.id
							INNER JOIN %s usersDicts ON usersDicts.dict_id = dictsWords.dict_id 
							WHERE word.id = $1 AND usersDicts.user_id = $2`,
		wordsTable, dictsWordsTable, usersDictsTable)

	if err := r.db.Get(&word, query, wordId, userId); err != nil {
		return word, err
	}

	return word, nil
}

func (r *WordPsql) Delete(userId, wordId int) error {
	query := fmt.Sprintf(`DELETE FROM %s word USING %s dictsWords, %s usersDicts 
									WHERE word.id = dictsWords.word_id AND dictsWords.dict_id = usersDicts.dict_id AND usersDicts.user_id = $1 AND word.id = $2`,
		wordsTable, dictsWordsTable, usersDictsTable)
	_, err := r.db.Exec(query, userId, wordId)
	return err
}

func (r *WordPsql) Update(userId, wordId int, input entity.UpdateWordInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Word != nil {
		setValues = append(setValues, fmt.Sprintf("word=$%d", argId))
		args = append(args, *input.Word)
		argId++
	}

	if input.Translation != nil {
		setValues = append(setValues, fmt.Sprintf("translation=$%d", argId))
		args = append(args, *input.Translation)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s word SET %s FROM %s dictsWords, %s usersDicts
									WHERE word.id = dictsWords.word_id AND dictsWords.dict_id = usersDicts.dict_id AND usersDicts.user_id = $%d AND word.id = $%d`,
		wordsTable, setQuery, dictsWordsTable, usersDictsTable, argId, argId+1)
	args = append(args, userId, wordId)

	_, err := r.db.Exec(query, args...)
	return err

}
