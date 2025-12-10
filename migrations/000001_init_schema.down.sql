-- Necesito borrar las tablas con foreign key --

drop table if exists payments;
drop table if exists order_items;
drop table if exists orders;
drop table if exists lost_inventories;
drop table if exists needed_ingredients;
drop table if exists inventories;

-- Ahora borrar las tables bases q no tienen foreign keys --

drop table if exists ingredients;
drop table if exists dishes;
drop table if exists restaurants;
drop table if exists users;

-- Borrar dominios y tipos enums que cree --

drop type if exists payment_method;
drop type  if exists payment_status;
drop type if exists order_status;
drop type if exists measure_unit;
drop type if exists role;

drop domain if exists non_negative_number;