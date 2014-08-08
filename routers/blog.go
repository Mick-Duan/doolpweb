package routers

import (
    "strings"

    "doolpweb/models"
)

// BlogRouter serves about page.
type BlogRouter struct {
    baseRouter
}

// Get implemented Get method for BlogRouter.
func (this *BlogRouter) Get() {
    this.Data["IsBlog"] = true
    this.TplNames = "blog.html"

    reqUrl := this.Ctx.Request.URL.String()
    fullName := reqUrl[strings.LastIndex(reqUrl, "/")+1:]
    if qm := strings.Index(fullName, "?"); qm > -1 {
        fullName = fullName[:qm]
    }

    df := models.GetBlog(fullName, this.Lang)
    if df == nil {
        this.Redirect("/blog", 302)
        return
    }

    this.Data["Title"] = df.Title
    this.Data["Data"] = string(df.Data)
    this.Data["IsHasMarkdown"] = true
}
