package handler

import (
	"net/http"

	"github.com/NordSec/status-checker-go"
	"github.com/labstack/echo/v4"
)

type EchoHandler struct {
	checkerPool status.Pool
}

func NewEchoHandler(pool status.Pool) *EchoHandler {
	return &EchoHandler{checkerPool: pool}
}

func (ec *EchoHandler) Status(context echo.Context) error {
	return context.JSON(http.StatusOK, echo.Map{"status": ec.checkerPool.Status()})
}

func (ec *EchoHandler) Details(context echo.Context) error {
	return context.JSON(http.StatusOK, ec.checkerPool.Details())
}
