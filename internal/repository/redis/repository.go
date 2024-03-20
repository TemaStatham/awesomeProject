package redis

import (
	"awesomeProject/internal/domain/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"time"
)

type Repository struct {
	log *slog.Logger
	r   *redis.Client
}

func NewRedis(r *redis.Client, l *slog.Logger) *Repository {
	return &Repository{
		r:   r,
		log: l,
	}
}

func (r *Repository) Cache(ctx context.Context, name string, projectID int) (models.Goods, error) {
	const op = "repository.cache"

	log := r.log.With(
		slog.String("op", op),
	)

	goods := models.Goods{Name: name, ProjectID: projectID}

	goodsJSON, err := json.Marshal(goods)
	if err != nil {
		log.Error("Error marshalling goods to JSON:", err)
		return models.Goods{}, err
	}

	key := fmt.Sprintf("goods:%s:%d", name, projectID)

	err = r.r.Set(ctx, key, goodsJSON, 24*time.Hour).Err()
	if err != nil {
		log.Error("Error caching data in Redis:", err)
		return goods, err
	}

	log.Info("Cached successfully")

	return goods, nil
}

func (r *Repository) Get(ctx context.Context, id int, projectID int) (models.Goods, error) {
	const op = "repository.get"

	log := r.log.With(
		slog.String("op", op),
	)

	key := fmt.Sprintf("goods:%s:%d", id, projectID)

	val, err := r.r.Get(ctx, key).Result()
	if err != nil {
		log.Error("Error getting goods from Redis:", err)
		return models.Goods{}, err
	}

	var goods models.Goods
	err = json.Unmarshal([]byte(val), &goods)
	if err != nil {
		log.Error("Error unmarshalling goods from JSON:", err)
		return models.Goods{}, err
	}

	log.Info("Retrieved successfully")

	return goods, nil
}

func (r *Repository) UpdateToRemove(ctx context.Context, id int, projectID int) (models.Projects, error) {
	const op = "repository.updatetoremove"

	log := r.log.With(
		slog.String("op", op),
	)

	key := fmt.Sprintf("goods:%d:%d", id, projectID)

	val, err := r.r.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			log.Info("Key not found in Redis")
			return models.Projects{}, nil
		}
		log.Error("Error getting goods from Redis:", err)
		return models.Projects{}, err
	}

	var goods models.Goods
	err = json.Unmarshal([]byte(val), &goods)
	if err != nil {
		log.Error("Error unmarshalling goods from JSON:", err)
		return models.Projects{}, err
	}

	goods.Removed = true

	goodsJSON, err := json.Marshal(goods)
	if err != nil {
		log.Error("Error marshalling goods to JSON:", err)
		return models.Projects{}, err
	}

	err = r.r.Set(ctx, key, goodsJSON, 24*time.Hour).Err()
	if err != nil {
		log.Error("Error updating goods in Redis:", err)
		return models.Projects{}, err
	}

	log.Info("Updated to remove successfully")

	return models.Projects{}, nil
}

func (r *Repository) Delete(ctx context.Context, id int, projectID int) error {
	const op = "repository.updatetoremove"

	log := r.log.With(
		slog.String("op", op),
	)

	key := fmt.Sprintf("goods:%d:%d", id, projectID)

	err := r.r.Del(ctx, key).Err()
	if err != nil {
		log.Error("Error deleting goods from Redis:", err)
		return err
	}

	log.Info("Updated to remove successfully")

	return nil
}

func (r *Repository) GetAll(ctx context.Context) ([]models.Goods, []string, error) {
	const op = "repository.getall"

	log := r.log.With(
		slog.String("op", op),
	)

	keys, err := r.r.Keys(ctx, "goods:*").Result()
	if err != nil {
		log.Error("Error retrieving keys from Redis:", err)
		return nil, []string{}, err
	}

	var allGoods []models.Goods

	for _, key := range keys {
		val, err := r.r.Get(ctx, key).Result()
		if err != nil {
			log.Error("Error getting goods from Redis:", err)
			continue
		}

		var goods models.Goods
		err = json.Unmarshal([]byte(val), &goods)
		if err != nil {
			log.Error("Error unmarshalling goods from JSON:", err)
			continue
		}

		allGoods = append(allGoods, goods)
	}

	log.Info("Retrieved successfully")

	return allGoods, keys, nil
}

func (r *Repository) DeleteAll(ctx context.Context, keys []string) {
	for _, key := range keys {
		if err := r.r.Del(ctx, key).Err(); err != nil {
			r.log.Error("Error deleting goods from Redis:", err)
			continue
		}
	}
}
