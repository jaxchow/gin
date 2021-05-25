package Todo

import (
	"time"
)

type Todo struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	Author    string `json:"author"`
	Owner     string `json:"owner"`
	Priority  string `json:"priority"`
	CreatedAt time.Time
	Duration  int64 `json:"duration"`
	StartTime int64 `json:"startTime"`
	TaskId    uint  `json:"taskId"`
	// gorm.Model
	// Task Task `gorm:"foreignKey:taskId"`
}
