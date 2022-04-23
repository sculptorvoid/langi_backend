package entity

type Dictionary struct {
	Id    int    `json:"id" db:"id"`
	Title string `json:"title" db:"title" binding:"required"`
}

type DictionariesWords struct {
	Id     int
	DictId int
	WordId int
}
