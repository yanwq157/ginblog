package model

import (
	"fmt"
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Category Category `gorm:"foreignKey:Cid"`
	Title    string   `gorm:"type:varchar(100);not null" Json:"title"`
	Cid      int      `gorm:"type:int;not null" Json:"cid"`
	Desc     string   `gorm:"type:varchar(200);not null" Json:"desc"`
	Content  string   `gorm:"type:longtext;not null" Json:"content"`
	Img      string   `gorm:"type:varchar(100);not null" Json:"img"`
}

//新增文章

func CreateArt(data *Article) int {
	err := Db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}

//查询分类下所有文章
func GatCateArt(id int, pageSize int, pageMum int) ([]Article, int) {
	var crateArtist []Article
	err := Db.Preload("Category").Offset((pageMum-1)*pageSize).Limit(pageSize).Where("cid = ?", id).Find(&crateArtist).Error
	if err != nil {
		return nil, errmsg.ErrorArtNotExits
	}
	return crateArtist, errmsg.SUCCESS
}

//查询单个文章

func GetArtInfo(id int) (Article, int) {
	var art Article
	err := Db.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ErrorArtNotExits
	}
	return art, errmsg.SUCCESS
}

//查询文章列表

func GetArt(pageSize int, pageNum int) ([]Article, int) {
	var articlelist []Article
	err = Db.Preload("Category").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&articlelist).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	return articlelist, errmsg.SUCCESS
}

//编辑分类

func EditArt(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{}) //value有字符串，有整数，所以用interface
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err := Db.Model(&art).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除分类

func DeleteArt(id int) int {
	var art Article
	err := Db.Where("id= ?", id).Delete(&art).Error
	fmt.Println(err)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
