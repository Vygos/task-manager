package domain

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id        uuid.UUID  `json:"id"`
	Title     string     `json:"title"`
	Status    TaskStatus `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func (t *Task) Merge(task Task) {

	if task.Status != "" {
		t.Status = task.Status
	}
	if task.Title != "" {
		t.Title = task.Title
	}

}
