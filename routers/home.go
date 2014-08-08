package routers

// HomeRouter serves home page.
type HomeRouter struct {
    baseRouter
}

// Get implemented Get method for HomeRouter.
func (this *HomeRouter) Get() {
    this.Data["IsHome"] = true
    this.TplNames = "home.html"
}
