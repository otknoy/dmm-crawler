package service

import (
	"encoding/json"
	"testing"

	"github.com/otknoy/dmm-crawler/domain/model"
)

type ItemRepositoryMock struct{}

func (r *ItemRepositoryMock) Search(keyword string, hits int, offset int) (model.ItemResponse, error) {
	return model.ItemResponse{}, nil
}

func TestNewItemService(t *testing.T) {
	NewItemService(&ItemRepositoryMock{})
}

func TestGetItems(t *testing.T) {
	s := NewItemService(&ItemRepositoryMock{})

	s.GetItems("keyword", 10, 1)
}

func TestParseActress(t *testing.T) {
	jsonStr := `
{
  "content_id": "hoge",
  "iteminfo": {
    "actress": [
      {
        "id": 123,
        "name": "sample-name"
      },
      {
        "id": "123_ruby",
        "name": "さんぷるねーむ"
      },
      {
        "id": "123_classify",
        "name": "av"
      },
      {
        "id": 1011199,
        "name": "上原亜衣"
      },
      {
        "id": "1011199_ruby",
        "name": "うえはらあい"
      },
      {
        "id": "1011199_classify",
        "name": "av"
      }
    ]
  }
}`

	var dmmItem model.DmmItem
	json.Unmarshal(([]byte)(jsonStr), &dmmItem)

	expected := parseActress(dmmItem)

	actual := []string{"sample-name", "上原亜衣"}

	if expected[0] != actual[0] || expected[1] != actual[1] {
		t.Errorf("fail: expected=%s, actual%s", expected, actual)
	}
}
