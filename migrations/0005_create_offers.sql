-- +goose Up

-- допустимые типы аттрибутов
create type if not exists product_information.attributes_type as enum(
	'string',
	'integer',
	'float',
	'date',
	'time',
	'fixed',
	'media'
);

-- товары
create table if not exists product_information.products(
	id bigserial primary key,
	item_id varchar not null default '',
	shop_id bigint references product_information.shops(id),
	name varchar not null default '',
    available boolean not null default true,
    type varchar not null default '',
    url varchar not null default '',
    price int not null default 0,
    currency_code varchar references catalogs.currencies (code),
    category_id bigint not null,
    vendor varchar not null default '',
    description varchar not null default '',
	created_at timestamp default current_timestamp,
	updated_at timestamp default current_timestamp,
	deleted_at timestamp,
	unique (shop_id, item_id),
	foreign key (shop_id, category_id) references product_information.categories (shop_id, item_id)
);

create trigger set_timestamp after update on product_information.products
for each row execute procedure trigger_set_timestamp();

-- категории конкретных товаров
create table if not exists product_information.products_categories(
	product_id bigint references product_information.products(id),
	category_id bigint references product_information.categories(id)
);

-- аттрибуты конкретных товаров
create table if not exists product_information.products_attributes(
	id bigserial primary key,
	product_id bigint references product_information.products(id),
    name varchar not null default '',
    type product_information.attributes_type not null,
	value jsonb,
	created_at timestamp default current_timestamp,
	updated_at timestamp default current_timestamp,
    deleted_at timestamp,
    unique (product_id, name)
);

create trigger set_timestamp after update on product_information.products_attributes
for each row execute procedure trigger_set_timestamp();

-- имена аттрибутов на разных языках
create table if not exists product_information.products_attributes_translations(
	attribute_id int references product_information.products_attributes(id),
	lang_code varchar references catalogs.languages(code),
	translation varchar not null default '',
    created_at timestamp not null default now(),
	updated_at timestamp not null default now(),
	deleted_at timestamp
);

create trigger set_timestamp after update on product_information.products_attributes_translations
for each row execute procedure trigger_set_timestamp();

-- +goose Down
drop table if exists product_information.products_attributes_translations cascade;
drop table if exists product_information.products_attributes cascade;
drop table if exists product_information.products_categories cascade;
drop table if exists product_information.products cascade;
drop type if exists product_information.attributes_type cascade;



