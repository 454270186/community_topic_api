package service

import (
	"errors"
	"log"
	"sync"

	"github.com/454270186/CommuTopicPage/repository"
)

// Service层流程: 参数校验 -> 准备数据 -> 组装实体

type PageInfo struct {
	Topic *repository.Topic
	PostList []*repository.Post
}

func QueryPageInfo(topicId int64) (*PageInfo, error) {
	return NewQueryInfoFlow(topicId).Do()
}

func NewQueryInfoFlow(topicId int64) *QueryPageInfoFlow {
	return &QueryPageInfoFlow{
		topicId: topicId,
		pageInfo: &PageInfo{},
	}
}

type QueryPageInfoFlow struct {
	topicId int64
	pageInfo *PageInfo
}

func (q *QueryPageInfoFlow) Do() (*PageInfo, error) {
	if err := q.checkParam(); err != nil {
		return nil, err
	}

	if err := q.prepareInfo(); err != nil {
		return nil, err
	}
	log.Println("page info is ")
	log.Println(q.pageInfo)
	return q.pageInfo, nil
}

func (q *QueryPageInfoFlow) checkParam() error {
	if q.topicId <= 0 {
		return errors.New("topic id must be larger than 0")
	}

	return nil
}

func (q *QueryPageInfoFlow) prepareInfo() error {
	// 获取topic和posts信息
	var wg sync.WaitGroup
	wg.Add(2)

	go func ()  {
		defer wg.Done()
		topic := repository.NewTopicDaoInstance().QueryById(q.topicId)
		log.Println(topic)
		q.pageInfo.Topic = topic
	}()

	go func ()  {
		defer wg.Done()
		posts := repository.NewPostDaoInstance().QueryByParentId(q.topicId)
		q.pageInfo.PostList = posts
	}()
	
	wg.Wait()
	return nil
}

// func (q *QueryPageInfoFlow) packageInfo() error {

// }
