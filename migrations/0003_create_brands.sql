-- +goose Up

-- товары
create schema if not exists products;

create table if not exists products.brands(
	name varchar primary key,
	logo varchar not null default '',
	site varchar not null default '',
	company varchar not null default '',
	created_at timestamp not null default now(),
	updated_at timestamp not null default now(),
	deleted_at timestamp
);

create trigger set_timestamp after update on products.brands
for each row execute procedure trigger_set_timestamp();


-- +goose Down
drop schema if exists products cascade;
drop trigger if exists products_brands_tmstp;