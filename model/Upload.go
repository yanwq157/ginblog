package model

import (
	"context"
	"ginblog/utils"
	"ginblog/utils/errmsg"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
)

var AccessKey = utils.AccessKey
var SecretKey = utils.SecretKey
var Bucket = utils.Bucket
var ImgUrl = utils.QiniuServer

func UploadFile(file multipart.File, fileSize int64) (string, int) {
	//上传的凭证
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)
	//构建一个上传用的Config对象
	cfg := storage.Config{
		Zone:          &storage.ZoneHuabei,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	//表单上传的额外可选项
	putExtra := storage.PutExtra{}
	//构建一个表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	//PutRet 为七牛标准的上传回复内容。
	ret := storage.PutRet{}
	//用来以表单方式上传一个文件
	//ctx 是请求的上下文.ret 是上传成功后返回的数据。
	//UpToken 是由业务服务器颁发的上传凭证。
	//data 是文件内容的访问接口（io.Reader）。
	//FSize 是要上传的文件大小。
	//extra 是上传的一些可选项。
	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return "", errmsg.ERROR
	}
	url := ImgUrl + ret.Key
	return url, errmsg.SUCCESS
}
