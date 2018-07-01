package service

import (
	"github.com/otknoy/dmm-crawler/domain/model"
	"github.com/otknoy/dmm-crawler/domain/repository"
)

type ItemGetServiceImpl struct {
	is repository.ItemSearcher
}

func NewItemGetServiceImpl(is repository.ItemSearcher) (ItemGetService, error) {
	return &ItemGetServiceImpl{is}, nil
}

func (igs *ItemGetServiceImpl) GetItems(keyword string, hits int, offset int, sort string) ([]model.DmmItem, error) {
	response, err := igs.is.Search(keyword, hits, offset, sort)
	return response.Result.Items, err
}
