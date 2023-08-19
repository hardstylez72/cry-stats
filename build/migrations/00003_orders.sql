-- +goose Up
-- +goose StatementBegin
create table if not exists orders (
    id uuid PRIMARY KEY,
    account_id uuid,
    net text not null,
    addr text not null,
    confirmed_by_tx_id uuid,
    income_expected numeric(8,3),
    income_received numeric(8,3),
    status text,
    meta text,
    exchange_rate real,
    created_at timestamp not null default now(),
    confirmed_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists orders;
-- +goose StatementEnd
