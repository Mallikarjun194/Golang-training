package model

import "time"

type Blog struct {
	ID        string    `json:"id"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

type Error struct {
	Msg string `json:"msg"`
}

type Msg struct {
	Message string `json:"message"`
}
