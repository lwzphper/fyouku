package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/beego/beego/v2/core/config"
	beego "github.com/beego/beego/v2/server/web"
	"time"
)

type CommonController struct {
	beego.Controller
}

type JsonStruct struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Items interface{} `json:"items"`
	Count int64       `json:"count"`
}

func ReturnSuccess(code int, msg interface{}, items interface{}, count int64) (json *JsonStruct) {
	json = &JsonStruct{Code: code, Msg: msg, Items: items, Count: count}
	return
}

func ReturnError(code int, msg interface{}) (json *JsonStruct) {
	json = &JsonStruct{Code: code, Msg: msg}
	return
}

func MD5V(password string) string {
	md5code, _ := config.String("md5code")
	h := md5.New()
	h.Write([]byte(password + md5code))
	return hex.EncodeToString(h.Sum(nil))
}

func DateFormat(times int64) string {
	videoTime := time.Unix(times, 0)
	return videoTime.Format("2006-01-02")
}