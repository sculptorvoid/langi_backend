package entity

type Word struct {
	Id          int    `json:"id" db:"id"`
	Word        string `json:"word" db:"word" binding:"required"`
	Translation string `json:"translation" db:"translation"`
}
