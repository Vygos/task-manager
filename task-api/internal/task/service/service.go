package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"golang.org/x/sync/errgroup"

	"github.com/vygos/task/task-api/internal/task/domain"
	"github.com/vygos/task/task-api/pkg/pagination"
)

type TaskRepository interface {
	Save(ctx context.Context, task domain.Task) (*domain.Task, error)
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Task, error)
	GetAll(ctx context.Context, page pagination.Page) ([]domain.Task, error)
	GetCount(context.Context) (int64, error)
	Update(ctx context.Context, task domain.Task) (*domain.Task, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type TaskService struct {
	repository TaskRepository
}

func NewService(taskRepository TaskRepository) *TaskService {
	return &TaskService{
		repository: taskRepository,
	}
}

func (t *TaskService) GetAll(ctx context.Context, page pagination.Page) (*pagination.Page, error) {
	g := errgroup.Group{}

	cTasks := make(chan []domain.Task, 1)

	g.Go(func() error {
		defer close(cTasks)
		results, err := t.repository.GetAll(ctx, page)
		if err != nil {
			return err
		}
		cTasks <- results
		return nil
	})

	count, err := t.repository.GetCount(ctx)
	if err != nil {
		return nil, err
	}

	if err = g.Wait(); err != nil {
		return nil, err
	}

	page.Data, page.TotalElements = <-cTasks, int(count)

	return &page, nil
}

func (t *TaskService) SaveTask(ctx context.Context, task domain.Task) (*domain.Task, error) {
	return t.repository.Save(ctx, task)
}

func (t *TaskService) UpdateTask(ctx context.Context, task domain.Task) (*domain.Task, error) {
	taskPersisted, err := t.repository.GetByID(ctx, task.Id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, TaskNotFoundErr
		}
		return nil, err
	}
	taskPersisted.Merge(task)

	result, err := t.repository.Update(ctx, *taskPersisted)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t *TaskService) DeleteTask(ctx context.Context, id uuid.UUID) error {
	return t.repository.Delete(ctx, id)
}
