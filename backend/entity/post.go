package entity

import "time"

type Post struct {
	Id       int       `json:"id"`
	VKId     string    `json:"vk_id"`
	Author   string    `json:"author"`
	Text     string    `json:"text"`
	AuthorId string    `json:"author_id"`
	GroupId  string    `json:"groupId"`
	Hash     string    `json:"hash"`
	Status   bool      `json:"status"`
	Date     time.Time `json:"date"`
	FromId   int       `json:"from_id"`
	PostUrl  string    `json:"post_url"`
}
