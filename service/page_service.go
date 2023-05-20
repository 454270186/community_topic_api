package service

import (
	"errors"
	"sync"

	"github.com/454270186/CommuTopicPage/repository"
)

// Service层流程: 参数校验 -> 准备数据 -> 组装实体

type PageInfo struct {
	Topic *repository.Topic
	PostList []*repository.Post
}

type PageService struct {
	Repo repository.DataRepo
}

func NewPageService(repo repository.DataRepo) PageService {
	return PageService{
		Repo: repo,
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
		topic, err := ps.Repo.FindById(topicId)
		if err != nil {
			topicErr = err
			return
		}
		
		pageInfo.Topic = topic
	} ()

	go func ()  {
		defer wg.Done()
		posts, err := ps.Repo.FindByParentId(topicId)
		if err != nil {
			postsErr = err
			return
		}

		pageInfo.PostList = posts
	} ()

	wg.Wait()

	if topicErr != nil || postsErr != nil {
		return nil, errors.New("error while get page info")
	}

	return &pageInfo, nil
}
