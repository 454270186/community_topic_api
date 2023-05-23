### 社区话题页面 Backend API

#### Based on
- Gin
- Redis
- PostgreSQL
- GORM
- Go-Viper

#### API
- GET ```/community/page/:id``` 获取指定id页面的所有话题和回帖数据
- POST ```/community/page/topic``` 发布新的topic
    - post body
    ```json
    {
        "title": <topic title>,
        "content": <topic content>
    }
    ```
- POST ```/community/page/post``` 在指定话题下发布新的post
    - post body 
    ```json
    {
        "topic_id": <topic id>,
        "content": <post content>
    }
    ```
- DELETE ```/community/page/topic/:id``` 删除指定topic
- DELETE ```/community/page/post/:id``` 删除指定post

- PUT ```/post/:postid/like``` 点赞指定post
- GET ```"/post/:topicid/like``` 获取指定topic下的所有post(按点赞数降序排列)
