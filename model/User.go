package model

import (
	"encoding/base64"
	"fmt"
	"ginblog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
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
	Db.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ErrorUsernameUsed //1001
	}
	return errmsg.SUCCESS //200

}

//新增用户
//结构体为引用类型，在函数中作为入参传他的指针
func CreateUser(data *User) int {
	//data.Password = ScryptPw(data.Password)
	err := Db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}

//查询用户列表
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	fmt.Printf("User46:%v\n", users)
	fmt.Printf("User47:%v\n", pageNum)
	fmt.Printf("User48:%v\n", pageSize)
	fmt.Printf("User47:%v\n", Db.Select(&pageNum))
	fmt.Printf("User48:%v\n", Db.Select(&pageSize))
	err = Db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&users).Error
	fmt.Printf("User50:%v\n", users)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	fmt.Printf("User54:%v\n", users)
	return users
}

//编辑用户
func EditUser() {

}

//删除

func (u *User) BeforeSave() {
	u.Password = ScryptPw(u.Password)
}

//密码加密
func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{123, 13, 34, 44, 53, 12, 71, 81}
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}
