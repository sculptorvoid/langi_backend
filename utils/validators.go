package utils

import "errors"

type UpdateDictionaryInput struct {
	Title *string `json:"title"`
}

func (i UpdateDictionaryInput) Validate() error {
	if i.Title == nil {
		return errors.New("update dictionary failed: no such values")
	}

	return nil
}

type UpdateWordInput struct {
	Word        *string `json:"word"`
	Translation *bool   `json:"done"`
}

func (i UpdateWordInput) Validate() error {
	if i.Word == nil && i.Translation == nil {
		return errors.New("update word failed: no such values")
	}

	return nil
}
