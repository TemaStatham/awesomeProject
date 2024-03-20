package service

import (
	"awesomeProject/internal/domain/models"
	"context"
	"errors"
	"log/slog"
)

var (
	ErrNameIsEmpty = errors.New("name is empty")
	ErrInternal    = errors.New("internal error")
)

func (s *Service) Create(ctx context.Context, name string, projectID int) (models.Goods, error) {
	const op = "service.create"

	log := s.log.With(
		slog.String("op", op),
	)

	if name == "" {
		log.Error("name is null", ErrNameIsEmpty.Error())
		return models.Goods{}, ErrNameIsEmpty
	}

	id, err := s.SaveGood(ctx, name, projectID)
	if err != nil {
		log.Error("good is not saved", ErrInternal.Error())
		return models.Goods{}, ErrInternal
	}

	g, err := s.GetGoodByID(ctx, id)
	if err != nil {
		log.Error("good is not getted", ErrInternal.Error())
		return models.Goods{}, ErrInternal
	}

	log.Info("service create successfully")

	return g, nil
}

func (s *Service) Update(
	ctx context.Context,
	name string,
	description string,
	id int,
	projectID int,
) (models.Goods, error) {
	const op = "service.update"

	log := s.log.With(
		slog.String("op", op),
	)

	if name == "" {
		log.Error("name is null", ErrNameIsEmpty.Error())
		return models.Goods{}, ErrNameIsEmpty
	}

	newID, err := s.ChangeDescription(ctx, name, description, id, projectID)
	if err != nil {
		log.Error("good is not update", ErrInternal.Error())
		return models.Goods{}, ErrInternal
	}

	g, err := s.GetGoodByID(ctx, newID)
	if err != nil {
		log.Error("good is not getted", ErrInternal.Error())
		return models.Goods{}, ErrInternal
	}

	log.Info("service update successfully")

	return g, nil
}

func (s *Service) Reprioritize(
	ctx context.Context,
	newPriority int,
	id int,
	projectID int,
) ([]models.Priorities, error) {
	const op = "service.reprioritiize"

	log := s.log.With(
		slog.String("op", op),
	)

	p, err := s.RedistributePriorities(ctx, newPriority, id, projectID)
	if err != nil {
		log.Error("reprioritiize is not allow", ErrInternal.Error())
		return []models.Priorities{}, ErrInternal
	}

	log.Info("service reprioritiize successfully")

	return p, nil
}

func (s *Service) Remove(ctx context.Context, id int, projectID int) (models.Projects, error) {
	const op = "service.remove"

	log := s.log.With(
		slog.String("op", op),
	)
	p, err := s.Remove(ctx, id, projectID)
	if err != nil {
		log.Error("remove error", ErrInternal.Error())
		return models.Projects{}, ErrInternal
	}

	log.Info("service remove successfully")

	return p, nil
}
