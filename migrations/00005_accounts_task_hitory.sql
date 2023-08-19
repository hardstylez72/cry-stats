-- +goose Up
-- +goose StatementBegin
create table if not exists accounts_task_history (
    task_id uuid PRIMARY KEY,
    task_type text not null,
    user_id uuid not null,
    process_id uuid not null,
    created_at timestamp not null,
    price numeric(8,3)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists accounts_task_history;
-- +goose StatementEnd
