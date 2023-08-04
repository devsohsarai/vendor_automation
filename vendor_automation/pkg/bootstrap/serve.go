package bootstrap

import (
	"github.com/gowaves/vendor_automation/pkg/cif"
	"github.com/gowaves/vendor_automation/pkg/database"
	"github.com/gowaves/vendor_automation/pkg/htmlparse"
	"github.com/gowaves/vendor_automation/pkg/routing"
	"github.com/gowaves/vendor_automation/pkg/sessions"
	"github.com/gowaves/vendor_automation/pkg/static"
)

func Serve() {
	cif.Set()

	database.Connect()

	routing.Init()

	sessions.Start(routing.GetRouter())

	static.LoadStatic(routing.GetRouter())

	htmlparse.LoadHTML(routing.GetRouter())

	routing.RegisterRoutes()

	routing.Serve()
}
