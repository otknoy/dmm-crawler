package infrastructure

import (
	"encoding/json"
	"strconv"

	"github.com/gomodule/redigo/redis"

	"github.com/otknoy/dmm-crawler/interfaces"
	"github.com/otknoy/dmm-crawler/model"
)

type ItemPublisher struct {
	redisHost string
	redisPort int
	redisConn redis.Conn
}

func NewItemPublisher(host string, port int) (interfaces.ItemPublisher, error) {
	conn, err := redis.Dial("tcp", host+":"+strconv.Itoa(port))
	if err != nil {
		return &ItemPublisher{}, err
	}

	return &ItemPublisher{
		host,
		port,
		conn,
	}, nil
}

func (ip *ItemPublisher) Publish(item model.DmmItem) error {
	bytes, err := json.Marshal(item)
	if err != nil {
		return err
	}

	_, err = ip.redisConn.Do("PUBLISH", "dmm-items", string(bytes))
	if err != nil {
		return err
	}

	return nil
}
