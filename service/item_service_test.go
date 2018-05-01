package service

import (
	"testing"

	"github.com/otknoy/dmm-crawler/model"
)

type ItemSearcherMock struct{}

func (r *ItemSearcherMock) Search(keyword string, hits int, offset int) (model.ItemResponse, error) {
	return model.ItemResponse{}, nil
}

func TestNewItemService(t *testing.T) {
	NewItemService(&ItemSearcherMock{})
}

func TestGetItems(t *testing.T) {
	s := NewItemService(&ItemSearcherMock{})

	s.GetItems("keyword", 10, 1)
}
