package application

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/otknoy/dmm-crawler/model"
	"github.com/otknoy/dmm-crawler/service"
)

type Crawler struct {
	itemSearchService service.ItemService
}

func NewCrawler(itemService service.ItemService) Crawler {
	return Crawler{itemService}
}

func (c *Crawler) Crawl(items chan<- model.DmmItem) {
	responses := make(chan model.ItemResponse, 2)

	go c.fetch(responses)

	for r := range responses {
		for _, dmmItem := range r.Result.Items {
			items <- dmmItem
		}
	}

	close(items)
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
