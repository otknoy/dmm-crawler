package main

import (
	"log"
	"os"

	"github.com/otknoy/dmm-crawler/application"
	"github.com/otknoy/dmm-crawler/domain/service"
	"github.com/otknoy/dmm-crawler/infrastructure"
)

func main() {
	apiid := os.Getenv("DMM_API_ID")
	affid := os.Getenv("DMM_AFFILIATE_ID")
	outputDir := os.Getenv("OUTPUT_DIR")

	dmmItemRepository := infrastructure.NewDmmItemRepository(apiid, affid)
	itemGetService, _ := service.NewItemGetServiceImpl(dmmItemRepository)

	itemSaveService, _ := service.NewItemSaveServiceImpl(outputDir)

	dmmCrawler, _ := application.NewDmmCrawler(itemGetService, itemSaveService)

	err := dmmCrawler.Crawl()
	if err != nil {
		log.Fatal(err)
	}
}
