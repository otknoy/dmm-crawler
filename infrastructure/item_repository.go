package infrastructure

import (
	"github.com/otknoy/dmm-crawler/interfaces"
	"github.com/otknoy/dmm-crawler/model"
)

type ItemRepository struct {
}

func NewItemRepository() interfaces.ItemRepository {
	return &ItemRepository{}
}

func (ir *ItemRepository) Insert(model.DmmItem) error {
	return nil
}
