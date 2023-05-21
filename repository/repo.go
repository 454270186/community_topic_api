package repository

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

type DataRepo interface {
	FindById(id int64) (*Topic, error)
	FindByParentId(parentId int64) ([]*Post, error)
	NewPost(post Post) error
	NewTopic(topic Topic) (int64, error)
}

func NewDataRepo(db *gorm.DB) DataRepo {
	return DataRepoDB{
		DB: db,
	}
}

type DataRepoDB struct {
	DB *gorm.DB
}

// find topic by id
func (p DataRepoDB) FindById(id int64) (*Topic, error) {
	var topic Topic
	if err := p.DB.First(&topic, id).Error; err != nil {
		log.Println("Error while find topic by id")
		return nil, err
	}

	return &topic, nil
}

// find post by parent id
func (p DataRepoDB) FindByParentId(parentId int64) ([]*Post, error) {
	var posts []*Post
	if err := p.DB.Where("parent_id = ?", parentId).Find(&posts).Error; err != nil {
		log.Println("Error while find posts by parent id")
		return nil, err
	}

	return posts, nil
}

// create new topic
func (p DataRepoDB) NewTopic(topic Topic) (int64, error) {
	topic.Id = 0
	result := p.DB.Create(&topic)
	if result.Error != nil {
		log.Println("Error while insert a new topic")
		return 0, errors.New("unexpect database error")
	}

	return topic.Id, nil
}

// Create new post by parent id
func (p DataRepoDB) NewPost(post Post) error {
	var cnt int64
	result := p.DB.Model(&Topic{}).Where("id = ?", post.ParentId).Count(&cnt)
	if result.Error != nil {
		log.Println("Error while check topic id")
		return errors.New("unexpect database error")
	} else if cnt == 0 {
		return errors.New("topic id does not exist")
	}

	result = p.DB.Create(&post)
	if result.Error != nil {
		log.Println("Error while insert a new post")
		return errors.New("unexpect database error")
	}

	return nil
}