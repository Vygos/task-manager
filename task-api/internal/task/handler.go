package task

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/vygos/task/task-api/internal/task/domain"
	"github.com/vygos/task/task-api/internal/task/service"
	"github.com/vygos/task/task-api/pkg/middleware/statuserr"
	"github.com/vygos/task/task-api/pkg/pagination"
	"github.com/vygos/task/task-api/pkg/rest"
)

type Handler struct {
	ts service.Service
}

func NewHandler(ts service.Service) *Handler {
	return &Handler{
		ts: ts,
	}
}

func (h *Handler) CreateTask(gCtx *gin.Context) error {
	ctx, cancel := context.WithTimeout(gCtx.Request.Context(), 10*time.Second)
	defer cancel()

	var task CreateTaskInput

	if err := gCtx.ShouldBindJSON(&task); err != nil {
		return statuserr.NewBadRequest(fmt.Sprintf("invalid request body: %v", err))
	}

	if err := task.Validate(); err != nil {
		return statuserr.NewBadRequest(err.Error())
	}

	result, err := h.ts.SaveTask(ctx, domain.Task{Title: task.Title, Status: task.Status})
	if err != nil {
		slog.Error("[SaveTask] error: ", slog.Any("error", err))
		return err
	}

	gCtx.JSON(http.StatusCreated, result)

	return nil
}

func (h *Handler) GetAll(gCtx *gin.Context) error {
	ctx, cancel := context.WithTimeout(gCtx.Request.Context(), 10*time.Second)
	defer cancel()

	size, _ := rest.GetInt(gCtx, "size")
	page, _ := rest.GetInt(gCtx, "page")

	newPage := pagination.NewPage(page, size, 10)

	results, err := h.ts.GetAll(ctx, newPage)
	if err != nil {
		slog.Error("[GetAll] error: ", slog.Any("error", err))
		return err
	}

	gCtx.JSON(http.StatusOK, results)
	return nil
}

func (h *Handler) UpdateTask(gCtx *gin.Context) error {
	ctx, cancel := context.WithTimeout(gCtx.Request.Context(), 10*time.Second)
	defer cancel()

	taskId, err := rest.GetUUID(gCtx, "taskId")
	if err != nil {
		return statuserr.NewBadRequest(err.Error())
	}

	var task UpdateTaskInput
	if err = gCtx.ShouldBindJSON(&task); err != nil {
		return statuserr.NewBadRequest(fmt.Sprintf("invalid request body: %v", err))
	}

	if err = task.Validate(); err != nil {
		return statuserr.NewBadRequest(err.Error())
	}

	result, err := h.ts.UpdateTask(ctx, domain.Task{
		Id:     taskId,
		Title:  task.Title,
		Status: task.Status,
	})
	if err != nil {
		if errors.Is(err, service.TaskNotFoundErr) {
			return statuserr.NewNotFoundErr(err.Error())
		}
		slog.Error("[UpdateTask] error: ", slog.Any("error", err))
		return err
	}

	gCtx.JSON(http.StatusOK, result)

	return nil
}

func (h *Handler) DeleteTask(gCtx *gin.Context) error {
	ctx, cancel := context.WithTimeout(gCtx.Request.Context(), 10*time.Second)
	defer cancel()

	taskId, err := rest.GetUUID(gCtx, "taskId")
	if err != nil {
		return statuserr.NewBadRequest(err.Error())
	}

	return h.ts.DeleteTask(ctx, taskId)

}
