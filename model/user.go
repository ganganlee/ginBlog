package model

import "time"

//用户表
type User struct {
	Id         int64     `json:"id"`
	Secret     string    `json:"secret" xorm:"varchar(50) notnull unique"` //用户密钥
	Username   string    `json:"username" xorm:"varchar(125) index(user_login) notnull unique"`
	Password   string    `json:"password" xorm:"varchar(125) index(user_login) notnull"`
	Avatar     string    `json:"avatar" xorm:"varchar(125)"` //用户头像
	Email      string    `json:"email" xorm:"varchar(125) unique"`
	CreateTime time.Time `json:"create_time" xorm:"created"`
	UpdateTime time.Time `json:"update_time" xorm:"updated"`
}
