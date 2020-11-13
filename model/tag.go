package model

import "time"

//标签表
type Tag struct {
	Id         int64     `json:"id"`
	UserId     int64     `json:"user_id" xorm:"notnull index unique(userTag)"`
	Name       string    `json:"name" xorm:"notnull index unique(userTag)"`
	Status     int8      `json:"status" xorm:"default 1"`
	CreateTime time.Time `json:"create_time" xorm:"created"`
	UpdateTime time.Time `json:"update_time" xorm:"updated"`
}