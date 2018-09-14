package user

import (
    "fmt"

    "github.com/kataras/iris"
    "github.com/kataras/iris/mvc"
    "github.com/kataras/iris/sessions"

    "github.com/yeweishuai/exweb/comm"
)

type UserForm struct {
    Username    string
    Password    string
}

type UController struct {
    Session     *sessions.Session
}

type formValue func(string) string

func (c *UController) BeforeActivation(b mvc.BeforeActivation) {
	b.Dependencies().Add(func(ctx iris.Context) formValue { return ctx.FormValue })
}

func (uc *UController) GetLogin() mvc.Result {
    return comm.UserLoginView
}

func (uc *UController) PostLogin(form formValue) mvc.Result {
	var (
		username = form("username")
		password = form("password")
	)
    fmt.Println("get user pass:", username, password)
    if comm.GConf.UserConf[username] != password {
        return comm.UserLoginView
    }

    uc.Session.Set("username", username)
    return comm.PathIndex
}
