package models

import "time"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorID  int       `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	Comments  []Comment `json:"comments"`
}

type Comment struct {
	ID        int       `json:"id"`
	PostID    int       `json:"post_id"`
	Content   string    `json:"content"`
	AuthorID  int       `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
}

var Users = []User{
	{ID: 1, Username: "admin", Password: "$2a$10$AtYqgFq8E/nyQVoNP9IXOuVr6TbWXe5fC7kfZOzIh7Vk8eQmjywuu", Role: "admin"},
	{ID: 2, Username: "user1", Password: "$2a$10$AtYqgFq8E/nyQVoNP9IXOuVr6TbWXe5fC7kfZOzIh7Vk8eQmjywuu", Role: "user"},
}

var Posts = []Post{
	{ID: 1, AuthorID: 1, Title: "First Post", Content: "This is the first post", CreatedAt: time.Now(), Comments: Comments},
}

var Comments = []Comment{
	{ID: 1, PostID: 1, AuthorID: 2, Content: "First comment", CreatedAt: time.Now()},
}

func (u *User) Create() string {
	Users = append(Users, *u)
	return "User created successfully"
}

func GetPosts() []Post {
	return Posts
}

func GetPostByID(id int) *Post {
	for i := range Posts {
		if Posts[i].ID == id {
			return &Posts[i]
		}
	}
	return nil
}

func (p *Post) Create() string {
	p.ID = len(Posts) + 1
	p.CreatedAt = time.Now()
	Posts = append(Posts, *p)
	return "Post created successfully"
}

func DeletePost(id int) string {
	for i := range Posts {
		if Posts[i].ID == id {
			Posts = append(Posts[:i], Posts[i+1:]...)
			return "Post deleted successfully"
		}
	}
	return "Post not found"
}

func (c *Comment) Create(postID int) string {
	c.ID = len(Comments) + 1
	c.PostID = postID
	c.CreatedAt = time.Now()
	Comments = append(Comments, *c)
	return "Comment added successfully"
}
