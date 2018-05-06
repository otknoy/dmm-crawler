package application

import (
	"github.com/otknoy/dmm-crawler/interfaces"
	"github.com/otknoy/dmm-crawler/model"
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
		dmmItems, _ := dc.igs.GetItems("", "date")

		for _, item := range dmmItems {
			items <- item
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
