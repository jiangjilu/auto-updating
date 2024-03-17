package model

import "gorm.io/gorm"

type News struct {
	gorm.Model
	Cid     int64  `json:"cid" column:"cid"`
	Title   string `json:"title" column:"title"`
	Content string `json:"content" column:"content"`
	State   int64  `json:"state" column:"state"`
}

func (u *News) TableName() string {
	return "news"
}
