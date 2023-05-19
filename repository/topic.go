package repository

import "sync"

type Topic struct {
	Id         int64  `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

type TopicDao struct {
}

var (
	topicDao  *TopicDao
	topicOnce sync.Once
)

// 使用sync.Once，确保高并发场景只执行一次，实现单例模式
func NewTopicDaoInstance() *TopicDao {
	topicOnce.Do(func() {
		topicDao = &TopicDao{}
	})

	return topicDao
}

func (*TopicDao) QueryById(id int64) *Topic {
	return topicIndexMap[id]
}




