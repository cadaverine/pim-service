create database pim_db template template0 encoding UTF8 LC_COLLATE "ru_RU.UTF-8" LC_CTYPE "ru_RU.UTF-8";

create or replace function trigger_set_timestamp()
returns trigger as $$
begin
    new.updated_at = now();
    return new;
end;
$$ language plpgsql;

-- справочники
create schema if not exists catalogs;

-- берем с api.hh.ru
create table if not exists catalogs.languages(
    code varchar primary key,
    name varchar not null default '',
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);

create trigger set_timestamp after update on catalogs.languages
for each row execute procedure trigger_set_timestamp();

insert into catalogs.languages(code, name)
values ('rus', 'Русский'), ('eng', 'Английский');

create table if not exists catalogs.currencies(
    code varchar not null default '',
    name varchar not null default '',
    rate integer not null default 1,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    primary key (code)
);

create trigger set_timestamp after update on catalogs.currencies
for each row execute procedure trigger_set_timestamp();

insert into catalogs.currencies(code, name, rate)
values ('RUB', 'Российский рубль', 1);

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
    unique (name, company)
);

create trigger set_timestamp before update on product_information.shops
for each row execute procedure trigger_set_timestamp();

-- категории товаров
create table if not exists product_information.categories(
    id bigserial primary key,
    item_id bigint not null default 0,
    parent_id bigint not null default 0,
    shop_id bigint references product_information.shops(id),
    name varchar not null default '',
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    unique (shop_id, item_id)
);

create trigger set_timestamp before update on product_information.categories
for each row execute procedure trigger_set_timestamp();

-- имена категорий на разных языках
create table if not exists product_information.categories_translations(
    category_id bigint references product_information.categories(id) on delete cascade,
    lang_code varchar references catalogs.languages(code),
    translation varchar not null default '',
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);

create trigger set_timestamp before update on product_information.categories_translations
for each row execute procedure trigger_set_timestamp();

-- допустимые типы аттрибутов
create type product_information.attributes_type as enum(
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
    vendor varchar not null default '',
    description varchar not null default '',
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    unique (shop_id, item_id)
);

create trigger set_timestamp after update on product_information.products
for each row execute procedure trigger_set_timestamp();

-- категории конкретных товаров
create table if not exists product_information.products_categories(
    product_id bigint references product_information.products(id) on delete cascade,
    category_id bigint references product_information.categories(id) on delete cascade,
    unique (product_id, category_id)
);

-- аттрибуты конкретных товаров
create table if not exists product_information.products_attributes(
    id bigserial primary key,
    product_id bigint references product_information.products(id) on delete cascade,
    name varchar not null default '',
    type product_information.attributes_type not null,
    value jsonb,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    unique (product_id, name)
);

create trigger set_timestamp after update on product_information.products_attributes
for each row execute procedure trigger_set_timestamp();

-- имена аттрибутов на разных языках
create table if not exists product_information.products_attributes_translations(
    attribute_id int references product_information.products_attributes(id) on delete cascade,
    lang_code varchar references catalogs.languages(code) on delete cascade,
    translation varchar not null default '',
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);

create trigger set_timestamp after update on product_information.products_attributes_translations
for each row execute procedure trigger_set_timestamp();

-- используем для нечеткого поиска по продуктам
create extension if not exists pg_trgm;

create index products_trgm_idx
on product_information.products
using gin (name gin_trgm_ops);
