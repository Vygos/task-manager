package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/vygos/task/task-api/internal/task/domain"
	"github.com/vygos/task/task-api/pkg/pagination"
)

func (q *Queries) Save(ctx context.Context, task domain.Task) (*domain.Task, error) {
	result, err := q.saveTask(ctx, saveTaskParams{
		ID:     pgtype.UUID{Valid: true, Bytes: uuid.New()},
		Title:  task.Title,
		Status: Status(task.Status),
	})
	if err != nil {
		return nil, err
	}

	return result.toDomain(), nil
}

func (q *Queries) GetByID(ctx context.Context, taskId uuid.UUID) (*domain.Task, error) {
	task, err := q.getByID(ctx, pgtype.UUID{Valid: true, Bytes: taskId})
	if err != nil {
		return nil, err
	}

	return task.toDomain(), nil
}

func (q *Queries) GetAll(ctx context.Context, page pagination.Page) ([]domain.Task, error) {
	tasks, err := q.getAll(ctx, getAllParams{
		Offset: int32((page.Page - 1) * page.Size),
		Limit:  int32(page.Size),
	})
	if err != nil {
		return nil, err
	}

	tasksDomain := make([]domain.Task, len(tasks))

	for i, t := range tasks {
		tasksDomain[i] = *t.toDomain()
	}

	return tasksDomain, nil
}

func (q *Queries) Update(ctx context.Context, task domain.Task) (*domain.Task, error) {
	updatedTask, err := q.updateTask(ctx, updateTaskParams{
		ID:     pgtype.UUID{Valid: true, Bytes: task.Id},
		Title:  task.Title,
		Status: Status(task.Status),
	})
	if err != nil {
		return nil, err
	}

	return updatedTask.toDomain(), nil
}

func (q *Queries) Delete(ctx context.Context, taskId uuid.UUID) error {
	_, err := q.deleteTask(ctx, pgtype.UUID{Valid: true, Bytes: taskId})
	if err != nil {
		return err
	}
	return err
}
