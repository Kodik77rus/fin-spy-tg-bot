package cron

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"time"

	"github.com/robfig/cron/v3"
)

type Cron struct {
	cron        *cron.Cron
	finhubToken string
}

func New(finhubToken string) *Cron {
	moscow, _ := time.LoadLocation("Europe/Moscow")

	return &Cron{
		cron:        cron.New(cron.WithLocation(moscow)),
		finhubToken: finhubToken,
	}
}

func (c *Cron) Start() {
	defer c.cron.Stop()

	c.cron.AddFunc("10 0  1 * *", CheckAssets)     //At 00:10 every month
	c.cron.AddFunc("*/31 * * * *", assetsScanning) //every  31 min
	c.cron.AddFunc("*/21 * * * *", assetsScanning) //every  21 min
	c.cron.AddFunc("*/16 * * * *", assetsScanning) //every  16 min
	c.cron.AddFunc("*/1 * * * *", assetsScanning)  //every  1 min

	// start scheduler
	go c.cron.Start()

	// shutdown trigger
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}

func CheckAssets() {
	fmt.Printf("%v", time.Now())
}

func assetsScanning() {
	fmt.Printf("%v", time.Now())
}
