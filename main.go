package main

import (
	"log"
	"os"
	"time"

	"github.com/otknoy/dmm-crawler/application"
	"github.com/otknoy/dmm-crawler/infrastructure"
	"github.com/otknoy/dmm-crawler/model"
	"github.com/otknoy/dmm-crawler/service"
)

func main() {
	apiid := os.Getenv("DMM_API_ID")
	affid := os.Getenv("DMM_AFFILIATE_ID")
	dmmItemRepository := infrastructure.NewDmmItemRepository(apiid, affid)
	itemService := service.NewItemService(dmmItemRepository)
	crawler := application.NewCrawler(itemService)

	itemRepository := infrastructure.NewItemRepository()
	itemPublisher, err := infrastructure.NewItemPublisher("localhost", 6379)
	if err != nil {
		log.Fatal(err)
	}

	items := make(chan model.DmmItem, 4)

	go crawler.Crawl(items)

	count := 0

	for item := range items {
		itemRepository.Insert(item)
		itemPublisher.Publish(item)

		time.Sleep(10 * time.Millisecond)

		count++
		log.Println(count)
	}
}
