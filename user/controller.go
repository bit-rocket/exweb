package user

import (
//    "fmt"

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
    Ctx         iris.Context
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
        username = form(comm.UsernameKey)
        password = form(comm.UserpassKey)
    )
    uc.Ctx.Application().Logger().Infof("ip[%s] try login user[%s] pass[%s]",
            uc.Ctx.RemoteAddr(), username, password)
    if comm.GConf.UserConf[username] != password {
        return comm.UserLoginView
    }

    uc.Session.Set(comm.UsernameKey, username)
    return comm.PathIndex
}

func (uc *UController) AnyLogout() mvc.Result {
    uc.Ctx.Application().Logger().Infof("ip[%s], user[%s] logout",
            uc.Ctx.RemoteAddr(), uc.Session.GetString("username"))
    uc.Session.Destroy()
    return comm.UserLoginView
}
