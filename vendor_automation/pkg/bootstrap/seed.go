package bootstrap

import (
	"github.com/gowaves/vendor_automation/internal/database/seeder"
	"github.com/gowaves/vendor_automation/pkg/cif"
	"github.com/gowaves/vendor_automation/pkg/database"
)

func Seed() {
	cif.Set()
	database.Connect()
	seeder.Seed()
}
