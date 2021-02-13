-- справочники
create schema if not exists catalogs;

-- берем с api.hh.ru
create table if not exists catalogs.languages(
	id bigserial primary key,
	code varchar not null default '',
	title varchar not null default ''
);

create table if not exists catalogs.brands(
	id bigserial primary key,
	title varchar not null default '',
	logo varchar not null default '',
	site varchar not null default ''
);



-- все, что относится к продуктам
create schema if not exists products;

-- категории товаров
create table if not exists products.categories_items(
	id bigserial primary key,
	title varchar not null default ''
);

-- категории имеют древовидную структуру
-- (может ли категория иметь более одного родителя?)
create table if not exists products.categories_relations(
	id bigint references products.categories_items(id),
	parent_id bigint references products.categories_items(id)
);

-- имена категорий на разных языках
create table if not exists products.categories_names(
	category_id bigint references products.categories_items(id),
	language_id bigint references catalogs.languages(id),
	name varchar not null default ''
);


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
create table if not exists products.attributes_items(
	id bigserial primary key,
	name varchar not null default '',
	type products.attributes_type not null,
	unique(name)
);

-- имена аттрибутов на разных языках
create table if not exists products.attributes_names(
	attribute_id bigint references products.attributes_items(id),
	language_id bigint references catalogs.languages(id),
	name varchar not null default ''
);


-- товары
-- (может ли товар иметь более одной категории?)
create table if not exists products.items(
	id bigserial primary key,
	brand_id bigint references catalogs.brands(id),
	created_at timestamp default current_timestamp,
	updated_at timestamp default current_timestamp,
	created_by varchar not null default ''
);

-- категории конкретных товаров
create table if not exists products.items_categories(
	product_id bigint references products.items(id),
	category_id bigint references products.categories_items(id)
);

-- аттрибуты конкретных товаров
create table if not exists products.items_attributes(
	item_id bigint references products.items(id),
	attribute_id bigint references products.attributes_items(id),
	value jsonb
);
