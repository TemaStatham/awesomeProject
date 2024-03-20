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
	GoodCacheSaver
	GoodCacheRemover
	GoodCacheGetter
	Logger
}

type GoodSaver interface {
	SaveGood(ctx context.Context, name string, projectID int) (models.Goods, error)
}

type GoodUpdater interface {
	ChangeDescription(ctx context.Context, name string, description string, id int, projectID int) (models.Goods, error)
	RedistributePriorities(ctx context.Context, newPriority int, id int, projectID int) ([]models.Priorities, error)
}

type GoodRemover interface {
	Remove(ctx context.Context, id int, projectID int) (models.Projects, error)
}

type GoodGetter interface {
	GetList(ctx context.Context, limit int, offset int) (models.List, error)
}

type GoodCacheSaver interface {
	Cache(ctx context.Context, name string, projectID int) (models.Goods, error)
}

type GoodCacheRemover interface {
	UpdateToRemove(ctx context.Context, id int, projectID int) (models.Projects, error)
	Delete(ctx context.Context, id int, projectID int) error
}

type GoodCacheGetter interface {
	Get(ctx context.Context, id int, projectID int) (models.Goods, error)
}

type Logger interface {
	Log(ctx context.Context, goods models.Goods) error
}

func NewService(
	_l *slog.Logger,
	c GoodSaver,
	u GoodUpdater,
	r GoodRemover,
	g GoodGetter,
	s GoodCacheSaver,
	cr GoodCacheRemover,
	cg GoodCacheGetter,
	l Logger,
) *Service {
	return &Service{
		log:              _l,
		GoodSaver:        c,
		GoodUpdater:      u,
		GoodRemover:      r,
		GoodGetter:       g,
		GoodCacheGetter:  cg,
		GoodCacheRemover: cr,
		GoodCacheSaver:   s,
		Logger:           l,
	}
}
