package main

import (
	"log"
	"os"

	"github.com/otknoy/dmm-crawler/application"
	"github.com/otknoy/dmm-crawler/infrastructure"
	"github.com/otknoy/dmm-crawler/service"
)

func main() {
	apiid := os.Getenv("DMM_API_ID")
	affid := os.Getenv("DMM_AFFILIATE_ID")
	outputDir := os.Getenv("OUTPUT_DIR")

	dmmItemRepository := infrastructure.NewDmmItemRepository(apiid, affid)
	itemGetService, _ := service.NewItemGetService(dmmItemRepository)

	itemSaveService, _ := service.NewItemSaveService(outputDir)

	dmmCrawler, _ := application.NewDmmCrawler(itemGetService, itemSaveService)

	err := dmmCrawler.Crawl()
	if err != nil {
		log.Fatal(err)
	}
}
