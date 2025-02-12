package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	"github.com/gin-gonic/gin"

	"github.com/vygos/task/task-api/internal/config"
	"github.com/vygos/task/task-api/internal/db"
	"github.com/vygos/task/task-api/internal/task"
	"github.com/vygos/task/task-api/pkg/middleware"
)

func main() {
	g := gin.Default()

	cfg, err := config.NewConfig("dev.yaml")
	if err != nil {
		log.Fatal(err)
	}

	if cfg.Env == "dev" {
		g.Use(middleware.CORS)
	}

	database, err := db.NewDatabase(config.DB{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		User:     cfg.DB.User,
		Pass:     cfg.DB.Pass,
		Name:     cfg.DB.Name,
		MaxConns: cfg.DB.MaxConns,
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = database.Exec(context.Background(), db.InitDB)
	if err != nil {
		slog.Error("error while initializing database ", slog.Any("error", err))
	}

	task.NewApi(g, database)

	err = g.Run(fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatal(err)
	}
}
