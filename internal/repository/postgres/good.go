package postgres

import (
	"awesomeProject/internal/domain/models"
	"context"
	"log/slog"
)

func (r *Repository) SaveGood(ctx context.Context, name string, projectID int) (models.Goods, error) {
	const op = "repository.save"

	log := r.log.With(
		slog.String("op", op),
	)

	query := `
		INSERT INTO Goods (project_id, name, description, priority, removed, created_at)
		VALUES ($1, $2, '', 0, false, NOW())
		RETURNING id, project_id, name, description, priority, removed, created_at
	`

	row := r.db.QueryRowContext(ctx, query, projectID, name)

	good := models.Goods{}

	err := row.Scan(
		&good.ID,
		&good.ProjectID,
		&good.Name,
		&good.Description,
		&good.Priority,
		&good.Removed,
		&good.CreatedAt,
	)

	if err != nil {
		log.Error("row scan error", "save error")
		return models.Goods{}, err
	}

	log.Info("save successfully")

	return good, nil
}

func (r *Repository) ChangeDescription(
	ctx context.Context,
	name string,
	description string,
	id int,
	projectID int,
) (models.Goods, error) {
	const op = "repository.changedescription"

	log := r.log.With(
		slog.String("op", op),
	)

	query := `
		UPDATE Goods 
		SET description = $1
		WHERE id = $2 AND project_id = $3
		RETURNING id, project_id, name, description, priority, removed, created_at
	`

	row := r.db.QueryRowContext(ctx, query, description, id, projectID)

	var good models.Goods

	err := row.Scan(
		&good.ID,
		&good.ProjectID,
		&good.Name,
		&good.Description,
		&good.Priority,
		&good.Removed,
		&good.CreatedAt,
	)

	if err != nil {
		log.Error("row scan error", "change desc error")
		return models.Goods{}, err
	}

	log.Info("change successfully")

	return good, nil
}

func (r *Repository) RedistributePriorities(
	ctx context.Context,
	newPriority int,
	id int,
	projectID int,
) ([]models.Priorities, error) {
	const op = "repository.redistributepriority"

	log := r.log.With(
		slog.String("op", op),
	)

	updateQuery := `
		UPDATE Goods 
		SET priority = $1
		WHERE project_id = $2 AND priority <= $3
		RETURNING id, priority
	`

	rows, err := r.db.QueryContext(ctx, updateQuery, newPriority, projectID, id)
	if err != nil {
		log.Error("row scan error", "redistribute error")
		return nil, err
	}
	defer rows.Close()

	var priorities []models.Priorities

	for rows.Next() {
		var p models.Priorities
		if err := rows.Scan(&p.ID, &p.PriorityID); err != nil {
			return nil, err
		}
		priorities = append(priorities, p)
	}
	if err := rows.Err(); err != nil {
		log.Error("rows error", "redistribute error")
		return nil, err
	}

	log.Info("redistribute successfully")

	return priorities, nil
}

func (r *Repository) Remove(ctx context.Context, id int, projectID int) (models.Projects, error) {
	const op = "repository.remove"

	log := r.log.With(
		slog.String("op", op),
	)

	updateQuery := `
		UPDATE Goods 
		SET removed = true
		WHERE id = $1 AND project_id = $2
	`

	_, err := r.db.ExecContext(ctx, updateQuery, id, projectID)
	if err != nil {
		log.Error("exec context err", "get error")
		return models.Projects{}, err
	}

	selectQuery := `
		SELECT id, name, created_at
		FROM Projects
		WHERE id = $1
	`

	row := r.db.QueryRowContext(ctx, selectQuery, projectID)

	project := models.Projects{}

	err = row.Scan(
		&project.ID,
		&project.Name,
		&project.CreatedAt,
	)

	if err != nil {
		log.Error("row scan error", "get error")
		return models.Projects{}, err
	}

	log.Info("remove successfully")

	return project, nil
}
