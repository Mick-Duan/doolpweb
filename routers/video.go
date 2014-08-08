package routers

import "doolpweb/models"

// HomeRouter serves home page.
type VideoRouter struct {
    baseRouter
}

// Get implemented Get method for VideoRouter.
func (this *VideoRouter) Get() {
    this.Data["IsVideo"] = true
    this.TplNames = "video.html"

    df := models.GetDoc("screencasts", this.Lang)
    this.Data["Section"] = "video toturial"
    this.Data["Data"] = string(df.Data)
    this.Data["IsHasMarkdown"] = true
}
