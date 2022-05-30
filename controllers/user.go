package controllers

import (
	"fyouku/models"
	beego "github.com/beego/beego/v2/server/web"
	"regexp"
)

type UserController struct {
	beego.Controller
}

// SaveRegister 用户注册
// @router /register/save [post]
func (u *UserController) SaveRegister() {
	var (
		mobile string
		password string
		err error
	)
	mobile = u.GetString("mobile")
	password = u.GetString("password")

	checkJson := checkMobileAddPassword(mobile, password)
	if checkJson != nil {
		u.Data["json"] = checkJson
		err = u.ServeJSON()
		if err != nil {
			return
		}
	}

	if models.ExistUserMobile(mobile) {
		u.Data["json"] = ReturnError(4004, "此手机号已经注册")
		err = u.ServeJSON()
		if err != nil {
			return
		}
	} else {
		err = models.UserSave(mobile, MD5V(password))
		if	err == nil {
			u.Data["json"] = ReturnSuccess(0, "注册成功", nil, 0)
			err := u.ServeJSON()
			if err != nil {
				return
			}
		} else {
			u.Data["json"] = ReturnError(5000, err)
			err := u.ServeJSON()
			if err != nil {
				return
			}
		}
	}
}

// LoginDo 用户登录
// @router /login/do [*]
func (u *UserController) LoginDo()  {
	var (
		mobile string
		password string
		err error
	)
	mobile = u.GetString("mobile")
	password = u.GetString("password")

	checkJson := checkMobileAddPassword(mobile, password)
	if checkJson != nil {
		u.Data["json"] = checkJson
		err = u.ServeJSON()
		if err != nil {
			return
		}
	}

	uid, name := models.HandleMobileLogin(mobile, MD5V(password))
	if uid != 0 {
		u.Data["json"] = ReturnSuccess(0, "登录成功", map[string]interface{}{"uid":uid, "username": name}, 1)
		u.ServeJSON()
	} else {
		u.Data["json"] = ReturnError(4004, "手机号或密码不正确")
		u.ServeJSON()
	}
}

func checkMobileAddPassword(mobile string, password string)  *JsonStruct {
	if mobile == "" {
		return ReturnError(4001, "手机号码不能为空")
	}

	isOrNo, _ := regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, mobile)
	if !isOrNo {
		return ReturnError(4002, "手机号码格式不正确")
	}

	if password == "" {
		return ReturnError(4003, "密码不能为空")
	}

	return nil
}