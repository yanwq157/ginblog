package model

import (
	"encoding/base64"
	"ginblog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=14" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=4,max=14" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色"` //validate里面为0是空，改为1为管理员，2为用户,默认为2
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

func GetUsers(pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var total int64
	err = Db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&users).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, total
}

//编辑用户

func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{}) //value有字符串，有整数，所以用interface
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := Db.Model(&user).Where("id = ?", id).Updates(maps).Error //传User 在Model模型更新
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除用户

func DeleteUser(id int) int {
	var user User
	err := Db.Where("id= ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//密码加密

func (u *User) BeforeSave(*gorm.DB) (err error) {
	u.Password = ScryptPw(u.Password)
	return
}

func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 13, 34, 44, 53, 12, 71, 81}
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}

//登录验证

func CheckLogin(username string, password string) int {
	var user User
	Db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ErrorUserNotExist
	}
	if ScryptPw(password) != user.Password {
		return errmsg.ErrorPasswordWrong
	}
	if user.Role != 1 {
		return errmsg.ErrorUserNoRight
	}
	return errmsg.SUCCESS
}
