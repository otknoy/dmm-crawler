package application

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/otknoy/dmm-crawler/domain/model"
	"github.com/otknoy/dmm-crawler/domain/service"
	"github.com/otknoy/dmm-crawler/infrastructure"
)

type Crawler struct {
	itemSearchService service.ItemSearchService
}

func NewCrawler() Crawler {
	apiid := os.Getenv("DMM_API_ID")
	affid := os.Getenv("DMM_AFFILIATE_ID")
	dmmSearchRepository := infrastructure.NewDmmSearchRepository(apiid, affid)

	itemSearchService := service.NewItemSearchService(dmmSearchRepository)

	return Crawler{itemSearchService}
}

func (c *Crawler) Crawl() {
	responses := make(chan []model.Item, 4)
	go c.fetch(responses)

	items := make(chan model.Item)
	go c.process(responses, items)

	c.save(items)
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
