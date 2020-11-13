package model

import "time"

//文章表
type Article struct {
	Id         int64     `json:"id"`
	UserId     int64     `json:"user_id" xorm:"notnull"`
	Title      string    `json:"title" xorm:"notnull"`
	Synopsis   string    `json:"synopsis" xorm:"notnull"` //简介
	Content    string    `json:"content" xorm:"text notnull"`
	Tag        []int64   `json:"tag" xorm:"notnull"`
	TagStr     []string  `json:"tag_str" xorm:"notnull"` //冗余文章标签字符串，防止连表查询
	Sort       int8      `json:"sort" xorm:"default 1"`
	View       int64     `json:"view" xorm:"default 0"`
	Status     int8      `json:"status" xorm:"default 1"`
	Comment    int64     `json:"comment" xorm:"default 0"`
	CreateTime time.Time `json:"create_time" xorm:"created"`
	UpdateTime time.Time `json:"update_time" xorm:"updated"`
}
