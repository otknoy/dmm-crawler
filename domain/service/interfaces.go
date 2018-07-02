package service

import "github.com/otknoy/dmm-crawler/domain/model"

type ItemGetService interface {
	GetItems(request model.CrawlRequest) ([]model.DmmItem, error)
}

type ItemSaveService interface {
	SaveItem(filename string, item model.DmmItem) error
}
