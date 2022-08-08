package model

import (
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

//查询分类是否存在
func CheckCategory(name string) (code int) {
	var cate Category
	Db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ErrorCateNameUsed //2001
	}
	return errmsg.SUCCESS //200

}

//新增分类

func CreateCate(data *Category) int {
	err := Db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}

//查询分类列表

func GetCate(pageSize int, pageNum int) []Category {
	var cate []Category
	err = Db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&cate).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cate
}

//编辑分类

func EditCate(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{}) //value有字符串，有整数，所以用interface
	maps["id"] = data.ID
	maps["name"] = data.Name
	err := Db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除分类

func DeleteCate(id int) int {
	var cate Category
	err := Db.Where("id= ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
