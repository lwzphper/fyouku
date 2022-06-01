package models

import "github.com/beego/beego/v2/client/orm"

type Region struct {
	Id   int
	Name string
}
type Type struct {
	Id   int
	Name string
}

func GetChannelRegion(channelId int) (num int64, regions []Region, err error) {
	o := orm.NewOrm()
	num, err = o.Raw("SELECT id,name FROM channel_region WHERE status=1 AND channel_id=? ORDER BY sort DESC", channelId).QueryRows(&regions)
	return
}

func GetChannelType(channelId int) (num int64, types []Type, err error) {
	o := orm.NewOrm()
	num, err = o.Raw("SELECT id,name FROM channel_type WHERE status=1 AND channel_id=? ORDER BY sort DESC", channelId).QueryRows(&types)
	return
}