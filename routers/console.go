package routers

import (
    "doolpweb/models"
    "github.com/astaxie/beego"
)

// DocsRouter serves about page.
type ConsoleRouter struct {
    baseRouter
}

func (this *ConsoleRouter) Get() {
    if !checkAccount(this.Ctx) {
        this.Redirect("/login", 301)
    }
    this.Data["IsLogin"] = checkAccount(this.Ctx)
    this.Data["IsConsole"] = true

    op := this.Input().Get("op")

    switch op {
    case "add":
        sname := this.Input().Get("sname")
        if len(sname) == 0 {
            break
        }

        err := models.AddServer(sname)
        if err != nil {
            beego.Error(err)
        }

        this.Redirect("/console", 301)
        return

    case "del":
        id := this.Input().Get("id")
        if len(id) == 0 {
            break
        }
        err := models.DelServer(id)
        if err != nil {
            beego.Error(err)
        }
        this.Redirect("/console", 301)
        return
    }
    this.TplNames = "console.html"
    var err error
    this.Data["Servers"], err = models.GetAllServerNmaes()

    if err != nil {
        beego.Error(err)
    }
}
