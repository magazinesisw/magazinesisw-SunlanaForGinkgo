package public

import "time"

// Posts Posts是数据库中的posts表,用于数据库查询后进行结构体映射
type Posts struct {
	PostId      int64     `json:"post_id"`
	UserId      int64     `json:"user_id"`
	PostTitle   string    `json:"post_title"`
	PostContent string    `json:"post_content"`
	PostType    int       `json:"post_type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"update_at"`
	PostStatus  bool      `json:"post_status"`
}

// PostsObj PostsObj是接口sunlana.com/api/v1/posts?page=1返回的格式用于反序列化
type PostsObj struct {
	PostId    int64             `json:"post_id"`
	PostTitle string            `json:"post_title"`
	PostImage map[string]string `json:"post_image"`
	UserName  string            `json:"user_name"`
	Avatar    string            `json:"avatar"`
	CreatedAt time.Time         `json:"created_at"`
}
