package routers

// DocsRouter serves about page.
type LaunchRouter struct {
    baseRouter
}

func (this *LaunchRouter) Get() {
    this.TplNames = "launch.html"
}
