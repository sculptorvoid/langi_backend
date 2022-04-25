package entity

import "errors"

type Dictionary struct {
	Id    int    `json:"id" db:"id"`
	Title string `json:"title" db:"title" binding:"required"`
}

type DictionariesWords struct {
	Id     int
	DictId int
	WordId int
}

type UpdateDictionaryInput struct {
	Title *string `json:"title"`
}

func (i UpdateDictionaryInput) Validate() error {
	if i.Title == nil {
		return errors.New("update dictionary failed: no such values")
	}

	return nil
}
