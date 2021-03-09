-- +goose Up

-- справочники
create schema if not exists catalogs;

-- берем с api.hh.ru
create table if not exists catalogs.languages(
	code varchar primary key,
	name varchar not null default '',
	created_at timestamp not null default now(),
	updated_at timestamp not null default now(),
	deleted_at timestamp
);

create trigger set_timestamp after update on catalogs.languages
for each row execute procedure trigger_set_timestamp();

create table if not exists catalogs.currencies(
	code varchar primary key,
	rate integer not null default 1,
	created_at timestamp not null default now(),
	updated_at timestamp not null default now(),
	deleted_at timestamp
);

create trigger set_timestamp after update on catalogs.languages
for each row execute procedure trigger_set_timestamp();

-- +goose Down
drop schema if exists catalogs cascade;


