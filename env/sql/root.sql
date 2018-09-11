
create database if not exists crypto_coin;

# change db password !

create user 'trader'@'localhost' identified by 'traderxyz';
grant all privileges on crypto_coin.* to 'trader'@'localhost';
