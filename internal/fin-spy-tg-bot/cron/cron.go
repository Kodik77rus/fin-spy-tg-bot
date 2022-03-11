package cron

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var wg sync.WaitGroup

func (c *Cron) Start() {
	defer c.cron.Stop()

	c.cron.AddFunc("10 0  1 * *", c.checkAssets)  //At 00:10 every month
	c.cron.AddFunc("* 31 * * *", assetsScanning)  //every  31 min
	c.cron.AddFunc("* 21 * * *", assetsScanning)  //every  21 min
	c.cron.AddFunc("*/10 * * * *", c.checkAssets) //every  16 min
	c.cron.AddFunc("*/1 * * * *", assetsScanning) //every  1 min

	// start scheduler
	c.cron.Start()

	// shutdown trigger
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}

func (c *Cron) checkAssets() {
	c.logger.Info("Starting update markets")

	markets, err := c.storage.GetAllMarkets()
	if err != nil {
		c.logger.Panicf("Troubls with db: %v", err)
	}

	wg.Add(len(*markets))

	for _, m := range *markets {
		go c.updateAssets(m.Code)
	}

	wg.Wait()

	c.logger.Info("Markets update completed successfully")
}

func assetsScanning() {
	fmt.Printf("%v", time.Now())
}

func (c *Cron) updateAssets(market string) error {
	defer wg.Done()

	c.logger.Infof("Starting update %s market", market)

	time.Sleep(3 * time.Second) // time  delay

	companies := new([]finHubResponse)

	res, err := c.finHubClient.R().
		SetQueryParam("exchange", market).
		Get("/stock/symbol")
	if err != nil {
		c.logger.Panicf("Something went wrong with finhub api: %v", err)
	}

	err = json.Unmarshal(res.Body(), companies)
	if err != nil {

		c.logger.Errorf("%v, %v, \n%+v", err, market, res)
		os.Exit(1)
	}
	fmt.Print(companies)
	return nil
}
