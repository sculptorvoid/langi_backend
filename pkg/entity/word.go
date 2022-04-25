package entity

import "errors"

type Word struct {
	Id          int    `json:"id" db:"id"`
	Word        string `json:"word" db:"word" binding:"required"`
	Translation string `json:"translation" db:"translation"`
}

type UpdateWordInput struct {
	Word        *string `json:"word"`
	Translation *string `json:"translation"`
}

func (i UpdateWordInput) Validate() error {
	if i.Word == nil || i.Translation == nil {
		return errors.New("update word failed: no such values")
	}

	return nil
}
