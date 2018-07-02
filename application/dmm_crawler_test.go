package application

import (
	"testing"

	"github.com/otknoy/dmm-crawler/domain/model"
	"github.com/otknoy/dmm-crawler/domain/service"
)

type ItemGetServiceMock struct {
	service.ItemGetService
	GetItemsMock func(request model.CrawlRequest) ([]model.DmmItem, error)
}

func (s *ItemGetServiceMock) GetItems(request model.CrawlRequest) ([]model.DmmItem, error) {
	return s.GetItemsMock(request)
}

type ItemSaveServiceMock struct {
	service.ItemSaveService
	SaveItemMock func(filename string, item model.DmmItem) error
}

func (iss *ItemSaveServiceMock) SaveItem(filename string, item model.DmmItem) error {
	return iss.SaveItemMock(filename, item)
}

func TestNewDmmCrawler(t *testing.T) {
	_, err := NewDmmCrawler(
		&ItemGetServiceMock{},
		&ItemSaveServiceMock{},
	)

	if err != nil {
		t.Error("Failed to create new instance")
	}
}

func TestCrawl(t *testing.T) {
	igsMock := &ItemGetServiceMock{
		GetItemsMock: func(request model.CrawlRequest) ([]model.DmmItem, error) {
			return []model.DmmItem{
				model.DmmItem{},
				model.DmmItem{},
			}, nil
		},
	}

	issMock := &ItemSaveServiceMock{
		SaveItemMock: func(filename string, item model.DmmItem) error {
			return nil
		},
	}

	dc, _ := NewDmmCrawler(igsMock, issMock)

	err := dc.Crawl()

	if err != nil {
		t.Error("fail: DmmCrawler.Crawl")
	}
}
