-- +goose Up
-- +goose StatementBegin
create table if not exists tx_offsets (
    id uuid PRIMARY KEY,
    net text not null,
    addr text not null,
    tx_total integer not null default 0,
    tx_offset integer not null default 0,
    updated_at timestamp not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists tx_offsets;
-- +goose StatementEnd
