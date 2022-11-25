package main

import (
	"os"

	"github.com/plgd-dev/hub/v2/pkg/log"
	"github.com/plgd-dev/hub/v2/resource-aggregate/maintenance"
)

func main() {
	if err := maintenance.PerformMaintenance(); err != nil {
		log.Error(err)
		os.Exit(2)
	}
}