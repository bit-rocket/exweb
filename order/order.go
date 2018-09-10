package order

import (
    "github.com/kataras/iris"
)

type Users struct {
    UserId      string      `json:"usersId"`
    ExName      string      `json:"exName"`
    TradingPair string      `json:"tradingPair"`
    Price       string      `json:"price"`
    Holding     string      `json:"holding"`
    Earn        string      `json:"earn"`
    OrderTime   string      `json:"orderTime"`
}

func ToDealOrders(ctx iris.Context) {
    var user []Users
    peter := Users{
        UserId: "1",
        ExName: "okex",
        TradingPair: "eos/usdt",
        Price: "5.01",
        Holding: "+1.02%",
        Earn: "20",
        OrderTime: "2017-05-10 10:30",
    }
    user = append(user, peter)
    ctx.JSON(user)
}

