package comm

import (
//    "github.com/kataras/iris"
)

type DBConfig struct {
    DBUser              string  `toml:"DBUser"`
    DBPass              string  `toml:"DBPass"`
    DBName              string  `toml:"DBName"`
}

type GlobalConfig struct {
    DbConf              DBConfig    `toml:"DbConf"`
}

var (
    GConf              GlobalConfig
)
