package handler

import (
	"github.com/gin-gonic/gin"
	"log/slog"
)

type Handler struct {
	log *slog.Logger
}

func NewHandler(_log *slog.Logger) *Handler {
	return &Handler{
		log: _log,
	}
}

type GoodCreator interface {
	Create()
}

type GoodUpdater interface {
	Update()
	Reprioritize()
}

type GoodRemover interface {
	Remove()
}

type ListProvider interface {
	GetList()
}

func (h *Handler) Init() *gin.Engine {
	const op = "handler.init"

	log := h.log.With(
		slog.String("op", op),
	)

	router := gin.New()

	good := router.Group("/good")
	{
		good.POST("/create")
		good.PATCH("/update")
		good.DELETE("/remove")
		good.PATCH("/reprioritiize")
	}
	goods := router.Group("/goods")
	{
		goods.GET("/list")
	}

	log.Info("Handler init")

	return router
}
