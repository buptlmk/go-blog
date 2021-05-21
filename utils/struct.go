package utils

import "time"

// ---------------------article-------------------------

type Comment struct {
	UserId      int64     `json:"userId"`
	User        string    `json:"user"`
	ArticleId   int64     `json:"articleId"`
	Content     string    `json:"content"`
	CreatedTime time.Time `json:"createdTime"`
}

type Article struct {
	Id            int64     `json:"id"`
	AuthorId      int64     `json:"authorId"`
	Author        string    `json:"author"`
	Title         string    `json:"title"`
	CreatedTime   time.Time `json:"createdTime"`
	VisitedNumber int64     `json:"visitedNumber"`
	StarNumber    int64     `json:"starNumber"`
	LikeNumber    int64     `json:"likeNumber"`
	CommentNumber int64     `json:"commentNumber"`
	Content       string    `json:"content"`
}

// ----------------user-----------
type User struct {
	Id     int64  `json:"id"`
	CardId string `json:"cardId"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
	Github string `json:"github"`
	QQ     string `json:"qq"`
	Wechat string `json:"wechat"`
}

type Activity struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Time        string `json:"time"`
	Img         string `json:"img"`
	Total       int    `json:"total"`
	Res         int    `json:"res"`
}

type Chat struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	PersonNumber int    `json:"personNumber"`
}
