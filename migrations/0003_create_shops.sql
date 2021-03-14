-- +goose Up

-- товары
create schema if not exists product_information;

create table if not exists product_information.shops(
	id bigserial primary key,
	name varchar not null default '',
    company varchar not null default '',
    url varchar not null default '',
    platform varchar not null default '',
	created_at timestamp not null default now(),
	updated_at timestamp not null default now(),
	deleted_at timestamp,
	unique (name, company)
);

create trigger set_timestamp before update on product_information.shops
for each row execute procedure trigger_set_timestamp();


-- +goose Down
drop schema if exists product_information cascade;
