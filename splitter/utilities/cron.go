package utilities

import (
	"github.com/robfig/cron/v3"
)

var Scheduler *cron.Cron

func LaunchCRON() {
	Scheduler = cron.New()
	Scheduler.AddFunc(
		"0 * * * * *",
		func() {

			Scheduler.Stop()
		},
	)
}
