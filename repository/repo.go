package repository

import (
	"log"

	"gorm.io/gorm"
)

type DataRepo interface {
	FindById(id int64) (*Topic, error)
	FindByParentId(parentId int64) ([]*Post, error)
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