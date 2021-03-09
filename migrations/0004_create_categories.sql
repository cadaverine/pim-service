-- +goose Up

-- категории товаров
create table if not exists products.categories(
	id bigserial primary key,
	parent_id bigint not null default 0,
	name varchar not null default '',
	created_at timestamp not null default now(),
	updated_at timestamp not null default now(),
	deleted_at timestamp
);

create trigger set_timestamp before update on products.categories
for each row execute procedure trigger_set_timestamp();

-- имена категорий на разных языках
create table if not exists products.categories_names(
	category_id bigint references products.categories(id),
	lang_code varchar references catalogs.languages(code),
	name varchar not null default ''
);

-- +goose Down
drop table if exists products.categories cascade;
drop table if exists products.categories_names cascade;
drop trigger if exists set_timestamp on products.categories;

