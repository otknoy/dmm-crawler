package infrastructure

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/otknoy/dmm-crawler/interfaces"
	"github.com/otknoy/dmm-crawler/model"
)

type ItemRepository struct {
	session *mgo.Session
}

func NewItemRepository() (interfaces.ItemRepository, error) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		return &ItemRepository{}, err
	}

	return &ItemRepository{session}, nil
}

func (ir *ItemRepository) Insert(item model.DmmItem) error {
	c := ir.session.DB("dmm").C("items")
	selector := bson.M{"contentid": item.ContentID}
	ci, err := c.Upsert(selector, item)
	if err != nil {
		return err
	}

	log.Printf("upsert: updated=%d, removed=%d", ci.Updated, ci.Removed)

	return nil
}
