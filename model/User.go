package model

import (
	"fmt"
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

//数据库操作的方法

//查询用户是否存在
func CheckUser(name string) (code int) {
	var users User
	db.Select("id").Where("username = ?", name).First(&users)
	fmt.Println(users)
	if users.ID > 0 {
		return errmsg.ErrorUsernameUsed //1001
	}
	return errmsg.SUCCESS //200
}

//新增用户
//结构体为引用类型，在函数中作为入参传他的指针
func CreateUser(data *User) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}
