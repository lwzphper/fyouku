package controllers

import (
	"fyouku/models"
	beego "github.com/beego/beego/v2/server/web"
)

type VideoController struct {
	beego.Controller
}

// ChannelAdvert 频道页 - 获取顶部广告
// @router /channel/advert [*]
func (v *VideoController) ChannelAdvert() {
	channelId,_ := v.GetInt("channelId")

	if channelId == 0 {
		v.Data["json"] = ReturnError(4001, "必须指定频道")
		v.ServeJSON()
	}

	num, videos, err := models.GetChannelAdvert(channelId)
	if err == nil {
		v.Data["json"] = ReturnSuccess(0, "success", videos, num)
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnError(4004, "请求数据失败，请稍后重试~")
		v.ServeJSON()
	}
}

// ChannelHotList 频道页-获取正在热播
// @router /channel/hot [*]
func (v *VideoController) ChannelHotList() {
	channelId,_ := v.GetInt("channelId")

	if channelId == 0 {
		v.Data["json"] = ReturnError(4001, "必须指定频道")
		v.ServeJSON()
	}

	num, videos, err := models.GetChannelHotList(channelId)
	if err == nil {
		v.Data["json"] = ReturnSuccess(0, "success", videos, num)
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnError(4004, "请求数据失败，请稍后重试~")
		v.ServeJSON()
	}
}

// ChannelRecommendRegionList 频道页-根据频道地区获取推荐的视频
// @router /channel/recommend/region [*]
func (v *VideoController) ChannelRecommendRegionList() {
	channelId, _ := v.GetInt("channelId")
	regionId, _ := v.GetInt("regionId")

	if channelId == 0 {
		v.Data["json"] = ReturnError(4001, "必须指定频道")
		v.ServeJSON()
	}
	if regionId == 0 {
		v.Data["json"] = ReturnError(4002, "必须指定频道地区")
		v.ServeJSON()
	}

	num, videos, err := models.GetChannelRecommendRegionList(channelId, regionId)
	if err == nil {
		v.Data["json"] = ReturnSuccess(0, "success", videos, num)
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnError(4004, "没有相关内容")
		v.ServeJSON()
	}
}

// GetChannelRecommendTypeList 频道页-根据频道类型获取推荐视频
// @router /channel/recommend/type [*]
func (v *VideoController) GetChannelRecommendTypeList() {
	channelId, _ := v.GetInt("channelId")
	typeId, _ := v.GetInt("typeId")

	if channelId == 0 {
		v.Data["json"] = ReturnError(4001, "必须指定频道")
		v.ServeJSON()
	}
	if typeId == 0 {
		v.Data["json"] = ReturnError(4002, "必须指定频道类型")
		v.ServeJSON()
	}

	num, videos, err := models.GetChannelRecommendTypeList(channelId, typeId)
	if err == nil {
		v.Data["json"] = ReturnSuccess(0, "success", videos, num)
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnError(4004, "没有相关内容")
		v.ServeJSON()
	}
}

// ChannelVideo 根据传入参数获取视频列表
// @router /channel/video [*]
func (v *VideoController) ChannelVideo() {
	//获取频道ID
	channelId, _ := v.GetInt("channelId")
	//获取频道地区ID
	regionId, _ := v.GetInt("regionId")
	//获取频道类型ID
	typeId, _ := v.GetInt("typeId")
	//获取状态
	end := v.GetString("end")
	//获取排序
	sort := v.GetString("sort")
	//获取页码信息
	limit, _ := v.GetInt("limit")
	offset, _ := v.GetInt("offset")

	if channelId == 0 {
		v.Data["json"] = ReturnError(4001, "必须指定频道")
		v.ServeJSON()
	}

	if limit == 0 {
		limit = 12
	}

	num, videos, err := models.GetChannelVideoList(channelId, regionId, typeId, end, sort, offset, limit)
	if err == nil {
		v.Data["json"] = ReturnSuccess(0, "success", videos, num)
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnError(4004, "没有相关内容")
		v.ServeJSON()
	}
}

// UserVideo 我的视频管理
// @router /user/video [*]
func (v *VideoController) UserVideo() {
	uid, _ := v.GetInt("uid")
	if uid == 0 {
		v.Data["json"] = ReturnError(4001, "必须指定用户")
		v.ServeJSON()
	}
	num, videos, err := models.GetUserVideo(uid)
	if err == nil {
		v.Data["json"] = ReturnSuccess(0, "success", videos, num)
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnError(4004, "没有相关内容")
		v.ServeJSON()
	}
}

// VideoInfo 获取视频详情
// @router /video/info [*]
func (v *VideoController) VideoInfo() {

}

// VideoEpisodesList 获取视频剧集列表
// @router /video/episodes/list [*]
func (v *VideoController) VideoEpisodesList() {

}

// VideoSave 保存用户上传视频信息
// @router /video/save [*]
func (v *VideoController) VideoSave() {

}

// Search 搜索接口
// @router /video/search [*]
func (v *VideoController) Search() {

}

// SaveAll 生成测试视频数据
// @router /video/save/all [*]
func (v *VideoController) SaveAll() {

}

// SendEs 导入ES脚本
// @router /video/send/es [*]
func (v *VideoController) SendEs() {

}