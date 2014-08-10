// Copyright 2013 Beego Web authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// An open source project for official documentation and blog website of beego app framework.
package main

import (
    "os"

    "github.com/astaxie/beego"
    "github.com/beego/i18n"

    "doolpweb/models"
    "doolpweb/routers"
    "github.com/astaxie/beego/orm"
)

const (
    APP_VER = "0.9.4.0406"
)

// We have to call a initialize function manully
// because we use `bee bale` to pack static resources
// and we cannot make sure that which init() execute first.
func initialize() {
    models.InitModels()

    routers.IsPro = beego.RunMode == "prod"
    if routers.IsPro {
        beego.SetLevel(beego.LevelInfo)
        os.Mkdir("./log", os.ModePerm)
        beego.BeeLogger.SetLogger("file", `{"filename": "log/log"}`)
    }

    routers.InitApp()
}

func main() {

    initialize()

    beego.Info(beego.AppName, APP_VER)

    beego.InsertFilter("/docs/images/:all", beego.BeforeRouter, routers.DocsStatic)

    if !routers.IsPro {
        beego.SetStaticPath("/static_source", "static_source")
        beego.DirectoryIndex = true
    }

    beego.SetStaticPath("/products/images", "products/images/")

    // 开启 ORM，进行注册, 请求测试
    models.RegisterDB()
    orm.Debug = true
    // 自动建表
    orm.RunSyncdb("default", false, true)

    // Register routers.
    beego.Router("/", &routers.HomeRouter{})
    beego.Router("/login", &routers.LoginRouter{})
    beego.Router("/dashboard", &routers.DashboardRouter{})
    beego.Router("/launch", &routers.LaunchRouter{})
    //beego.Router("/dashboard/list", &routers.DashboardRouter{})
    //beego.Router("/community", &routers.CommunityRouter{})
    //beego.Router("/quickstart", &routers.QuickStartRouter{})
    //beego.Router("/video", &routers.VideoRouter{})
    //beego.Router("/products", &routers.ProductsRouter{})
    //beego.Router("/team", &routers.PageRouter{})
    //beego.Router("/about", &routers.AboutRouter{})
    //beego.Router("/donate", &routers.DonateRouter{})
    //beego.Router("/docs/", &routers.DocsRouter{})
    //beego.Router("/docs/:all", &routers.DocsRouter{})
    //beego.Router("/blog", &routers.BlogRouter{})
    //beego.Router("/blog/:all", &routers.BlogRouter{})
    //beego.Router("/login", &routers.LoginRouter{})

    // Register template functions.
    beego.AddFuncMap("i18n", i18n.Tr)
    beego.Errorhandler("404", routers.Page_not_found)

    beego.Run()
}
