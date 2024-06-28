package data

import (
	"database/sql"
	"time"
)

type Task struct {
	ID          int64     `json:"id"`
	ActivatedAt time.Time `json:"activation_time" example:"2024-06-28"`
	Title       string    `json:"title" example:"simple title"`
	Done        bool      `json:"done" example:"true"`
}

type TaskModel struct {
	DB *sql.DB
}
