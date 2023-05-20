package controller

import (
	"strconv"

	"github.com/454270186/CommuTopicPage/service"
)

type PageData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Controller struct {
	pageService service.PageService
}

func NewController(ps service.PageService) Controller {
	return Controller{
		pageService: ps,
	}
}

func (ctrl Controller) QueryPageInfo(topicIdStr string) *PageData {
	topicId, err := strconv.ParseInt(topicIdStr, 10, 64)
	if err != nil {
		return &PageData{
			Code: -1,
			Msg: err.Error(),
		}
	}

	pageInfo, err := ctrl.pageService.QueryPageInfo(topicId)
	if err != nil {
		return &PageData{
			Code: -1,
			Msg: err.Error(),
		}
	}

	return &PageData{
		Code: 0,
		Msg: "success",
		Data: pageInfo,
	}
}