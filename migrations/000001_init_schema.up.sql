-- USUARIOS Y ROLES --

create type "role" as enum (
    'client',
    'admin',
    'chef',
    'cashier'
    );

create table users (
    "user_id" serial primary key,
    "user_name" varchar(25) not null unique,
    "email" varchar not null,
    "first_name" varchar(25) not null,
    "last_name" varchar(25) not null,
    "country_code" char(3) not null,
    "user_role" role not null,
    "created_at" timestamp with time zone default now(),
    "is_active" bool default true not null ,
    "password_hash" varchar(255) not null
);

-- RESTAURANTES --

create table restaurants (
    "restaurant_id" serial primary key,
    "restaurant_name" varchar not null ,
    "address" varchar not null ,
    "is_active" bool default true not null,
    "is_open" bool default false not null
);

-- MENU E INVENTARIO --

create type "measure_unit" as enum (
    'u',
    'kg',
    'g',
    'l',
    'ml'
    );

create table dishes (
    "dish_id" serial primary key ,
    "restaurant_id" int not null ,
    dish_name varchar not null ,
    "price" decimal(10,2) not null ,

    foreign key (restaurant_id) references restaurants(restaurant_id)
);

create table ingredients (
    "ingredient_id" varchar primary key ,
    "ingredient_name" varchar not null
);

-- Esto es como crea un politica en mi base de datos
create domain "non_negative_number" as numeric
check ( value >= 0 );

create table inventories (
    "inventory_id" serial primary key ,
    "restaurant_id" int not null ,
    "ingredient_id" varchar not null ,
    "stock" non_negative_number not null ,
    "unit" measure_unit not null,
    "price" non_negative_number not null ,
    "expiration_date" date,

    foreign key (restaurant_id) references restaurants(restaurant_id),
    foreign key (ingredient_id) references ingredients(ingredient_id, ingredient_name)
);

create table needed_ingredients (
    "dish_id" int not null ,
    "ingredient_id" varchar not null ,
    "needed_quantity" decimal(10,2) not null ,
    "unit" measure_unit not null ,
    "is_optional" bool default false,

    primary key (dish_id,ingredient_id),

    foreign key (ingredient_id) references ingredients(ingredient_id),
    foreign key (dish_id) references dishes(dish_id)
);

create table lost_inventories (
    "lost_id" serial primary key ,
    "inventory_id" int not null ,
    "quantity" decimal(10,2) not null ,
    "unit" measure_unit not null ,
    "unit_price" decimal(10,2) not null ,
    "date_event" timestamp with time zone default now(),

    foreign key (inventory_id) references inventories(inventory_id)
);

-- ORDENES --

create type "order_status" as enum (
    'received',
    'in_processing',
    'completed',
    'canceled'
    );

create table orders (
    "order_id" serial primary key ,
    "user_id" int not null ,
    "restaurant_id" int not null ,
    "created_at" timestamp with time zone default now() not null ,
    "status" order_status not null default 'received',

    foreign key (user_id) references users(user_id),
    foreign key (restaurant_id) references restaurants(restaurant_id)
);

create table order_items (
    "order_item_id" serial primary key ,
    "order_id" int not null ,
    "dish_id" int not null ,
    "quantity" int not null  default 1,
    "price_at_moment" decimal(10,2) not null,

    foreign key (order_id) references orders(order_id),
    foreign key (dish_id) references dishes(dish_id)
);

-- PAGOS --

create type "payment_status" as enum (
    'pending',
    'succeeded',
    'failed',
    'refunded'
    );

create type "payment_method" as enum (
    'cash',
    'paypal',
    'mercadopago'
    );

create table payments (
    "payment_id" serial primary key ,
    "user_id" int not null ,
    "order_id" int not null ,
    "amount" decimal(10,2) not null ,
    "currency" char(3) default 'USD',
    "method" payment_method not null ,
    "created_at" timestamp with time zone default now(),

    foreign key (user_id) references users(user_id),
    foreign key (order_id) references orders(order_id)
);