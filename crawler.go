package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/otknoy/dmm-crawler/dmm"
)

func main() {
	apiid := os.Getenv("DMM_API_ID")
	affid := os.Getenv("DMM_AFFILIATE_ID")

	client := dmm.NewItemSearchClientImpl(apiid, affid)

	keyword := "つぼみ"

	res, err := client.Search(keyword)
	if err != nil {
		log.Print(err)
	}

	items := res.Result.Items
	for i, item := range items {
		bytes, err := json.Marshal(item)
		if err != nil {
			log.Print(err)
		}

		filename := "/mnt/temp/dmm/" + fmt.Sprintf("%04d_%s.json", i, keyword)
		log.Printf("save file: %s", filename)
		save(filename, bytes)
	}
}

func save(filename string, bytes []byte) {
	err := ioutil.WriteFile(filename, bytes, os.ModePerm)
	if err != nil {
		log.Fatalf("failed to write file: %s", filename)
	}
}
