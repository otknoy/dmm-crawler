package service

import (
	"github.com/otknoy/dmm-crawler/interfaces"
	"github.com/otknoy/dmm-crawler/model"
)

type ItemService interface {
	GetItems(keyword string, hits int, offset int) ([]model.Item, error)
}

type itemService struct {
	itemSearcher interfaces.ItemSearcher
}

func NewItemService(itemSearcher interfaces.ItemSearcher) ItemService {
	return &itemService{
		itemSearcher,
	}
}

func (s *itemService) GetItems(keyword string, hits int, offset int) ([]model.Item, error) {
	itemResponse, _ := s.itemSearcher.Search(keyword, hits, offset)

	items := []model.Item{}
	for _, dmmItem := range itemResponse.Result.Items {

		item := process(dmmItem)
		items = append(items, item)
	}

	return items, nil
}
