package service

import (
	"github.com/otknoy/dmm-crawler/domain/model"
	"github.com/otknoy/dmm-crawler/domain/repository"
)

type ItemService interface {
	GetItems(keyword string, hits int, offset int) ([]model.Item, error)
}

type itemService struct {
	itemSearcher repository.ItemSearcher
}

func NewItemService(itemSearcher repository.ItemSearcher) ItemService {
	return &itemService{
		itemSearcher,
	}
}

func (s *itemService) GetItems(keyword string, hits int, offset int) ([]model.Item, error) {
	itemResponse, _ := s.itemSearcher.Search(keyword, hits, offset)

	items := []model.Item{}
	for _, dmmItem := range itemResponse.Result.Items {
		actresses := parseActress(dmmItem)
		genres := parseGenre(dmmItem)
		makers := parseMaker(dmmItem)

		item := model.Item{
			dmmItem.ContentID,
			dmmItem.Title,
			dmmItem.URL,
			dmmItem.ImageURL.Large,
			actresses,
			genres,
			makers,
		}
		items = append(items, item)
	}

	return items, nil
}

func parseActress(dmmItem model.DmmItem) []string {
	var actresses []string
	for i, actress := range dmmItem.Iteminfo.Actress {
		if i%3 != 0 {
			continue
		}
		actresses = append(actresses, actress.Name)
	}
	return actresses
}

func parseGenre(dmmItem model.DmmItem) []string {
	var genres []string
	for _, genre := range dmmItem.Iteminfo.Genre {
		genres = append(genres, genre.Name)
	}
	return genres
}

func parseMaker(dmmItem model.DmmItem) []string {
	var makers []string
	for _, maker := range dmmItem.Iteminfo.Maker {
		makers = append(makers, maker.Name)
	}
	return makers
}
