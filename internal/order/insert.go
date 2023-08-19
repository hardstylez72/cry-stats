package order

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type InsertOrderReq struct {
	AccountId      string  `db:"account_id"`
	Net            string  `db:"net"`
	IncomeExpected float64 `db:"income_expected"`
	Addr           string  `db:"addr"`
	Meta           string  `db:"meta"`
}

type insertOrderReqWithId struct {
	InsertOrderReq
	Id     string `db:"id"`
	Status string `db:"status"`
}

type InsertOrderRes struct {
	Id string
}

func (r *Repository) InsertOrder(ctx context.Context, req *InsertOrderReq) (resp *InsertOrderRes, err error) {
	arg := insertOrderReqWithId{
		InsertOrderReq: *req,
		Id:             uuid.NewString(),
		Status:         StatusCreated,
	}

	query := `insert into orders (
		id,
		account_id,
		net,
		addr,
		income_expected,
		status,
		meta
	) values (
		:id,
		:account_id,
		:net,
		:addr,
		:income_expected,
		:status,
		:meta
	);`

	res, err := r.conn.NamedExecContext(ctx, query, arg)
	if err != nil {
		return nil, err
	}

	if rows, err := res.RowsAffected(); err != nil {
		return nil, err
	} else if rows == 0 {
		return nil, errors.New("Order was not created")
	}

	return &InsertOrderRes{
		Id: arg.Id,
	}, nil
}
