package cron

import (
	"github.com/jasonlvhit/gocron"
)

type Cron struct {
	cron     *gocron.Scheduler
	fn       *func() error
	interval uint
}

func New(fn *func() error, interval uint) *Cron {
	return &Cron{
		cron:     gocron.NewScheduler(),
		fn:       fn,
		interval: interval,
	}
}
