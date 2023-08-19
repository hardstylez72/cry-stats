package history

import (
	"context"

	"github.com/Masterminds/squirrel"
)

func (r *Repository) MarkAsProcessed(ctx context.Context, id string) (err error) {
	sQuery := squirrel.Update("tx_history").
		Set("processed_at", squirrel.Expr("now()")).
		Where(squirrel.Eq{
			"id": id,
		})

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
