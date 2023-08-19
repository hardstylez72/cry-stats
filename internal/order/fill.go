package order

import (
	"context"

	"github.com/Masterminds/squirrel"
)

type FillReq struct {
	ConfirmedTxId  string
	Meta           string
	ExchangeRate   int
	IncomeReceived float64
}

func (r *Repository) Fill(ctx context.Context, id string, req *FillReq) (err error) {
	sQuery := squirrel.Update("orders").
		SetMap(map[string]interface{}{
			"confirmed_by_tx_id": req.ConfirmedTxId,
			"meta":               req.Meta,
			"exchange_rate":      req.ExchangeRate,
			"income_received":    req.IncomeReceived,
		}).
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
