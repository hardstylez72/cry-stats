-- +goose Up
alter table if exists accounts
    add if not exists promo text null;

-- +goose Down
alter table if exists accounts
    drop column if exists promo;
