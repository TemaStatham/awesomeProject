package postgres

import (
	"awesomeProject/internal/domain/models"
	"context"
	"log/slog"
)

func (r *Repository) GetList(ctx context.Context, limit int, offset int) (models.List, error) {
	const op = "repository.getlist"

	log := r.log.With(
		slog.String("op", op),
	)

	query := `
		SELECT id, project_id, name, description, priority, removed, created_at
		FROM Goods
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return models.List{}, err
	}
	defer rows.Close()

	var list models.List
	list.Goods = make([]models.Goods, 0)

	for rows.Next() {
		var good models.Goods
		if err := rows.Scan(
			&good.ID,
			&good.ProjectID,
			&good.Name,
			&good.Description,
			&good.Priority,
			&good.Removed,
			&good.CreatedAt,
		); err != nil {
			log.Error("row scan error", "get error")
			return models.List{}, err
		}
		list.Goods = append(list.Goods, good)
	}
	if err := rows.Err(); err != nil {
		log.Error("rows error", "get error")
		return models.List{}, err
	}

	countQuery := `
		SELECT COUNT(*)
		FROM Goods
	`
	removedQuery := `
		SELECT COUNT(*)
		FROM Goods
		WHERE removed = true
	`

	var total, removed int
	err = r.db.QueryRowContext(ctx, countQuery).Scan(&total)
	if err != nil {
		log.Error("query row error", "get error")
		return models.List{}, err
	}
	err = r.db.QueryRowContext(ctx, removedQuery).Scan(&removed)
	if err != nil {
		log.Error("query row error", "get error")
		return models.List{}, err
	}

	list.Meta = models.Meta{
		Total:   total,
		Removed: removed,
		Limit:   limit,
		Offset:  offset,
	}

	log.Info("get successfully")

	return list, nil
}
