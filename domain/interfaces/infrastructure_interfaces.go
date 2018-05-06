package interfaces

import "github.com/otknoy/dmm-crawler/domain/model"

type ItemSearcher interface {
	Search(keyword string, hits int, offset int, sort string) (model.ItemResponse, error)
}
