package task

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/vygos/task/task-api/internal/task/domain"
)

type CreateTaskInput struct {
	Title  string            `json:"title"`
	Status domain.TaskStatus `json:"status"`
}

func (c *CreateTaskInput) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Title, validation.Required),
		validation.Field(&c.Status, validation.Required, validation.In(domain.Completed, domain.Incomplete)),
	)
}

type UpdateTaskInput struct {
	Title  string            `json:"title"`
	Status domain.TaskStatus `json:"status"`
}

func (c *UpdateTaskInput) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Status, validation.In(domain.Completed, domain.Incomplete)),
	)
}
