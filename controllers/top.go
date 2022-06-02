package controllers

import (
	"fyouku/models"
	beego "github.com/beego/beego/v2/server/web"
)

type TopController struct {
	beego.Controller
}

// ChannelTop 根据频道获取排行榜
// @router /channel/top [*]
func (t *TopController) ChannelTop() {
	channelId, _ := t.GetInt("channelId")
	if channelId == 0 {
		t.Data["json"] = ReturnError(4001, "必须指定频道")
		t.ServeJSON()
	}

	num, videos, err := models.GetChannelTop(channelId)
	if err == nil {
		t.Data["json"] = ReturnSuccess(0, "success", videos, num)
		t.ServeJSON()
	} else {
		t.Data["json"] = ReturnError(4004, "没有相关内容")
		t.ServeJSON()
	}
}

//根据类型获取排行榜
// @router /type/top [*]
func (t *TopController) TypeTop() {
	typeId, _ := t.GetInt("typeId")
	if typeId == 0 {
		t.Data["json"] = ReturnError(4001, "必须指定类型")
		t.ServeJSON()
	}
	num, videos, err := models.GetTypeTop(typeId)
	if err == nil {
		t.Data["json"] = ReturnSuccess(0, "success", videos, num)
		t.ServeJSON()
	} else {
		t.Data["json"] = ReturnError(4004, "没有相关内容")
		t.ServeJSON()
	}
}