package application

import (
	"testing"

	"github.com/otknoy/dmm-crawler/domain/interfaces"
	"github.com/otknoy/dmm-crawler/domain/model"
)

type ItemGetServiceMock struct {
	interfaces.ItemGetService
	GetItemsMock func(keyword string, hits int, offset int, rank string) ([]model.DmmItem, error)
}

func (igs *ItemGetServiceMock) GetItems(keyword string, hits int, offset int, rank string) ([]model.DmmItem, error) {
	return igs.GetItemsMock(keyword, hits, offset, rank)
}

type ItemSaveServiceMock struct {
	interfaces.ItemSaveService
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
		GetItemsMock: func(keyword string, hits int, offset int, rank string) ([]model.DmmItem, error) {
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
