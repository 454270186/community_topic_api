package repository

import (
	"sync"
	"time"
)

type Post struct {
	Id         int64  `json:"id"`
	ParentId   int64  `json:"parent_id"`
	Content    string `json:"content"`
	CreateTime time.Time  `json:"create_time"`
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
