
create database if not exists crypto_coin;

use crypto_coin;

# create can be repeated

create table if not exists coin_order (
    order_id int auto_increment,
    exchange_name varchar(64),
    trading_pair varchar(64),
    price       decimal(18, 8),
    holding     decimal(18, 8),
    earn        decimal(18, 8),
    order_time  datetime,
    primary key (order_id)
) engine=innodb;
