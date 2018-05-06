package application

import (
	"log"
	"path/filepath"

	"github.com/otknoy/dmm-crawler/model"
	"github.com/otknoy/dmm-crawler/service"
)

type Crawler struct {
	itemSearchService service.ItemService
}

func NewCrawler(itemService service.ItemService) Crawler {
	return Crawler{itemService}
}

func (c *Crawler) Crawl(outputDir string) {
	responses := make(chan model.ItemResponse, 2)
	items := make(chan model.DmmItem, 100)

	go c.fetch(responses)

	go process(responses, items)

	for item := range items {
		filename := filepath.Join(outputDir, item.ContentID+".json")
		err := save(filename, item)
		if err != nil {
			log.Print(err)
		}
	}
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

func process(responses <-chan model.ItemResponse, items chan<- model.DmmItem) {
	for res := range responses {
		for _, item := range res.Result.Items {
			items <- item
		}
	}
	close(items)
}
func save(filename string, o interface{}) error {

}
