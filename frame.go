package main

import (

    "os"
    "fmt"
    "io/ioutil"

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
    return
}
