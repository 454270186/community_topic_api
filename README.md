### 社区话题页面 Backend API

#### Based on
- Gin
- 分层架构
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
