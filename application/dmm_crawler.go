package application

import (
	"github.com/otknoy/dmm-crawler/domain/model"
	"github.com/otknoy/dmm-crawler/domain/service"
)

type DmmCrawler struct {
	igs service.ItemGetService
	iss service.ItemSaveService
}

func NewDmmCrawler(igs service.ItemGetService, iss service.ItemSaveService) (DmmCrawler, error) {
	return DmmCrawler{igs, iss}, nil
}

func (dc *DmmCrawler) Crawl() error {
	items := make(chan model.DmmItem)

	// get items
	go func() {
		n := 100
		pageLimit := 100
		for page := 1; page <= pageLimit; page += 1 {
			request := model.NewCrawlRequest("", uint(n), uint(page), "date")
			dmmItems, _ := dc.igs.GetItems(request)
			for _, item := range dmmItems {
				items <- item
			}

			// dmmItems, _ = dc.igs.GetItems("", hits, offset, "rank")
			// for _, item := range dmmItems {
			// 	items <- item
			// }

			// dmmItems, _ = dc.igs.GetItems("", hits, offset, "review")
			// for _, item := range dmmItems {
			// 	items <- item
			// }
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
