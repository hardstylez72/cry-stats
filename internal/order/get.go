package order

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
)

type GetOrderReq struct {
	Id string `db:"id"`
}

type GetOrderRes struct {
	Status    string `db:"status"`
	AccountId string `db:"account_id"`
}

func (r *Repository) GetOrder(ctx context.Context, req *GetOrderReq) (resp *GetOrderRes, err error) {
	sQuery := squirrel.Select(
		"status",
		"account_id",
	).
		From("orders").
		Where(
			squirrel.Eq{
				"id": req.Id,
			},
		)

	q, a, err := sQuery.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	res := &GetOrderRes{}
	err = r.conn.GetContext(ctx, res, q, a...)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return res, nil
}

func (r *Repository) AmountUniq(ctx context.Context, am float64) (bool, error) {

	q := `select count(1) from orders where confirmed_at is null and income_expected = $1`

	var count = 0
	err := r.conn.GetContext(ctx, &count, q, am)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, err
	}

	return count == 0, nil
}
