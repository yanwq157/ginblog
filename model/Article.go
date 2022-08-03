package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Category Category `gorm:"foreignkey:Cid"`
	Title    string   `gorm:"type:varchar(100);not null" Json:"title"`
	Cid      int      `gorm:"type:int;not null" Json:"cid"`
	Desc     string   `gorm:"type:varchar(200);not null" Json:"desc"`
	Content  string   `gorm:"type:longtext;not null" Json:"content"`
	Img      string   `gorm:"type:varchar(100);not null" Json:"img"`
}
