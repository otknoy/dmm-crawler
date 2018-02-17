package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-redis/redis"
	"github.com/otknoy/dmm-crawler/dmm"
)

func main() {
	apiid := os.Getenv("DMM_API_ID")
	affid := os.Getenv("DMM_AFFILIATE_ID")

	itemSearchClient := dmm.NewItemSearchClientImpl(apiid, affid)
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	keyword := "つぼみ"

	res, err := itemSearchClient.Search(keyword)
	if err != nil {
		log.Print(err)
	}

	items := res.Result.Items
	for _, item := range items {
		bytes, err := json.Marshal(item)
		if err != nil {
			log.Print(err)
		}

		err = redisClient.Publish("item_feed", bytes).Err()
		if err != nil {
			panic(err)
		}

		// filename := "/mnt/temp/dmm/" + fmt.Sprintf("%04d_%s.json", i, keyword)
		// log.Printf("save file: %s", filename)
		// save(filename, bytes)
	}
}

func save(filename string, bytes []byte) {
	err := ioutil.WriteFile(filename, bytes, os.ModePerm)
	if err != nil {
		log.Fatalf("failed to write file: %s", filename)
	}
}
