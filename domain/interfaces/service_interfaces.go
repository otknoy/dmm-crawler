package interfaces

import "github.com/otknoy/dmm-crawler/domain/model"

type ItemGetService interface {
	GetItems(keyword string, hits int, offset int, sort string) ([]model.DmmItem, error)
}

type ItemSaveService interface {
	SaveItem(filename string, item model.DmmItem) error
}
