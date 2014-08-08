package routers

import (
    "io"
    "os"

    "github.com/astaxie/beego/context"
    "github.com/beego/i18n"

    "doolpweb/models"
)

// DocsRouter serves about page.
type DocsRouter struct {
    baseRouter
}

// Get implemented Get method for DocsRouter.
func (this *DocsRouter) Get() {
    this.Data["IsDocs"] = true
    this.TplNames = "docs.html"

    dRoot := models.GetDocByLocale(this.Lang)

    if dRoot == nil {
        this.Abort("404")
        return
    }

    link := this.GetString(":all")

    var doc *models.DocNode
    if len(link) == 0 {
        if dRoot.Doc.HasContent() {
            doc = dRoot.Doc
        } else {
            this.Redirect("/docs/intro/", 302)
            return
        }
    } else {
        doc, _ = dRoot.GetNodeByLink(link)
        if doc == nil {
            doc, _ = dRoot.GetNodeByLink(link + "/")
            if doc != nil {
                this.Redirect("/docs/"+link+"/", 301)
                return
            }
        }
    }

    if doc == nil {
        this.Abort("404")
        return
    }

    this.Data["DocRoot"] = dRoot
    this.Data["Doc"] = doc
    this.Data["Title"] = doc.Name
    this.Data["Data"] = doc.GetContent()
}

func DocsStatic(ctx *context.Context) {
    if uri := ctx.Input.Params[":all"]; len(uri) > 0 {
        lang := ctx.GetCookie("lang")
        if !i18n.IsExist(lang) {
            lang = "en-US"
        }

        f, err := os.Open("docs/" + lang + "/" + "images/" + uri)
        if err != nil {
            ctx.WriteString(err.Error())
            return
        }
        defer f.Close()

        _, err = io.Copy(ctx.ResponseWriter, f)
        if err != nil {
            ctx.WriteString(err.Error())
            return
        }
    }
}
