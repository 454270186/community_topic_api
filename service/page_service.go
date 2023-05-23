package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"

	"github.com/454270186/CommuTopicPage/repository"
	"github.com/redis/go-redis/v9"
)

// Service层流程: 参数校验 -> 准备数据 -> 组装实体

type PageInfo struct {
	Topic *repository.Topic
	PostList []*repository.Post
}

type PageService struct {
	repo repository.DataRepo
	rdb *redis.Client
}

func NewPageService(repo repository.DataRepo, redisdb *redis.Client) PageService {
	return PageService{
		repo: repo,
		rdb: redisdb,
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

		key := fmt.Sprintf("topic:%d", topicId)
		val, err := ps.rdb.Get(context.Background(), key).Result()
		if err == redis.Nil {
			fmt.Println("Not found in Redis, Query from DB")
			topic, err := ps.repo.FindById(topicId)
			if err != nil {
				topicErr = err
				return
			}

			data, _ := json.Marshal(topic)
			err = ps.rdb.Set(context.Background(), key, data, 0).Err()
			if err != nil {
				topicErr = err
				return
			}

			pageInfo.Topic = topic
		} else if err != nil {
			topicErr = err
			return
		} else {
			var topic repository.Topic
			err := json.Unmarshal([]byte(val), &topic)
			if err != nil {
				topicErr = err
				return
			}

			pageInfo.Topic = &topic
		}
	} ()

	go func ()  {
		defer wg.Done()
		key := fmt.Sprintf("post:%d", topicId)
		val, err := ps.rdb.Get(context.Background(), key).Result()
		if err == redis.Nil {
			fmt.Println("Not found in Redis, Query from DB")
			posts, err := ps.repo.FindByParentId(topicId)
			if err != nil {
				postsErr = err
				return
			}
			
			data, _ := json.Marshal(posts)
			err = ps.rdb.Set(context.Background(), key, data, 0).Err()
			if err != nil {
				postsErr = err
				return
			}

			pageInfo.PostList = posts
		} else if err != nil {
			postsErr = err
			return
		} else {
			var posts []*repository.Post
			err := json.Unmarshal([]byte(val), &posts)
			if err != nil {
				topicErr = err
				return
			}

			pageInfo.PostList = posts
		}

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
	newPost, err := ps.repo.NewPost(post)
	if err != nil {
		return 0, err
	}

	key := fmt.Sprintf("post:%d", post.ParentId)
	_, err = ps.rdb.Del(context.Background(), key).Result()
	if err != nil {
		return 0, err
	}

	// add to redis like sort
	likeKey := fmt.Sprintf("post:%d:like:", post.ParentId)
	member := newPost.Id
	score := 0
	err = ps.rdb.ZAdd(context.Background(), likeKey, redis.Z{Score: float64(score), Member: member}).Err()
	if err != nil {
		return 0, err
	}
	
	return newPost.Id, nil
}

func (ps PageService) DeleteTopic(id int64) error {
	err := ps.repo.DelTopic(id)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("topic:%d", id)
	_, err = ps.rdb.Del(context.Background(), key).Result()
	if err != nil {
		return err
	}

	return nil
}

func (ps PageService) DeletePost(id int64) error {
	delPostParentID, err := ps.repo.DelPost(id)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("post:%d", delPostParentID)
	_, err = ps.rdb.Del(context.Background(), key).Result()
	if err != nil {
		return err
	}

	// del postid in redis like list
	likeKey := fmt.Sprintf("post:%d:like:", delPostParentID)
	member := id
	err = ps.rdb.ZRem(context.Background(), likeKey, member).Err()
	if err != nil {
		return err
	}

	return nil
}

// return current post likes count after update
func (ps PageService) AddPostLike(postId int64) (int64, error) {
	curLikeCnt, topicId, err := ps.repo.AddPostLike(postId)
	if err != nil {
		return -1, err
	}

	// add in Redis sorted
	// key => post:[topicId]:like
	key := fmt.Sprintf("post:%d:like:", topicId)
	member := postId
	score := curLikeCnt

	err = ps.rdb.ZAdd(context.Background(), key, redis.Z{Score: float64(score), Member: member}).Err()
	if err != nil {
		return -1, err
	}

	return curLikeCnt, err
}

func (ps PageService) GetPostByLike(topicId int64) ([]*repository.Post, error) {
	key := fmt.Sprintf("post:%d:like:", topicId)

	result, err := ps.rdb.ZRevRange(context.Background(), key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	var postLists []*repository.Post
	for _, member := range result {
		postId, _ := strconv.ParseInt(member, 10, 64)
		post, err := ps.repo.FindPostById(postId)
		if err != nil {
			return nil, err
		}

		postLists = append(postLists, post)
	}

	return postLists, nil
}