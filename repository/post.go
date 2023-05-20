package repository

import (
	"sync"
	"time"
)

type Post struct {
	Id         int64     `gorm:"primaryKey"`
	ParentId   int64     `gorm:"column:parent_id"`
	Content    string    `gorm:"column:content"`
	CreateTime time.Time `gorm:"column:create_time"`
}

type PostDao struct {
}

var (
	postDao  *PostDao
	postOnce sync.Once
)

func NewPostDaoInstance() *PostDao {
	postOnce.Do(func() {
		postDao = &PostDao{}
	})

	return postDao
}

func (*PostDao) QueryByParentId(parentId int64) []*Post {
	return postIndexMap[parentId]
}
