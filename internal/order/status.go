package order

import (
	"context"

	"github.com/Masterminds/squirrel"
)

const StatusCreated = "Created"
const StatusInProgress = "InProgress"
const StatusProcessed = "Processed"
const StatusError = "Error"

func (r *Repository) SetStatus(ctx context.Context, id string, status string) (err error) {
	sQuery := squirrel.Update("orders").
		Set("status", status).
		Where(squirrel.Eq{
			"id": id,
		})

	if status == StatusProcessed {
		sQuery = sQuery.Set("confirmed_at", squirrel.Expr("now()"))
	}

	q, a, err := sQuery.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return err
	}

	_, err = r.conn.ExecContext(ctx, q, a...)
	if err != nil {
		return err
	}

	return nil
}
