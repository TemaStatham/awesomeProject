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

	log.Info("service get list successfully")

	return models.List{}, nil
}
