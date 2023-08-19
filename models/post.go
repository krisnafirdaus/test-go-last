package models

type Post struct {
	Id      int    `json:"id" gorm:"primaryKey"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
