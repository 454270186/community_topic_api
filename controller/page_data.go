package controller

import (
	"strconv"
	"time"

	"github.com/454270186/CommuTopicPage/dto"
	"github.com/454270186/CommuTopicPage/repository"
	"github.com/454270186/CommuTopicPage/service"
	"github.com/gin-gonic/gin"
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

func (ctrl Controller) AddNewTopic(topic dto.NewTopicReq) *PageData {
	if len(topic.Title) == 0 || len(topic.Content) == 0 {
		return &PageData{
			Code: -1,
			Msg: "title and content can't be empty",
		}
	}

	newTopic := repository.Topic{
		Title: topic.Title,
		Content: topic.Content,
		CreateTime: time.Now(),
	}

	newTopicId, err := ctrl.pageService.AddNewTopic(newTopic)
	if err != nil {
		return &PageData{
			Code: -1,
			Msg: err.Error(),
		}
	}
	

	return &PageData{
		Code: 0,
		Msg: "success",
		Data: gin.H{
			"new_topic_id": newTopicId,
		},
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

	newPostId, err := ctrl.pageService.AddNewPost(newPost)
	if err != nil {
		return &PageData{
			Code: -1,
			Msg: err.Error(),
		}
	}

	return &PageData{
		Code: 0,
		Msg: "success",
		Data: gin.H{"new_post_id": newPostId},
	}
}

func (ctrl Controller) DeleteTopic(idStr string) *PageData {
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return &PageData{
			Code: -1,
			Msg: err.Error(),
		}
	}

	err = ctrl.pageService.DeleteTopic(id)
	if err != nil {
		return &PageData{
			Code: -1,
			Msg: err.Error(),
		}
	}

	return &PageData{
		Code: 0,
		Msg: "delete topic successfully",
	}
}

func (ctrl Controller) DeletePost(idStr string) *PageData {
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return &PageData{
			Code: -1,
			Msg: err.Error(),
		}
	}

	err = ctrl.pageService.DeletePost(id)
	if err != nil {
		return &PageData{
			Code: -1,
			Msg: err.Error(),
		}
	}

	return &PageData{
		Code: 0,
		Msg: "delete post successfully",
	}
}

func (ctrl Controller) AddPostLike(postIdStr string) *PageData {
	postId, err := strconv.ParseInt(postIdStr, 10, 64)
	if err != nil {
		return &PageData{
			Code: -1,
			Msg: err.Error(),
		}
	}

	curLikeCnt, err := ctrl.pageService.AddPostLike(postId)
	if err != nil {
		return &PageData{
			Code: -1,
			Msg: err.Error(),
		}
	}

	return &PageData{
		Code: 0,
		Msg: "add like successfully",
		Data: gin.H{"current_likes": curLikeCnt},
	}
}

func (ctrl Controller) GetPostByLike(topicIdStr string) *PageData {
	topicId, err := strconv.ParseInt(topicIdStr, 10, 64)
	if err != nil {
		return &PageData{
			Code: -1,
			Msg: err.Error(),
		}
	}

	postLists, err := ctrl.pageService.GetPostByLike(topicId)
	if err != nil {
		return &PageData{
			Code: -1,
			Msg: err.Error(),
		}
	}

	return &PageData{
		Code: 0,
		Msg: "success",
		Data: postLists,
	}
}