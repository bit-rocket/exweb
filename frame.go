package main

import (

    "os"
    "fmt"
    "io/ioutil"

    "github.com/kataras/iris"
	"github.com/BurntSushi/toml"

    "github.com/yeweishuai/exweb/comm"
)

func LoadToml(conf string) (err error) {
    if _, err = os.Stat(conf); err != nil {
        err = fmt.Errorf("try stat conf[%s] error:%s", conf, err.Error())
        return err
    }

    data, err := ioutil.ReadFile(conf)
    if err != nil {
        return err
    }

    if _, err := toml.Decode(string(data), &comm.GConf); err != nil {
        return err
    }
    return
}

func ProcessInit(conf string) (err error) {
    err = LoadToml(conf)
    if err != nil {
        return
    }

    comm.OrderTypeMap[0] = "dft"
    comm.OrderTypeMap[1] = "buy2sell"
    comm.OrderTypeMap[2] = "sell2buy"

    comm.StatusMap[0] = "created"
    comm.StatusMap[1] = "buying"
    comm.StatusMap[2] = "holding"
    comm.StatusMap[3] = "selling"
    comm.StatusMap[4] = "quit"
    comm.StatusMap[5] = "finish"

    comm.SonStatusMap[0] = "created"
    comm.SonStatusMap[1] = "made"
    comm.SonStatusMap[2] = "dealt"
    comm.SonStatusMap[3] = "undo"
    comm.SonStatusMap[4] = "finish"
    return
}

func IndexHandler(ctx iris.Context) {
        session := comm.GSession.Start(ctx)
        exist := session.GetString("username")
        if exist == "" {
            ctx.Application().Logger().Infof("ip[%s] not login, redirect",
                    ctx.RemoteAddr())
            ctx.Redirect(comm.StrUserLogin)
            return
        }
        ctx.Header("Cache-Control", "no-cache")
        ctx.View("index.html")
}
