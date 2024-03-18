package handler

import (
	"github.com/gin-gonic/gin"
	"log/slog"
)

type Handler struct {
	log *slog.Logger
}

func (h *Handler) Init() *gin.Engine {
	const op = "handler.init"

	log := h.log.With(
		slog.String("op", op),
	)

	router := gin.New()

	log.Info("Handler init")

	return router
}
