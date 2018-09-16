package order

import (
    "fmt"
//    "time"
    "database/sql"

    "github.com/kataras/iris"
    "github.com/kataras/iris/sessions"
    _ "github.com/go-sql-driver/mysql"

    "github.com/yeweishuai/exweb/comm"

)

type Order struct {
    OrderId         string      `json:"orderId"`
    ExName          string      `json:"exName"`
    TradingPair     string      `json:"tradingPair"`
    OrderType       string      `json:"orderType"`
    BuyPrice        string      `json:"buyPrice"`
    OrderAmount     string      `json:"orderAmount"`
    Holding         string      `json:"holding"`
    SellPrice       string      `json:"sellPrice"`
    EarnRate        string      `json:"earnRate"`
    EarnAmount      string      `json:"earnAmount"`
    Status          string      `json:"status"`
    FinishTime      string      `json:"finishTime"`
    CreateTime      string      `json:"createTime"`
}

type OController struct {
    Ctx         iris.Context
    Session     *sessions.Session
}

func (oc *OController) GetTodeal() {
    var orders []Order
    if oc.Session.GetString(comm.UsernameKey) == "" {
        oc.Ctx.JSON(orders)
        return
    }

    // format:
    //          user:pass@/dbname
    db_conf := fmt.Sprintf("%s:%s@/%s",
            comm.GConf.DbConf.DBUser, comm.GConf.DbConf.DBPass,
            comm.GConf.DbConf.DBName)
    db, err := sql.Open("mysql", db_conf)
    if err != nil {
        oc.Ctx.Application().Logger().Infof("db error:%s", err.Error())
        return
    }
    defer db.Close()
    // db refer
    //    http://www.golangprograms.com/example-of-golang-crud-using-mysql-from-scratch.html
    query := `select id, exchange_name, trading_pair, buy_price, order_type,
            order_amount, holding, sell_price, earn_rate, earn_amount,
            status, create_time, finish_time
            from coin_order
            `
    res, err := db.Query(query)
    if err != nil {
        oc.Ctx.Application().Logger().Warnf("db error:%s", err.Error())
        return
    }
    for res.Next() {
        var id, status, order_type int
        var exchange_name, trading_pair string
        var buy_price, holding, sell_price, earn_rate, earn_amount, order_amount float32
        var finish_time, create_time []byte
        err = res.Scan(&id, &exchange_name, &trading_pair,
                &buy_price, &order_type, &order_amount, &holding,
                &sell_price, &earn_rate,
                &earn_amount, &status, &create_time, &finish_time)
        if err != nil {
            oc.Ctx.Application().Logger().Warnf("scan error:%s", err.Error())
            continue
        }
        order := Order {
            OrderId: fmt.Sprintf("%d", id),
            ExName: exchange_name,
            TradingPair: trading_pair,
            BuyPrice: fmt.Sprintf("%.4f", buy_price),
            OrderType: fmt.Sprintf("%s", comm.OrderTypeMap[order_type]),
            OrderAmount: fmt.Sprintf("%.4f", order_amount),
            Holding: fmt.Sprintf("%.4f", holding),
            SellPrice: fmt.Sprintf("%.4f", sell_price),
            EarnRate: fmt.Sprintf("%.4f", earn_rate),
            EarnAmount: fmt.Sprintf("%.4f", earn_amount),
            Status: fmt.Sprintf("%d", status),
            CreateTime: fmt.Sprintf("%s", string(create_time)),
            FinishTime: fmt.Sprintf("%s", string(finish_time)),
        }
        orders = append(orders, order)
    }
    oc.Ctx.JSON(orders)
}

func (oc *OController) PostNew() {
    var order Order
    err := oc.Ctx.ReadJSON(&order)
    if err != nil {
        oc.Ctx.Application().Logger().Warnf("read json error:%s", err.Error())
        oc.Ctx.JSON(map[string]string{"msg":err.Error()})
        return
    }
    oc.Ctx.Application().Logger().Infof("read json :%v", order)
    oc.Ctx.JSON(map[string]string{"msg":"ok"})
}
