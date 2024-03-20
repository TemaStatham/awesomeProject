package clickhouse

import (
	"awesomeProject/internal/domain/models"
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"log/slog"
)

type Repository struct {
	log *slog.Logger
	ch  clickhouse.Conn
}

func NewClickhouse(ch clickhouse.Conn, l *slog.Logger) *Repository {
	return &Repository{
		ch:  ch,
		log: l,
	}
}

func (r *Repository) Log(ctx context.Context, goods models.Goods) error {
	query := fmt.Sprintf("INSERT INTO goods (Id, ProjectId, Name, Description, Priority, Removed, EventTime) VALUES (%d, %d, '%s', '%s', %d, %t, '%s')",
		goods.ID, goods.ProjectID, goods.Name, goods.Description, goods.Priority, goods.Removed, goods.CreatedAt.Format("2006-01-02 15:04:05"))

	err := r.ch.Exec(ctx, query)
	if err != nil {
		r.log.Info("Error inserting data into ClickHouse:", err)
		return err
	}

	return nil
}
