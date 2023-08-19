package order

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	conn *sqlx.DB
}

func NewRepository(conn *sqlx.DB) *Repository {
	return &Repository{
		conn: conn,
	}
}

type Order struct {
	Id          string    `db:"id"`
	Net         string    `db:"net"`
	Addr        string    `db:"addr"`
	Status      string    `db:"status"`
	CreatedAt   time.Time `db:"created_at"`
	ConfirmedAt time.Time `db:"confirmed_at"`
	Amount      float64   `db:"income_expected"`
}

func (r *Repository) GetOrderHistory(ctx context.Context, userId string) ([]Order, error) {

	q := `select id, net, addr, status, created_at, confirmed_at, income_expected from orders 
		where account_id = $1 and status in ('Processed', 'Error') order by created_at desc`

	orders := make([]Order, 0)
	err := r.conn.SelectContext(ctx, &orders, q, userId)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
