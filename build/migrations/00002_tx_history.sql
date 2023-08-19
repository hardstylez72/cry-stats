-- +goose Up
-- +goose StatementBegin
create table if not exists tx_history (
    id uuid PRIMARY KEY,
    net text not null,
    addr text not null,
    tx_hash text not null,
    income_received integer not null,
    created_at timestamp not null default now(),
    processed_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists tx_history;
-- +goose StatementEnd
