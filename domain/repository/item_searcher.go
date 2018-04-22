package repository

import "github.com/otknoy/dmm-crawler/domain/model"

type ItemSearcher interface {
	Search(keyword string, hits int, offset int) (model.ItemResponse, error)
}
