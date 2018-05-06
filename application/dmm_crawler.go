package application

import (
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
		for offset := 1; offset <= offsetLimit; offset += hits {
			dmmItems, _ := dc.igs.GetItems("", hits, offset, "date")
			for _, item := range dmmItems {
				items <- item
			}

			dmmItems, _ = dc.igs.GetItems("", hits, offset, "rank")
			for _, item := range dmmItems {
				items <- item
			}

			dmmItems, _ = dc.igs.GetItems("", hits, offset, "review")
			for _, item := range dmmItems {
				items <- item
			}
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
