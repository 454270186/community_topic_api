package service

import (
	"sync"

	"github.com/454270186/CommuTopicPage/repository"
)

// Service层流程: 参数校验 -> 准备数据 -> 组装实体

type PageInfo struct {
	Topic *repository.Topic
	PostList []*repository.Post
}

type PageService struct {
	repo repository.DataRepo
}

func NewPageService(repo repository.DataRepo) PageService {
	return PageService{
		repo: repo,
	}
}

func (ps PageService) QueryPageInfo(topicId int64) (*PageInfo, error) {
	var topicErr error = nil
	var postsErr error = nil
	var pageInfo PageInfo
	var wg sync.WaitGroup
	wg.Add(2)

	go func ()  {
		defer wg.Done()
		topic, err := ps.repo.FindById(topicId)
		if err != nil {
			topicErr = err
			return
		}
		
		pageInfo.Topic = topic
	} ()

	go func ()  {
		defer wg.Done()
		posts, err := ps.repo.FindByParentId(topicId)
		if err != nil {
			postsErr = err
			return
		}

		pageInfo.PostList = posts
	} ()

	wg.Wait()

	if topicErr != nil {
		return nil, topicErr
	} else if postsErr != nil {
		return nil, postsErr
	}

	return &pageInfo, nil
}

func (ps PageService) AddNewTopic(topic repository.Topic) (int64, error) {
	newTopicId, err := ps.repo.NewTopic(topic)
	if err != nil {
		return 0, err
	}

	return newTopicId, nil
}

func (ps PageService) AddNewPost(post repository.Post) (int64, error) {
	postId, err := ps.repo.NewPost(post)
	if err != nil {
		return 0, err
	}
	
	return postId, nil
}

func (ps PageService) DeleteTopic(id int64) error {
	err := ps.repo.DelTopic(id)
	if err != nil {
		return err
	}

	return nil
}
