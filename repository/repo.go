package repository

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

type DataRepoDB struct {
	DB *gorm.DB
}

// find topic by id
func (p DataRepoDB) FindById(id int64) (*Topic, error) {
	var topic Topic
	if err := p.DB.First(&topic, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("topic not found")
		}

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
func (p DataRepoDB) NewPost(post Post) (int64, error) {
	var cnt int64
	result := p.DB.Model(&Topic{}).Where("id = ?", post.ParentId).Count(&cnt)
	if result.Error != nil {
		log.Println("Error while check topic id")
		return 0, errors.New("unexpect database error")
	} else if cnt == 0 {
		return 0, errors.New("topic id does not exist")
	}

	result = p.DB.Create(&post)
	if result.Error != nil {
		log.Println("Error while insert a new post")
		return 0, errors.New("unexpect database error")
	}

	return post.Id, nil
}

func (p DataRepoDB) DelTopic(id int64) error {
	result := p.DB.Delete(&Topic{}, id)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return errors.New("unexpect database error")
	}
	if result.RowsAffected == 0 {
		return errors.New("topic id does not exist")
	}

	return nil
}

func (p DataRepoDB) DelPost(id int64) (int64, error) {
	var post Post
	err := p.DB.First(&post, id).Error
	if err != nil {
		return 0, errors.New("post id does not exist")
	}

	parentId := post.ParentId

	result := p.DB.Delete(&post)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return 0, errors.New("unexpect database error")
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("post id does not exist")
	}

	return parentId, nil
}

func (p DataRepoDB) AddPostLike(id int64) (int64, int64, error) {
	var curLikeCnt int64
	var topicId int64
	err := p.DB.Transaction(func(tx *gorm.DB) error {
		var post Post
		err := p.DB.First(&post, id).Error
		if err != nil {
			return err
		}

		post.LikeCnt++
		curLikeCnt = post.LikeCnt
		topicId = post.ParentId

		if err := p.DB.Model(&post).Update("like_count", post.LikeCnt).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		log.Println("Error while add post likes")
		return -1, -1, errors.New("unexpect database error")
	}

	return curLikeCnt, topicId, nil
}