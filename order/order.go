package order

import (
    "fmt"
    "github.com/kataras/iris"
)

type Users struct {
    UserId      string      `json:"usersId"`
    UserName    string      `json:"userName"`
    UserEmail   string      `json:"userEmail"`
    UserSex     string      `json:"userSex"`
    UserStatus  string      `json:"userStatus"`
    UserGrade   string      `json:"userGrade"`
    UserEndTime string      `json:"userEndTime"`
}

func ToDealOrders(ctx iris.Context) {
    var user []Users
    peter := Users{
        UserId: "1",
        UserName: "abc",
        UserEmail: "abc@xcd.com",
        UserSex: "male",
        UserStatus: "buok",
        UserGrade: "middle",
        UserEndTime: "2017-05-10 10:30",
    }
    user = append(user, peter)
    ctx.JSON(user)
}

