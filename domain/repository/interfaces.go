package repository

import "github.com/otknoy/dmm-crawler/domain/model"

type ItemSearcher interface {
	Search(request model.SearchRequest) (model.ItemResponse, error)
}
