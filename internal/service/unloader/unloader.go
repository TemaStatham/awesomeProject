package unloader

import (
	"awesomeProject/internal/domain/models"
	"context"
	"fmt"
	"log/slog"
	"time"
)

type Unloader struct {
	log *slog.Logger
	UnloaderFromDB
	LoaderToDB
	Logger
}

type UnloaderFromDB interface {
	GetAll(ctx context.Context) ([]models.Goods, []string, error)
	DeleteAll(ctx context.Context, keys []string)
}

type LoaderToDB interface {
	SaveGood(ctx context.Context, name string, projectID int) (models.Goods, error)
}

type Logger interface {
	Log(ctx context.Context, goods models.Goods) error
}

func NewUnloader(l *slog.Logger, u UnloaderFromDB, _l LoaderToDB, logger Logger) *Unloader {
	return &Unloader{
		log:            l,
		UnloaderFromDB: u,
		LoaderToDB:     _l,
		Logger:         logger,
	}
}

func (u *Unloader) UnloadOnceWhile(ctx context.Context, t time.Duration) {
	const op = "unloader.Unload"

	log := u.log.With(
		slog.String("op", op),
	)

	log.Info("unload product")

	ticker := time.NewTicker(t)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopping product data unload...")
			return
		case <-ticker.C:
			log.Info("start unload product")

			gs, ks, err := u.GetAll(ctx)
			if err != nil {
				log.Error("db is empty")
				continue
			}

			u.DeleteAll(ctx, ks)

			for _, g := range gs {
				_g, err := u.SaveGood(ctx, g.Name, g.ProjectID)
				if err != nil {
					log.Error("save to db error")
					continue
				}

				err = u.Log(ctx, _g)
				if err != nil {
					log.Error("log error")
					continue
				}
			}

			log.Info("unload successfully")
		default:
			continue
		}
	}
}
