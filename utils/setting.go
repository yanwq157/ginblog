package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var AppMode string
var HttpPort string
var JwtKey string

//Db         string

var DbHost string
var DbPort string
var DbUser string
var DbPassword string
var DbName string

var AccessKey string
var SecretKey string
var Bucket string
var QiniuServer string

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("config reload err", err)
	}
	LoadServer(file)
	LoadDataBase(file)
	LoadQiniu(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":8080")
	JwtKey = file.Section("server").Key("JwtKey").MustString("test123!@#$")

}

func LoadDataBase(file *ini.File) {
	//Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassword = file.Section("database").Key("DbPassword").MustString("123123")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}

func LoadQiniu(file *ini.File) {

	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SecretKey = file.Section("qiniu").Key("SecretKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	QiniuServer = file.Section("qiniu").Key("QiniuServer").String()
}
