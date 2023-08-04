package bootstrap

import (
	"github.com/gowaves/vendor_automation/internal/database/migration"
	"github.com/gowaves/vendor_automation/pkg/cif"
	"github.com/gowaves/vendor_automation/pkg/database"
)

func Migrate() {
	cif.Set()
	database.Connect()
	migration.Migrate()
}
