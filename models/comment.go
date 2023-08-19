package models

type Comment struct {
	Id     int    `json:"id" gorm:"primaryKey"`
	Text   string `json:"text"`
	PostId int    `gorm:"index"` // Foreign key untuk relasi dengan Post
}
