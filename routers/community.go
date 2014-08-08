package routers

import (
    "doolpweb/models"
)

// CommunityRouter serves community page.
type CommunityRouter struct {
    baseRouter
}

// Get implemented Get method for CommunityRouter.
func (this *CommunityRouter) Get() {
    this.Data["IsCommunity"] = true
    this.TplNames = "community.html"

    df := models.GetDoc("usecases", this.Lang)
    this.Data["Section"] = "usecases"
    this.Data["Data"] = string(df.Data)
}
