-- +goose Up

-- допустимые типы аттрибутов
create type products.attributes_type as enum(
	'string',
	'integer',
	'float',
	'date',
	'time',
	'fixed',
	'media'
);

-- каталог аттрибутов
create table if not exists products.attributes(
	id bigserial primary key,
	name varchar not null default '' ,
	type products.attributes_type not null,
    created_at timestamp not null default now(),
	updated_at timestamp not null default now(),
	deleted_at timestamp,
	unique(name)
);

create trigger set_timestamp after update on products.attributes
for each row execute procedure trigger_set_timestamp();


-- имена аттрибутов на разных языках
create table if not exists products.attributes_names(
	attribute_id bigint references products.attributes(id),
	lang_code varchar references catalogs.languages(code),
	name varchar not null default '',
    created_at timestamp not null default now(),
	updated_at timestamp not null default now(),
	deleted_at timestamp
);

create trigger set_timestamp after update on products.attributes_names
for each row execute procedure trigger_set_timestamp();

-- товары
create table if not exists products.items(
	id varchar primary key,
	name varchar not null default '',
	brand_name varchar references products.brands(name),
    available boolean not null default true,
    type varchar not null default '',
    url varchar not null default '',
    price varchar not null default '0',
    currency_code varchar references catalogs.currencies (code),
    category_id bigint references products.categories (id),
    vendor varchar not null default '',
    description varchar not null default '',
	created_at timestamp default current_timestamp,
	updated_at timestamp default current_timestamp,
	deleted_at timestamp
);

create trigger set_timestamp after update on products.items
for each row execute procedure trigger_set_timestamp();

-- категории конкретных товаров
create table if not exists products.items_categories(
	product_id varchar references products.items(id),
	category_id bigint references products.categories(id)
);

-- аттрибуты конкретных товаров
create table if not exists products.items_attributes(
	item_id varchar references products.items(id),
	attribute_id bigint references products.attributes(id),
	value jsonb,
	created_at timestamp default current_timestamp,
	updated_at timestamp default current_timestamp,
    deleted_at timestamp
);


-- +goose Down
drop table if exists products.items_categories cascade;
drop table if exists products.items_attributes cascade;
drop table if exists products.items cascade;
drop trigger if exists set_timestamp on products.items cascade;
drop table if exists products.attributes cascade;
drop trigger if exists set_timestamp on products.attributes cascade;
drop table if exists products.attributes_names cascade;
drop trigger if exists set_timestamp on products.attributes_names cascade;
drop table if exists products.items_categories cascade;
drop type products.attributes_type cascade;