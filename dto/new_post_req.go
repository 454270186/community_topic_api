package dto

type NewPostReq struct {
	TopicId int64 `json:"topic_id"`
	Content string `json:"content"`
}