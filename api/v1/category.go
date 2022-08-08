package v1

import (
	"fmt"
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//添加分类

func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		model.CreateCate(&data)
	}
	if code == errmsg.ErrorCateNameUsed {
		code = errmsg.ErrorCateNameUsed
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询分类列表

func GetCate(c *gin.Context) {
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

	data := model.GetCate(pageSize, pageNum)
	fmt.Printf("user55:%v\n", data)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
	})
}

//编辑分类名

func EditCate(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		model.EditCate(id, &data)
	}
	if code == errmsg.ErrorCateNameUsed {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//删除分类

func DeleteCate(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteCate(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询单个分类下的文章
