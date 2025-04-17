package models

import "time"

type Todo struct{
	ID uint `gorm:"primaryKey" json:"id"`
	Title string `json:"title"`
	IsDone bool `json:"is_done"`
	CreatedAt time.Time `json:"created_at"`
}