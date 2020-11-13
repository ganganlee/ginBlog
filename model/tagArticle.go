package model

import "time"

//标签文章表
type TagArticle struct {
	Id         int64     `json:"id"`
	TagId      int64     `json:"tag_id" xorm:"notnull unique(tag_article)"`
	ArticleId  int64     `json:"article_id" xorm:" notnull unique(tag_article)"`
	CreateTime time.Time `json:"create_time" xorm:"created"`
	UpdateTime time.Time `json:"update_time" xorm:"updated"`
}
