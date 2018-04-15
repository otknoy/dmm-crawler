package main

import (
	"os"

	"github.com/otknoy/dmm-crawler/application"
	"github.com/otknoy/dmm-crawler/domain/repository"
	"github.com/otknoy/dmm-crawler/domain/service"
)

func main() {
	apiid := os.Getenv("DMM_API_ID")
	affid := os.Getenv("DMM_AFFILIATE_ID")
	dmmItemRepository := repository.NewDmmItemRepository(apiid, affid)

	itemService := service.NewItemService(dmmItemRepository)

	crawler := application.NewCrawler(itemService)
	crawler.Crawl()
}
