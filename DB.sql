CREATE TABLE IF NOT EXISTS users(
    id int not null unique,
    balance real not null);

CREATE TABLE IF NOT EXISTS orders(
    user_id int not null,
    product_id int not null,
    order_id int not null unique,
    price real not null);

CREATE TABLE IF NOT EXISTS report(
    product_id int not null,
    balance real not null,
    month varchar not null);

CREATE TABLE IF NOT EXISTS operations_journal(
    user_id int not null,
    amount real not null,
    date timestamp not null,
    message varchar not null);