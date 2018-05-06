package interfaces

import "github.com/otknoy/dmm-crawler/model"

type ItemGetService interface {
	GetItems(keyword string, hits int, offset int, rank string) ([]model.DmmItem, error)
}

type ItemSaveService interface {
	SaveItem(filename string, item model.DmmItem) error
}
