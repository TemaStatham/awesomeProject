package service

import (
	"awesomeProject/internal/domain/models"
	"context"
	"log/slog"
)

type Service struct {
	log *slog.Logger
	GoodSaver
	GoodUpdater
	GoodRemover
	GoodGetter
}

type GoodSaver interface {
	SaveGood(ctx context.Context, name string, projectID int) (int, error)
}

type GoodUpdater interface {
	ChangeDescription(ctx context.Context, name string, description string, id int, projectID int) (int, error)
	RedistributePriorities(ctx context.Context, newPriority int, id int, projectID int) ([]models.Priorities, error)
}

type GoodRemover interface {
	Remove(ctx context.Context, id int, projectID int) (models.Projects, error)
}

type GoodGetter interface {
	GetGoodByID(ctx context.Context, id int) (models.Goods, error)
	GetList(ctx context.Context, limit int, offset int) (models.List, error)
}

func NewService(_l *slog.Logger, c GoodSaver, u GoodUpdater, r GoodRemover, g GoodGetter) *Service {
	return &Service{
		log:         _l,
		GoodSaver:   c,
		GoodUpdater: u,
		GoodRemover: r,
		GoodGetter:  g,
	}
}
