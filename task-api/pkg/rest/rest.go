package rest

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetInt(gCtx *gin.Context, value string) (int, error) {
	size, err := strconv.Atoi(gCtx.Query(value))
	if err != nil {
		return 0, fmt.Errorf("could not convert value of %s to int", value)
	}

	return size, nil
}

func GetUUID(gCtx *gin.Context, value string) (uuid.UUID, error) {
	return uuid.Parse(gCtx.Param(value))
}
