package service

import (
	"github.com/otknoy/dmm-crawler/interfaces"
	"github.com/otknoy/dmm-crawler/model"
)

type ItemService interface {
	GetItems(keyword string, hits int, offset int) (model.ItemResponse, error)
}

type itemService struct {
	itemSearcher interfaces.ItemSearcher
}

func NewItemService(itemSearcher interfaces.ItemSearcher) ItemService {
	return &itemService{
		itemSearcher,
	}
}

func (s *itemService) GetItems(keyword string, hits int, offset int) (model.ItemResponse, error) {
	itemResponse, err := s.itemSearcher.Search(keyword, hits, offset)
	return itemResponse, err
}
