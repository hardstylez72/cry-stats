package history

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type AppendTxReq struct {
	Addr           string
	Net            string
	Hash           string
	IncomeReceived int
	IncomeTime     time.Time
}

func (r *Repository) Append(ctx context.Context, req AppendTxReq) (id string, err error) {
	sQuery := squirrel.Insert("tx_history").
		SetMap(map[string]interface{}{
			"id":              uuid.NewString(),
			"net":             req.Net,
			"addr":            req.Addr,
			"tx_hash":         req.Hash,
			"income_received": req.IncomeReceived,
			"created_at":      req.IncomeTime,
		}).
		Suffix("returning \"id\"")

	q, a, err := sQuery.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return "", err
	}

	err = r.conn.GetContext(ctx, &id, q, a...)
	if err != nil {
		return "", err
	}

	return id, nil
}
