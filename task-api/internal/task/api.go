package task

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	repository "github.com/vygos/task/task-api/internal/task/repository/queries"
	"github.com/vygos/task/task-api/internal/task/service"
	"github.com/vygos/task/task-api/pkg/middleware"
)

type Api struct {
	gin gin.Engine
	h   Handler
}

func NewApi(engine *gin.Engine, db *pgxpool.Pool) {

	taskRepository := repository.New(db)
	taskService := service.NewService(taskRepository)
	h := NewHandler(taskService)

	r := engine.Group("tasks")

	r.POST("", middleware.ErrorHandler(h.CreateTask))
	r.GET("", middleware.ErrorHandler(h.GetAll))
	r.PATCH(":taskId", middleware.ErrorHandler(h.UpdateTask))
	r.DELETE(":taskId", middleware.ErrorHandler(h.DeleteTask))
}
