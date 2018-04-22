package application

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/otknoy/dmm-crawler/infrastructure"
	"github.com/otknoy/dmm-crawler/model"
	"github.com/otknoy/dmm-crawler/service"
)

type Crawler struct {
	itemSearchService service.ItemService
}

func NewCrawler(itemService service.ItemService) Crawler {
	return Crawler{itemService}
}

func (c *Crawler) Crawl() {
	responses := make(chan model.ItemResponse, 4)
	items := make(chan model.Item)

	go c.fetch(responses)
	go c.process(responses, items)
	c.feed(items)
}

func (c *Crawler) fetch(responses chan<- model.ItemResponse) {
	hits := 100
	offsetLimit := 50000
	for offset := 1; offset <= offsetLimit; offset += hits {
		res, _ := c.itemSearchService.GetItems("", hits, offset)
		responses <- res
	}
	close(responses)
}

func (c *Crawler) process(responses <-chan model.ItemResponse, items chan<- model.Item) {
	ips := service.NewItemProcessService()
	for res := range responses {
		for _, dmmItem := range res.Result.Items {
			item := ips.Process(dmmItem)
			items <- item
		}
	}
	close(items)
}

func (c *Crawler) feed(items <-chan model.Item) {
	solr := infrastructure.NewSolrRepository()
	for item := range items {
		filename := fmt.Sprintf("/mnt/temp/dmm/%s.json", item.ID)
		save(filename, item)

		err := solr.Add(item)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func save(filename string, o interface{}) error {
	bytes, _ := json.Marshal(o)
	err := ioutil.WriteFile(filename, bytes, os.ModePerm)
	if err != nil {
		log.Fatalf("failed to write file: %s", filename)
		return err
	}
	log.Printf("success to save file: %s", filename)

	return nil
}
