package order

import (
    "fmt"
//    "time"
    "database/sql"

    "github.com/kataras/iris"
    _ "github.com/go-sql-driver/mysql"

    "github.com/yeweishuai/exweb/comm"

)

type Order struct {
    OrderId     string      `json:"orderId"`
    ExName      string      `json:"exName"`
    TradingPair string      `json:"tradingPair"`
    BuyPrice    string      `json:"buyPrice"`
    OrderAmount string      `json:"orderAmount"`
    Holding     string      `json:"holding"`
    SellPrice   string      `json:"sellPrice"`
    EarnRate    string      `json:"earnRate"`
    EarnAmount  string      `json:"earnAmount"`
    Status      string      `json:"status"`
    FinishTime  string      `json:"finishTime"`
    CreateTime  string      `json:"createTime"`
}

func ToDealOrders(ctx iris.Context) {
    var orders []Order

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
    query := `select order_id, exchange_name, trading_pair, buy_price,
            order_amount, holding, sell_price, earn_rate, earn_amount,
            status, create_time, finish_time
            from coin_order
            `
    res, err := db.Query(query)
    if err != nil {
        ctx.Application().Logger().Warnf("db error:%s", err.Error())
        return
    }
    for res.Next() {
        var order_id, status int
        var exchange_name, trading_pair string
        var buy_price, holding, sell_price, earn_rate, earn_amount, order_amount float32
        var finish_time, create_time []byte
        err = res.Scan(&order_id, &exchange_name, &trading_pair,
                &buy_price, &order_amount, &holding, &sell_price, &earn_rate,
                &earn_amount, &status, &create_time, &finish_time)
        if err != nil {
            ctx.Application().Logger().Warnf("scan error:%s", err.Error())
            continue
        }
        order := Order {
            OrderId: fmt.Sprintf("%d", order_id),
            ExName: exchange_name,
            TradingPair: trading_pair,
            BuyPrice: fmt.Sprintf("%f", buy_price),
            OrderAmount: fmt.Sprintf("%f", order_amount),
            Holding: fmt.Sprintf("%f", holding),
            SellPrice: fmt.Sprintf("%f", sell_price),
            EarnRate: fmt.Sprintf("%f", earn_rate),
            EarnAmount: fmt.Sprintf("%f", earn_amount),
            Status: fmt.Sprintf("%d", status),
            CreateTime: fmt.Sprintf("%s", string(create_time)),
            FinishTime: fmt.Sprintf("%s", string(finish_time)),
        }
        orders = append(orders, order)
    }

    ctx.JSON(orders)
}

