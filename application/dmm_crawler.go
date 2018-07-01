package application

import (
	"log"
	"time"

	"github.com/otknoy/dmm-crawler/domain/interfaces"
	"github.com/otknoy/dmm-crawler/domain/model"
)

type DmmCrawler struct {
	igs interfaces.ItemGetService
	iss interfaces.ItemSaveService
}

func NewDmmCrawler(igs interfaces.ItemGetService, iss interfaces.ItemSaveService) (DmmCrawler, error) {
	return DmmCrawler{igs, iss}, nil
}

func (dc *DmmCrawler) Crawl() error {
	items := make(chan model.DmmItem)

	// get items
	go func() {
		hits := 100
		offsetLimit := 50000
		sortList := []string{"date", "rank", "review"}

		for offset := 1; offset <= offsetLimit; offset += hits {
			go func() {
				for _, sort := range sortList {
					dmmItems, _ := dc.igs.GetItems("", hits, offset, sort)
					log.Printf("[get]\tget %d items\n", len(dmmItems))
					for _, item := range dmmItems {
						log.Printf("[get]\t%s\n", item.ContentID)
						items <- item
					}
				}
			}()

			<-time.NewTicker(1 * time.Second).C
		}
		close(items)
	}()

	// save items
	for item := range items {
		filename := item.ContentID + ".json"
		dc.iss.SaveItem(filename, item)
	}

	return nil
}
