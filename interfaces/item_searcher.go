package interfaces

import "github.com/otknoy/dmm-crawler/model"

type ItemSearcher interface {
	Search(keyword string, hits int, offset int) (model.ItemResponse, error)
}
