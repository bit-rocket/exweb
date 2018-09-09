package main

// modified from iris/_examples/overview/main.go

import (
    "github.com/kataras/iris"
    "github.com/yeweishuai/exweb/order"
)

func main() {
    app := iris.New()
    // app.Logger().SetLevel("disable") to disable the logger

    // Define templates using the std html/template engine.
    // Parse and load all files inside "./views" folder with ".html" file extension.
    // Reload the templates on each request (development mode).
    app.RegisterView(iris.HTML("./layuicms", ".html").Reload(true))

    // Register custom handler for specific http errors.
    app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
        // .Values are used to communicate between handlers, middleware.
        errMessage := ctx.Values().GetString("error")
        if errMessage != "" {
            ctx.Writef("Internal server error: %s", errMessage)
            return
        }

        ctx.Writef("(Unexpected) internal server error")
    })

    app.Use(func(ctx iris.Context) {
        ctx.Application().Logger().Infof("Begin request for path: %s", ctx.Path())
        ctx.Next()
    })
    // app.Done(func(ctx iris.Context) {]})

    orderRoutes := app.Party("order", logThisMiddleware)
    // flowing is code block, not related above as a function
    {
        // set handler for: /order/todeal
        orderRoutes.Get("/todeal", order.ToDealOrders)
    }

    app.StaticWeb("layui", "layuicms/layui")
    app.StaticWeb("js", "layuicms/js")
    app.StaticWeb("json", "layuicms/json")
    app.StaticWeb("css", "layuicms/css")
    app.StaticWeb("page", "layuicms/page")
    app.StaticWeb("images", "layuicms/images")

    app.Get("/", func(ctx iris.Context) {
        ctx.Header("Cache-Control", "no-cache")
        ctx.View("index.html")
    })

    // Listen for incoming HTTP/1.x & HTTP/2 clients on localhost port 8080.
    app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"), iris.WithoutVersionChecker)
}

func logThisMiddleware(ctx iris.Context) {
    ctx.Application().Logger().Infof("Path: %s | IP: %s", ctx.Path(), ctx.RemoteAddr())

    // .Next is required to move forward to the chain of handlers,
    // if missing then it stops the execution at this handler.
    ctx.Next()
}
