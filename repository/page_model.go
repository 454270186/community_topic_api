package repository

import (
	"time"

	"gorm.io/gorm"
)

type Topic struct {
	Id         int64     `gorm:"primaryKey" json:"id"`
	Title      string    `gorm:"column:title" json:"title"`
	Content    string    `gorm:"column:content" json:"content"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
}

type Post struct {
	Id         int64     `gorm:"primaryKey" json:"id"`
	ParentId   int64     `gorm:"column:parent_id" json:"parent_id"`
	Content    string    `gorm:"column:content" json:"content"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	LikeCnt    int64     `gorm:"column:like_count" json:"like_count"`
}

type DataRepo interface {
	FindById(id int64) (*Topic, error)
	FindPostById(id int64) (*Post, error)
	FindByParentId(parentId int64) ([]*Post, error)
	NewPost(post Post) (*Post, error)
	NewTopic(topic Topic) (int64, error)
	DelTopic(id int64) error
	DelPost(id int64) (int64, error)
	AddPostLike(id int64) (int64, int64, error)
}

func NewDataRepo(db *gorm.DB) DataRepo {
	return DataRepoDB{
		DB: db,
	}
}
