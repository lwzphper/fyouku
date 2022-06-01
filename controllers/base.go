package controllers

import (
	"fyouku/models"
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
}

// ChannelRegion 获取频道地区列表
// @router /channel/region [*]
func (b *BaseController) ChannelRegion() {
	channelId, _ := b.GetInt("channelId")
	if channelId == 0 {
		b.Data["json"] = ReturnError(4001, "必须指定频道")
		b.ServeJSON()
	}
	num, regions, err := models.GetChannelRegion(channelId)
	if err == nil {
		b.Data["json"] = ReturnSuccess(0, "success", regions, num)
		b.ServeJSON()
	} else {
		b.Data["json"] = ReturnError(4004, "没有相关内容")
		b.ServeJSON()
	}
}

// ChannelType 获取频道类型列表
// @router /channel/type [*]
func (b *BaseController) ChannelType() {
	channelId, _ := b.GetInt("channelId")
	if channelId == 0 {
		b.Data["json"] = ReturnError(4001, "必须指定频道")
		b.ServeJSON()
	}

	num, types, err := models.GetChannelType(channelId)
	if err == nil {
		b.Data["json"] = ReturnSuccess(0, "success", types, num)
		b.ServeJSON()
	} else {
		b.Data["json"] = ReturnError(4004, "没有相关内容")
		b.ServeJSON()
	}
}