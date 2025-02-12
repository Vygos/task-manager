package repository

import (
	"github.com/vygos/task/task-api/internal/task/domain"
)

func (t *Task) toDomain() *domain.Task {
	return &domain.Task{
		Id:        t.ID.Bytes,
		Status:    domain.TaskStatus(t.Status),
		Title:     t.Title,
		CreatedAt: t.CreatedAt.Time,
		UpdatedAt: t.UpdatedAt.Time,
	}
}
