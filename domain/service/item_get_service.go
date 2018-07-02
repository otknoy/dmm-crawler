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

func (igs *ItemGetServiceImpl) GetItems(crawlRequest model.CrawlRequest) ([]model.DmmItem, error) {
	searchRequest := model.NewSearchRequest(
		crawlRequest.GetKeyword(),
		crawlRequest.GetN(),
		crawlRequest.GetOffset(),
		crawlRequest.GetSort(),
	)

	response, err := igs.is.Search(searchRequest)
	return response.Result.Items, err
}
