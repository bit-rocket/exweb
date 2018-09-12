
create database if not exists crypto_coin;

use crypto_coin;

# create can be repeated

create table if not exists coin_order (
    order_id int auto_increment,
    exchange_name varchar(64),
    trading_pair varchar(64),
    order_amount    decimal(18,8) default 0.0,
    holding     decimal(18, 8) default 0.0,
    buy_price   decimal(18, 8) default 0.0,
    sell_price  decimal(18,8) default 0.0,
    earn_rate        decimal(18, 8) default 0.0,
    earn_amount decimal(18,8) default 0.0,
    status      tinyint default 0
        comment "0 created, 1 under buy, 2 holding, 3 under sell, 4 finished",
    create_time datetime,
    finish_time  datetime,
    primary key (order_id)
) engine=innodb;

# insert into coin_order (exchange_name, trading_pair, buy_price, order_amount,
#    holding, sell_price, earn_rate, earn_amount, status, create_time, finish_time)
#    values ("okex", "eos/usdt", 4.32, 200, 100.23, 5.82, 0.15, 30, 3,
#    "2018-09-09 11:23:58", "2018-09-10 11:23:58")
