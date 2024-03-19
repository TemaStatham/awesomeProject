package handler

import (
	"awesomeProject/internal/domain/models"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
)

const (
	InvalidInputBodyErr = "invalid input body"
)

type Handler struct {
	log *slog.Logger
	GoodCreator
	GoodUpdater
	GoodRemover
	ListProvider
}

func NewHandler(_log *slog.Logger) *Handler {
	return &Handler{
		log: _log,
	}
}

type GoodCreator interface {
	Create(ctx context.Context, name string, projectID int) (models.Goods, error)
}

type GoodUpdater interface {
	Update(ctx context.Context, name string, description string, id int, projectID int) (models.Goods, error)
	Reprioritize(ctx context.Context, newPriority int, id int, projectID int) ([]models.Priorities, error)
}

type GoodRemover interface {
	Remove(ctx context.Context, id int, projectID int) (models.Projects, error)
}

type ListProvider interface {
	GetList(ctx context.Context, limit int, offset int) (models.List, error)
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
