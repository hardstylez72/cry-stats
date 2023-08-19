-- +goose Up
-- +goose StatementBegin
create table if not exists accounts (
    id uuid PRIMARY KEY,
    login text unique,
    funds numeric(8,3),
    status text,
    task_price numeric(8,3)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists accounts;
-- +goose StatementEnd
