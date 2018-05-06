package service

import (
	"github.com/otknoy/dmm-crawler/domain/interfaces"
	"github.com/otknoy/dmm-crawler/domain/model"
)

type ItemGetService struct {
	is interfaces.ItemSearcher
}

func NewItemGetService(is interfaces.ItemSearcher) (interfaces.ItemGetService, error) {
	return &ItemGetService{is}, nil
}

func (igs *ItemGetService) GetItems(keyword string, hits int, offset int, sort string) ([]model.DmmItem, error) {
	response, err := igs.is.Search(keyword, hits, offset, sort)
	return response.Result.Items, err
}
