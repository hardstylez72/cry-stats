package account

import (
	"context"
	"database/sql"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Repository struct {
	conn *sqlx.DB
}

func New(conn *sqlx.DB) *Repository {

	return &Repository{
		conn: conn,
	}
}

type GetByIdRes struct {
	Id     string `db:"id"`
	Login  string `db:"login"`
	Funds  int    `db:"funds"`
	Status string `db:"status"`
}

func (r *Repository) GetByLogin(ctx context.Context, login string) (*GetByIdRes, error) {
	query := "select id, login, funds, status from accounts where login = $1 "

	account := &GetByIdRes{}
	err := r.conn.GetContext(ctx, account, query, login)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrAccountNotFound
		}
		return nil, err
	}

	return account, nil
}

const AccountStatusCreated = "CREATED"
const DefaultInitialValue = 100

func (r *Repository) CreateAccount(ctx context.Context, id, login string, funds, taskPrice float64) (err error) {

	sQuery := squirrel.Insert("accounts").
		Columns("id", "login", "funds", "status", "task_price").
		Values(id, login, funds, AccountStatusCreated, taskPrice)

	q, a, err := sQuery.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return err
	}

	_, err = r.conn.ExecContext(ctx, q, a...)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

func (r *Repository) AccountExist(ctx context.Context, id string) (*bool, error) {

	q := `select count(1) from accounts where id = $1`

	rows, err := r.conn.QueryContext(ctx, q, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	defer rows.Close()
	rows.Next()
	var count = 0
	if err := rows.Scan(&count); err != nil {
		return nil, err
	}

	exist := count == 1

	return &exist, nil
}

type Account struct {
	Id        string  `db:"id"`
	Login     string  `db:"login"`
	Funds     float64 `db:"funds"`
	Status    string  `db:"status"`
	TaskPrice float64 `db:"task_price"`
}

var ErrAccountNotFound = errors.New("account not found")

func (r *Repository) GetAccount(ctx context.Context, id string) (*Account, error) {

	q := `select id, login, funds, status, task_price from accounts where id = $1`

	var acc Account
	err := r.conn.GetContext(ctx, &acc, q, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrAccountNotFound
		}
		return nil, err
	}

	return &acc, nil
}

func (r *Repository) GetHistoryRecords(ctx context.Context, userId string, offset int) ([]TaskHistoryRecord, error) {

	q := `select task_id, task_type, user_id, process_id, created_at, price from accounts_task_history
   where user_id = $1 order by created_at desc limit 10 offset $2`

	out := make([]TaskHistoryRecord, 0)
	err := r.conn.SelectContext(ctx, &out, q, userId, offset)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrAccountNotFound
		}
		return nil, err
	}

	return out, nil
}

func (r *Repository) GetFundsById(ctx context.Context, id string) (funds float64, isFound bool, err error) {
	sQuery := squirrel.Select("funds").
		From("accounts").
		Where("id = ?", id)

	q, a, err := sQuery.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return 0, false, err
	}

	err = r.conn.GetContext(ctx, &funds, q, a...)
	if err == sql.ErrNoRows {
		return 0, false, nil
	}
	if err != nil {
		return 0, false, err
	}

	return funds, true, nil
}

func (r *Repository) GetFundsByLogin(ctx context.Context, login string) (funds float64, isFound bool, err error) {
	sQuery := squirrel.Select("funds").
		From("accounts").
		Where("login = ?", login)

	q, a, err := sQuery.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return 0, false, err
	}

	err = r.conn.GetContext(ctx, &funds, q, a...)
	if err == sql.ErrNoRows {
		return 0, false, nil
	}
	if err != nil {
		return 0, false, err
	}

	return funds, true, nil
}

type DecrementFundsRes struct {
	Funds int `db:"funds"`
}

func (r *Repository) DecrementFundsByUserId(ctx context.Context, id string, value float64) (err error) {

	q := `update accounts set funds = funds - $1 where id = $2`

	_, err = r.conn.ExecContext(ctx, q, value, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) AddFundsById(ctx context.Context, id string, value float64) (newFunds int, err error) {

	q := `update accounts set funds = funds + $1 where id = $2`

	err = r.conn.GetContext(ctx, &newFunds, q, value, id)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	return newFunds, nil
}
func (r *Repository) TaskExist(ctx context.Context, taskId string) (bool, error) {
	q := `select count(*) from accounts_task_history where task_id = $1`

	var number int
	if err := r.conn.GetContext(ctx, &number, q, taskId); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return number > 0, nil
}
func (r *Repository) AddTaskRecord(ctx context.Context, rec *TaskHistoryRecord) error {

	q := `insert into accounts_task_history (task_id, task_type, user_id, process_id, created_at, price)
    values (:task_id, :task_type, :user_id, :process_id, now(), :price) on conflict (task_id) do nothing `

	_, err := r.conn.NamedExecContext(ctx, q, rec)
	if err != nil {
		return err
	}

	return nil
}

type TaskHistoryRecord struct {
	TaskId    string    `db:"task_id"`
	TaskType  string    `db:"task_type"`
	UserId    string    `db:"user_id"`
	ProcessId string    `db:"process_id"`
	CreatedAt time.Time `db:"created_at"`
	Price     float64   `db:"price"`
}
