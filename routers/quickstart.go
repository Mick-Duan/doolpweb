package routers

import (
    "doolpweb/models"
)

// QuickStartRouter serves about page.
type QuickStartRouter struct {
    baseRouter
}

// Get implemented Get method for QuickStartRouter.
func (this *QuickStartRouter) Get() {
    this.Data["IsQuickStart"] = true
    this.TplNames = "quickstart.html"

    df := models.GetDoc("quickstart", this.Lang)
    this.Data["Section"] = "quickstart"
    this.Data["Title"] = df.Title
    this.Data["Data"] = string(df.Data)
    this.Data["IsHasMarkdown"] = true
}
