package routing

import (
	"fmt"
	"log"

	"github.com/gowaves/vendor_automation/pkg/cif"
)

func Serve() {
	r := GetRouter()
	configs := cif.Get()
	err := r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))

	if err != nil {
		log.Fatal("Error in rotuing")
		return
	}
}
