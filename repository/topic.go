package repository

import (
	"sync"
	"time"
)

type Topic struct {
	Id         int64     `gorm:"primaryKey"`
	Title      string    `gorm:"column:title"`
	Content    string    `gorm:"column:content"`
	CreateTime time.Time `gorm:"column:create_time"`
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
