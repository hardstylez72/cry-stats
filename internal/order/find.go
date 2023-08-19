package order

import (
	"context"
	"database/sql"
	"time"

	"github.com/Masterminds/squirrel"
)

type FindOrderReq struct {
	Addr           string
	Net            string
	IncomeExpected float64
	CreatedSince   *time.Time
}

type FindOrderRes struct {
	Id        string `db:"id"`
	Status    string `db:"status"`
	AccountId string `db:"account_id"`
}

func (r *Repository) FindOrder(ctx context.Context, req FindOrderReq) (resp *FindOrderRes, isFound bool, err error) {
	sQuery := squirrel.Select(
		"id",
		"status",
		"account_id",
	).
		From("orders").
		Where(
			squirrel.Eq{
				"net":             req.Net,
				"income_expected": req.IncomeExpected,
			},
		)

	if req.CreatedSince != nil {
		sQuery = sQuery.Where("created_at > ?", req.CreatedSince)
	}

	q, a, err := sQuery.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, false, err
	}

	res := FindOrderRes{}
	err = r.conn.GetContext(ctx, &res, q, a...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false, nil
		}
		return nil, false, err
	}

	return &res, true, nil
}
