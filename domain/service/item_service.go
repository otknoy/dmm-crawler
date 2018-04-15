package service

import (
	"github.com/otknoy/dmm-crawler/domain/model"
	"github.com/otknoy/dmm-crawler/domain/repository"
)

type ItemService interface {
	GetItems(keyword string, hits int, offset int) ([]model.Item, error)
}

type itemService struct {
	itemSearchRepository repository.ItemRepository
}

func NewItemService(itemRepository repository.ItemRepository) ItemService {
	return &itemService{
		itemRepository,
	}
}

func (s *itemService) GetItems(keyword string, hits int, offset int) ([]model.Item, error) {
	itemResponse, _ := s.itemSearchRepository.Search(keyword, hits, offset)

	items := []model.Item{}
	for _, dmmItem := range itemResponse.Result.Items {
		var actresses []string
		for i, actress := range dmmItem.Iteminfo.Actress {
			if i != 0 {
				continue
			}
			actresses = append(actresses, actress.Name)
		}

		var genres []string
		for _, genre := range dmmItem.Iteminfo.Genre {
			genres = append(genres, genre.Name)
		}

		var makers []string
		for _, maker := range dmmItem.Iteminfo.Maker {
			makers = append(makers, maker.Name)
		}

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
