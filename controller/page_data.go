package controller

import (
	"strconv"
	"time"

	"github.com/454270186/CommuTopicPage/dto"
	"github.com/454270186/CommuTopicPage/repository"
	"github.com/454270186/CommuTopicPage/service"
)

type PageData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
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

func (ctrl Controller) AddNewPost(post dto.NewPostReq) *PageData {
	if len(post.Content) == 0 {
		return &PageData{
			Code: -1,
			Msg: "content can't be empty",
		}
	}

	newPost := repository.Post{
		ParentId: post.TopicId,
		Content: post.Content,
		CreateTime: time.Now(),
	}

	if err := ctrl.pageService.AddNewPost(post.TopicId, newPost); err != nil {
		return &PageData{
			Code: -1,
			Msg: err.Error(),
		}
	}

	return &PageData{
		Code: 0,
		Msg: "success",
	}
}