package routers

import (
    "strings"

    "doolpweb/models"
)

// SamplesRouter serves about page.
type SamplesRouter struct {
    baseRouter
}

// Get implemented Get method for SamplesRouter.
func (this *SamplesRouter) Get() {
    this.Data["IsSamples"] = true

    reqUrl := this.Ctx.Request.URL.String()
    sec := reqUrl[strings.LastIndex(reqUrl, "/")+1:]
    if qm := strings.Index(sec, "?"); qm > -1 {
        sec = sec[:qm]
    }

    if len(sec) == 0 || sec == "samples" {
        this.Redirect("/samples/Samples_Introduction", 302)
        return
    } else {
        this.Data[sec] = true
    }

    // Get language.
    curLang, _ := this.Data["LangVer"].(langType)
    df := models.GetDoc(sec, curLang.Lang)
    if df == nil {
        this.Redirect("/samples/Samples_Introduction", 302)
        return
    }

    this.Data["Title"] = df.Title
    this.Data["Data"] = string(df.Data)
    this.Data["IsHasMarkdown"] = true
    this.TplNames = "samples_" + curLang.Lang + ".html"
}
