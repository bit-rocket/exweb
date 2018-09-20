package order

import (
    "fmt"
    "encoding/json"
//    "time"
    "database/sql"

    "github.com/kataras/iris"
    "github.com/kataras/iris/sessions"
    _ "github.com/go-sql-driver/mysql"

    "github.com/yeweishuai/exweb/comm"

)

type Order struct {
    OrderId         int         `json:"orderId"`
    ExName          string      `json:"exName"`
    TradingPair     string      `json:"tradingPair"`
    OrderType       int         `json:"orderType"`
    BuyPrice        float64     `json:"buyPrice"`
    OrderAmount     float64     `json:"orderAmount"`
    Holding         string      `json:"holding"`
    SellPrice       float64     `json:"sellPrice"`
    EarnRate        string      `json:"earnRate"`
    EarnAmount      string      `json:"earnAmount"`
    Status          string      `json:"status"`
    FinishTime      string      `json:"finishTime"`
    CreateTime      string      `json:"createTime"`
}

type TodealRes struct {
    Code            int     `json:"code"`
    Msg             string  `josn:"msg"`
    Count           int     `json:"count"`
    Data            []Order `json:"data"`
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
        var buy_price, holding, sell_price, earn_rate, earn_amount, order_amount float64
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
            OrderId: id,
            ExName: exchange_name,
            TradingPair: trading_pair,
            BuyPrice: buy_price,
            OrderType: order_type,
            OrderAmount: order_amount,
            Holding: fmt.Sprintf("%.4f", holding),
            SellPrice: sell_price,
            EarnRate: fmt.Sprintf("%.4f", earn_rate),
            EarnAmount: fmt.Sprintf("%.4f", earn_amount),
            Status: fmt.Sprintf("%s", comm.StatusMap[status]),
            CreateTime: fmt.Sprintf("%s", string(create_time)),
            FinishTime: fmt.Sprintf("%s", string(finish_time)),
        }
        orders = append(orders, order)
    }

    todealRes := TodealRes {
        Code: 0,
        Msg: "",
        Count: len(orders),
        Data: orders,
    }
    oc.Ctx.JSON(todealRes)
}

func (oc *OController) PostNew() {
    dft_msg := map[string]string{"msg":"access denied."}
    if oc.Session.GetString(comm.UsernameKey) == "" {
        oc.Ctx.Application().Logger().Warnf("deny unauthorized access from:%s",
                oc.Ctx.RemoteAddr())
        oc.Ctx.JSON(dft_msg)
        return
    }

    var order Order
    err := oc.Ctx.ReadJSON(&order)
    if err != nil {
        oc.Ctx.Application().Logger().Warnf("read json error:%s", err.Error())
        oc.Ctx.JSON(map[string]string{"msg":err.Error()})
        return
    }
    order_json, err := json.Marshal(order)
    oc.Ctx.Application().Logger().Infof("read json :%s", order_json)
    if order.OrderId > 0 {
        oc.Ctx.Application().Logger().Warnf("ip[%s] order id > 0! order:%s",
                oc.Ctx.RemoteAddr(), order_json)
        dft_msg["msg"] = "post json error!"
        oc.Ctx.JSON(dft_msg)
        return
    }
    // try insert into db
    db_conf := fmt.Sprintf("%s:%s@/%s",
            comm.GConf.DbConf.DBUser, comm.GConf.DbConf.DBPass,
            comm.GConf.DbConf.DBName)
    db, err := sql.Open("mysql", db_conf)
    if err != nil {
        oc.Ctx.Application().Logger().Infof("ip[%s] db error:%s",
                oc.Ctx.RemoteAddr(), err.Error())
        return
    }
    defer db.Close()

    price_field_name := "buy_price"
    price_field_value := order.BuyPrice
    if order.OrderType == comm.OrderSellThenBuy {
        price_field_name = "sell_price"
        price_field_value = order.SellPrice
    }
    prepare_sql := fmt.Sprintf("insert into coin_order (exchange_name, trading_pair, order_type," +
            "order_amount, %s, create_time) values (?, ?, ?, ?, ?, now())", price_field_name)
    insForm, err := db.Prepare(prepare_sql)
    if err != nil {
        oc.Ctx.Application().Logger().Warnf("prepare sql[%s] error:%s",
                prepare_sql, err.Error())
        dft_msg["msg"] = "internal sql error."
        oc.Ctx.JSON(dft_msg)
        return
    }
    insForm.Exec(order.ExName, order.TradingPair, order.OrderType,
            order.OrderAmount, price_field_value)

    oc.Ctx.JSON(map[string]string{"msg":"ok"})
    return
}

func (oc *OController) PostDigest() {
    dft_msg := map[string]string{"msg":"access denied."}
    if oc.Session.GetString(comm.UsernameKey) == "" {
        oc.Ctx.Application().Logger().Warnf("digest deny unauthorized access from:%s",
                oc.Ctx.RemoteAddr())
        oc.Ctx.JSON(dft_msg)
        return
    }
    // read post json
    var order Order
    err := oc.Ctx.ReadJSON(&order)
    if err != nil {
        oc.Ctx.Application().Logger().Warnf("digest read json error:%s", err.Error())
        oc.Ctx.JSON(map[string]string{"msg":err.Error()})
        return
    }
    order_json, err := json.Marshal(order)
    oc.Ctx.Application().Logger().Infof("digest read json :%s", order_json)

    // conditions judge

    // - parent exists
    // - holding amount can cover
    // - TODO price judge
    // - order type match
    // - insert a son order

    oc.Ctx.JSON(map[string]string{"msg":"ok"})
    return
}
