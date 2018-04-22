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
	responses := make(chan []model.Item, 4)
	go c.fetch(responses)

	items := make(chan model.Item)
	go c.process(responses, items)

	// c.save(items)

	solr := infrastructure.NewSolrRepository()
	for item := range items {
		err := solr.Add(item)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func (c *Crawler) fetch(responses chan []model.Item) {
	hits := 100
	offsetLimit := 50000
	for offset := 1; offset <= offsetLimit; offset += hits {
		items, _ := c.itemSearchService.GetItems("", hits, offset)
		responses <- items
	}
	close(responses)
}

func (c *Crawler) process(in chan []model.Item, out chan model.Item) {
	for items := range in {
		for _, item := range items {
			out <- item
		}
	}
	close(out)
}

func (c *Crawler) save(items chan model.Item) {
	for item := range items {
		bytes, _ := json.Marshal(item)

		outputDir := "/mnt/temp/dmm/"
		filename := outputDir + fmt.Sprintf("%s.json", item.ID)
		err := ioutil.WriteFile(filename, bytes, os.ModePerm)
		if err != nil {
			log.Fatalf("failed to write file: %s", filename)
		}
		// log.Printf("success to save file: %s", filename)
	}
}
