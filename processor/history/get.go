package history

import (
	"context"

	"github.com/Masterminds/squirrel"
)

type GetUnprocessedRes struct {
	Id             string `db:"id"`
	Addr           string `db:"addr"`
	Net            string `db:"net"`
	IncomeReceived int    `db:"income_received"`
}

func (r *Repository) GetUnprocessed(ctx context.Context) ([]GetUnprocessedRes, error) {
	sQuery := squirrel.Select("id", "addr", "net", "income_received").
		From("tx_history").
		Where("processed_at IS NULL")

	q, a, err := sQuery.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var res []GetUnprocessedRes
	err = r.conn.SelectContext(ctx, &res, q, a...)
	if err != nil {
		return nil, err
	}

	return res, nil
}
