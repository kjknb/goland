package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	UserId     int64
	TageId     int64
	Type       int
	Media      int
	Content    string
	CreateTimr uint64
	ReadTime   uint64
	Pic        string
	Url        string
	Desc       string
	Amount     int
}

func (table *Message) TableName() string {
	return "message"

}
