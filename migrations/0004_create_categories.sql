-- +goose Up

-- категории товаров
create table if not exists product_information.categories(
	id bigserial primary key,
	item_id bigint not null default 0,
	parent_id bigint not null default 0,
	shop_id bigint references product_information.shops(id),
	name varchar not null default '',
	created_at timestamp not null default now(),
	updated_at timestamp not null default now(),
	deleted_at timestamp,
	unique (shop_id, item_id)
);

create trigger set_timestamp before update on product_information.categories
for each row execute procedure trigger_set_timestamp();

-- имена категорий на разных языках
create table if not exists product_information.categories_translations(
	category_id bigint references product_information.categories(id),
	lang_code varchar references catalogs.languages(code),
	translation varchar not null default '',
	created_at timestamp not null default now(),
	updated_at timestamp not null default now(),
	deleted_at timestamp
);

create trigger set_timestamp before update on product_information.categories_translations
for each row execute procedure trigger_set_timestamp();

-- +goose Down
drop table if exists product_information.categories cascade;
drop table if exists product_information.categories_translations cascade;
