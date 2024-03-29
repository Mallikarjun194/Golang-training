package model

import "time"

type Blog struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Author    string    `json:"author" binding:"required"`
	Title     string    `json:"title" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
}

type Error struct {
	Msg string `json:"msg"`
}

type Msg struct {
	Message string `json:"message"`
}

type Result struct {
	Data        []Blog `json:"data"`
	Status      bool   `json:"status"`
	NoOfRecords int    `json:"noOfRecords"`
}
