package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/garyburd/redigo/redis"
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
		panic(err)
	}

	items := make(chan model.DmmItem, 4)

	go crawler.Crawl(items)

	for item := range items {
		itemRepository.Insert(item)
		itemPublisher.Publish(item)
	}
}

func subscriber() {
	log.Println("Subscriber")

	conn, _ := redis.Dial("tcp", "localhost:6379")
	psc := redis.PubSubConn{Conn: conn}
	psc.Subscribe("dmm-items")
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			item := &model.DmmItem{}
			if err := json.Unmarshal(v.Data, item); err != nil {
				fmt.Println(err)
			}

			// fmt.Println(string(v.Data))
			fmt.Println(item.Title)
			// fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Println("error")
		}
	}
}
