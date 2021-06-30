package handler

import (
	"net/http"

	"github.com/NordSec/status-checker-go"
	"github.com/gin-gonic/gin"
)

type GinHandler struct {
	checkerPool status.Pool
}

func NewGinHandler(pool status.Pool) *GinHandler {
	return &GinHandler{checkerPool: pool}
}

func (c *GinHandler) Status(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"status": c.checkerPool.Status()})
}

func (c *GinHandler) Details(context *gin.Context) {
	context.JSON(http.StatusOK, c.checkerPool.Details())
}
