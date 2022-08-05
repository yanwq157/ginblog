package v1

import (
	"fmt"
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var code int

//查询用户是否存在
func UserExist(c *gin.Context) {

}

//添加用户
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}
	if code == errmsg.ErrorUsernameUsed {
		code = errmsg.ErrorUsernameUsed
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询单个用户
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	fmt.Printf("user40:%v\n", pageSize)
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	fmt.Printf("user42:%v\n", pageNum)
	if pageNum == 0 {
		pageNum = -1
	}
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	fmt.Printf("user52:%v\n", pageNum)

	data := model.GetUsers(pageSize, pageNum)
	fmt.Printf("user55:%v\n", data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询用户列表

//编辑用户
func EditUser(c *gin.Context) {

}

//删除用户
func DeleteUser(c *gin.Context) {

}
