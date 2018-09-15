
create database if not exists crypto_coin;

use crypto_coin;

# create can be repeated

create table if not exists coin_order (
    id int auto_increment,
    exchange_name varchar(64),
    trading_pair varchar(64),
    order_amount    decimal(18,8) default 0.0,
    holding     decimal(18, 8) default 0.0,
    buy_price   decimal(18, 8) default 0.0,
    buy_cost    decimal(18, 8) default 0.0,
    sell_price  decimal(18,8) default 0.0,
    sell_income decimal(18, 8) default 0.0,
    earn_rate        decimal(18, 8) default 0.0,
    earn_amount decimal(18,8) default 0.0,
    status      tinyint default 0
        comment "0 created, 1 under buy, 2 holding, 3 under sell, 4 finished",
    create_time datetime,
    finish_time  datetime,
    primary key (id)
) engine=innodb;

 insert into coin_order (exchange_name, trading_pair, buy_price, order_amount,
    holding, sell_price, earn_rate, earn_amount, status, create_time, finish_time)
    values ("okex", "eos/usdt", 4.32, 200, 100.23, 5.82, 0.15, 30, 3,
    "2018-09-09 11:23:58", "2018-09-10 11:23:58");

create table if not exists coin_son_order (
    id  int auto_increment,
    son_order_id  int  comment "exchange order id",
    parent_id   int   comment "sql record id",
    exchange_name   varchar(64),
    trading_pair varchar(64),
    order_amount  decimal(18, 8) default 0.0,
    deal_amount decimal(18, 8) default 0.0,
    price   decimal(18, 8) default 0.0,
    order_type  tinyint default 0  comment " 1 buy, 2 sell",
    status  tinyint default 0
        comment "0 created, 1 order made, 2 partial dealt, 3 manually undo, 4 finish",
    create_time datetime,
    finish_time  datetime,

    primary key (id)
) engine=innodb;
