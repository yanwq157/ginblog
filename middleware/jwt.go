package middleware

import (
	"fmt"
	"ginblog/utils"
	"ginblog/utils/errmsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var JetKey = []byte(utils.JwtKey) //声明用于加密解密的秘钥

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var code int

//生成token

func SetToken(username string) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour)
	//自定义结构体配置加密的参数
	SetClaims := MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //到期时间声明
			Issuer:    "gin-blog",        //发行人声明
		},
	}
	//使用NewWithClaims new一个token，两个参数，返回一个token
	//参数SigningMethodES256为加密的方法
	//参数SetClaims为加密的一些参数，可以使用自带的MapClaims,也可以自定义结构体，
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaim.SignedString(JetKey) //获取完整的签名令牌
	fmt.Printf("token:%v\n", token)
	if err != nil {
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCESS
}

//验证token

func CheckToken(token string) (*MyClaims, int) {
	//自定义结构体使用ParseWithClaims方法解密，有三个参数：
	//1:加密后的token字符串
	//2:加密使用的模版，当前使用的是MyClaims
	//3:一个自带的回调函数，将秘钥和错误return出来
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JetKey, nil
	})
	//token结构体 ，令牌的第二段，setToken.Claims.(*MyClaims)
	//Valid 验证令牌是否有效，解析/验证令牌时填充
	if key, _ := setToken.Claims.(*MyClaims); setToken.Valid {
		return key, errmsg.SUCCESS
	} else {
		return nil, errmsg.ERROR
	}
}

//jwt中间件
//HandlerFunc 将gin中间件使用的处理程序定义为返回值。

func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取值，如果没有关联的值，返回""，不区分大小写
		tokenHarder := c.Request.Header.Get("Authorization")
		if tokenHarder == "" {
			code = errmsg.ErrorTokenExist
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return

		}
		//按空格切割为分片，每个两个字符
		checkToken := strings.SplitN(tokenHarder, " ", 2)
		fmt.Println(checkToken)
		if len(checkToken) != 2 && checkToken[8] != "Bearer" {
			code = errmsg.ErrorTokenTypeWrong
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		key, tCode := CheckToken(checkToken[1])
		if tCode == errmsg.ERROR {
			code = errmsg.ErrorTokenWrong
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ErrorTokenRuntime
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		c.Set("username", key.Username)
		c.Next()
	}
}
