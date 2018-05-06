package interfaces

import "github.com/otknoy/dmm-crawler/model"

type ItemSearcher interface {
	Search(keyword string, hits int, offset int) (model.ItemResponse, error)
}

type ItemSaver interface {
	Save(key string, data interface{}) error
}

type ItemRepository interface {
	Insert(model.DmmItem) error
}

type ItemPublisher interface {
	Publish(model.DmmItem) error
}
