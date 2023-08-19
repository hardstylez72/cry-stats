package task

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	conn *sqlx.DB
}

func New(conn *sqlx.DB) *Repository {
	return &Repository{
		conn: conn,
	}
}

type GetOffsetsReq struct {
	Net  string
	Addr string
}

type GetOffsetsRes struct {
	TxTotal  int `db:"tx_total"`
	TxOffset int `db:"tx_offset"`
}

func (r *Repository) GetOffsets(ctx context.Context, req *GetOffsetsReq) (*GetOffsetsRes, error) {
	sQuery := squirrel.Select("tx_total", "tx_offset").
		From("tx_offsets").
		Where(
			squirrel.Eq{
				"net":  req.Net,
				"addr": req.Addr,
			},
		)

	q, a, err := sQuery.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	res := GetOffsetsRes{}
	err = r.conn.GetContext(ctx, &res, q, a...)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &res, nil
}

func (r *Repository) InsertOffsets(ctx context.Context, req *UpdateOffsetsReq) error {
	sQuery := squirrel.Insert("tx_offsets").
		Columns("id", "net", "addr", "tx_total", "tx_offset").
		Values(uuid.NewString(), req.Net, req.Addr, req.TxTotal, req.TxOffsetIncrBy)

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

type UpdateOffsetsReq struct {
	Net            string
	Addr           string
	TxTotal        int
	TxOffsetIncrBy int
}

func (r *Repository) UpdateOffsets(ctx context.Context, req *UpdateOffsetsReq) error {
	log.Printf("Req IncrBy %+v \n", req)
	sQuery := squirrel.Update("tx_offsets").
		Set("tx_total", req.TxTotal).
		Set("tx_offset", squirrel.Expr(fmt.Sprintf("tx_offset+%d", req.TxOffsetIncrBy))).
		Set("updated_at", squirrel.Expr("now()")).
		Where(
			squirrel.Eq{
				"net":  req.Net,
				"addr": req.Addr,
			},
		)
	q, a, err := sQuery.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return err
	}

	sqlRes, err := r.conn.ExecContext(ctx, q, a...)
	if err != nil {
		return err
	}

	rowsAffected, err := sqlRes.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		r.InsertOffsets(ctx, req)
	}

	return nil
}
