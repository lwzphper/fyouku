package models

import "github.com/beego/beego/v2/client/orm"

type Advert struct {
	Id int
	Title string
	SubTitle string
	AddTime int64
	Img string
	Url string
}

func init()  {
	orm.RegisterModel(new(Advert))
}

func GetChannelAdvert(channelId int) (num int64, adverts []Advert, err error)  {
	o := orm.NewOrm()
	num, err = o.Raw("SELECT id, title, sub_title,img,add_time,url FROM advert WHERE status=1 AND channel_id=? ORDER BY sort DESC LIMIT 1", channelId).QueryRows(&adverts)
	return
}