package order

import (
    "fmt"
    "database/sql"

    "github.com/kataras/iris"
    _ "github.com/go-sql-driver/mysql"

    "github.com/yeweishuai/exweb/comm"

)

type Order struct {
    UserId      string      `json:"usersId"`
    ExName      string      `json:"exName"`
    TradingPair string      `json:"tradingPair"`
    Price       string      `json:"price"`
    Holding     string      `json:"holding"`
    Earn        string      `json:"earn"`
    OrderTime   string      `json:"orderTime"`
}

func ToDealOrders(ctx iris.Context) {
    var orders []Order
    mock := Order {
        UserId: "1",
        ExName: "okex",
        TradingPair: "eos/usdt",
        Price: "5.01",
        Holding: "+1.02%",
        Earn: "20",
        OrderTime: "2017-05-10 10:30",
    }

    // format:
    //          user:pass@/dbname
    db_conf := fmt.Sprintf("%s:%s@/%s",
            comm.GConf.DbConf.DBUser, comm.GConf.DbConf.DBPass,
            comm.GConf.DbConf.DBName)
    db, err := sql.Open("mysql", db_conf)
    if err != nil {
        ctx.Application().Logger().Infof("db error:%s", err.Error())
        return
    }
    defer db.Close()
    // db refer
    //    http://www.golangprograms.com/example-of-golang-crud-using-mysql-from-scratch.html


    orders = append(orders, mock)
    ctx.JSON(orders)
}

