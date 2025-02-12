package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/vygos/task/task-api/internal/task/domain"
	"github.com/vygos/task/task-api/pkg/pagination"
)

type Service interface {
	SaveTask(context.Context, domain.Task) (*domain.Task, error)
	GetAll(ctx context.Context, page pagination.Page) (*pagination.Page, error)
	UpdateTask(ctx context.Context, task domain.Task) (*domain.Task, error)
	DeleteTask(ctx context.Context, id uuid.UUID) error
}
