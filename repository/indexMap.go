package repository

import (
	"bufio"
	"log"
	"os"

	"github.com/goccy/go-json"
)

var (
	topicIndexMap map[int64]*Topic
	postIndexMap map[int64][]*Post
)

func InitIndexMap(filepath string) error {
	if err := initTopicIndexMap(filepath); err != nil {
		return err
	}

	if err := initPostIndexMap(filepath); err != nil {
		return err
	}

	log.Println(topicIndexMap)
	log.Println(postIndexMap)

	return nil
}

func initTopicIndexMap(filepath string) error {
	open, err := os.Open(filepath + "topic.txt")
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(open)
	topicTempMap := make(map[int64]*Topic)

	for scanner.Scan() {
		text := scanner.Text()
		var topic Topic
		if err := json.Unmarshal([]byte(text), &topic); err != nil {
			return err
		}
		log.Println(topic)
		topicTempMap[topic.Id] = &topic
	}

	topicIndexMap = topicTempMap
	return nil
}

func initPostIndexMap(filepath string) error {
	open, err := os.Open(filepath + "post.txt")
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(open)
	postTempMap := make(map[int64][]*Post)

	for scanner.Scan() {
		text := scanner.Text()
		var post Post
		if err := json.Unmarshal([]byte(text), &post); err != nil {
			return err
		}

		postTempMap[post.ParentId] = append(postTempMap[post.ParentId], &post)
	}

	postIndexMap = postTempMap
	return nil
}