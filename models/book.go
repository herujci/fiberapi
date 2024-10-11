package models

type Book struct {
	Id          int64  `gorm:"primary_key;auto_increment" json:"id"`
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	Author      string `gorm:"type:varchar(255)" json:"author"`
	PublishDate string `gorm:"type:date" json:"publish_date"`
}
