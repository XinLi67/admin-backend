package bootstrap

import (
	"gohub/app/jobs"

	"github.com/robfig/cron/v3"
)

func SetupSchedule() {
	c := cron.New()
	c.AddJob("@every 1m", jobs.GreetingJob{Name: "world!"})
	c.Start()
}
