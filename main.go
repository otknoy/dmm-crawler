package main

import (
	"os"

	"github.com/otknoy/dmm-crawler/application"
	"github.com/otknoy/dmm-crawler/infrastructure"
	"github.com/otknoy/dmm-crawler/service"
)

func main() {
	apiid := os.Getenv("DMM_API_ID")
	affid := os.Getenv("DMM_AFFILIATE_ID")
	dmmItemRepository := infrastructure.NewDmmItemRepository(apiid, affid)

	itemService := service.NewItemService(dmmItemRepository)

	crawler := application.NewCrawler(itemService)
	crawler.Crawl()
}
