package service

import (
	"github.com/otknoy/dmm-crawler/interfaces"
	"github.com/otknoy/dmm-crawler/model"
)

type ItemGetService struct {
	is interfaces.ItemSearcher
}

func NewItemGetService(is interfaces.ItemSearcher) (interfaces.ItemGetService, error) {
	return &ItemGetService{is}, nil
}

func (igs *ItemGetService) GetItems(keyword string, hits int, offset int, rank string) ([]model.DmmItem, error) {
	response, err := igs.is.Search(keyword, hits, offset)
	return response.Result.Items, err
}
