package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type User struct {
	Id       int
	Name     string
	Password string
	Status   int
	AddTime  int64
	Mobile   string
	Avatar   string
}

type UserInfo struct {
	Id      int    `json:"id"`
	Name    int    `json:"name"`
	AddTime int64  `json:"addTime"`
	Avatar  string `json:"avatar"`
}

func init() {
	orm.RegisterModel(new(User))
}

// ExistUserMobile 判断手机号是否存在
func ExistUserMobile(mobile string) bool {
	o := orm.NewOrm()
	user := User{Mobile: mobile}
	err := o.Read(&user, "Mobile")
	if err == orm.ErrNoRows {
		return false
	} else if err == orm.ErrMissPK {
		return false
	}
	return true
}

// UserSave 保存用户
func UserSave(mobile string, password string) error {
	o := orm.NewOrm()
	user := new(User)
	user.Name = ""
	user.Mobile = mobile
	user.Password = password
	user.Status = 1
	user.AddTime = time.Now().Unix()
	_, err := o.Insert(user)
	return err
}

// HandleMobileLogin 登录功能
func HandleMobileLogin(mobile string, password string) (int, string) {
	o := orm.NewOrm()
	user := User{}
	err := o.QueryTable("user").Filter("mobile", mobile).Filter("password", password).One(&user)
	if err == orm.ErrNoRows {
		return 0, ""
	} else if err == orm.ErrMissPK {
		return 0, ""
	}
	return user.Id, user.Name
}

// GetUserInfo 获取用户信息
func GetUserInfo(uid int) (UserInfo, error) {
	userInfo := UserInfo{}
	o := orm.NewOrm()
	err := o.Raw("SELECT id,name,add_time,avatar FROM user where id = ? LIMIT 1", uid).QueryRow(&userInfo)
	return userInfo, err
}
