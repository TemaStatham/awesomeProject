package service

import (
	"awesomeProject/internal/domain/models"
	"context"
	"log/slog"
)

func (s *Service) GetList(ctx context.Context, limit int, offset int) (models.List, error) {
	const op = "service.getlist"

	log := s.log.With(
		slog.String("op", op),
	)

	l, err := s.GoodGetter.GetList(ctx, limit, offset)
	if err != nil {
		log.Error("get list error")
		return models.List{}, err
	}

	log.Info("service get list successfully")

	return l, nil
}
