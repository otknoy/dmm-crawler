package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/otknoy/dmm-crawler/dmm"
)

func main() {
	actresses := []string{}
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		actress := stdin.Text()
		actresses = append(actresses, actress)
	}

	apiid := os.Getenv("DMM_API_ID")
	affid := os.Getenv("DMM_AFFILIATE_ID")
	itemSearchClient := dmm.NewItemSearchClientImpl(apiid, affid)

	responses := make(chan dmm.ItemResponse, 4)
	go func(keywords []string) {
		for _, keyword := range keywords {
			for i := 0; i < 50000; i++ {
				res, _ := itemSearchClient.Search(keyword, 100, i*100+1)
				if res.Result.ResultCount == 0 {
					break
				}

				responses <- res
			}
		}
		close(responses)
	}(actresses)

	items := make(chan dmm.Item)
	go func(response <-chan dmm.ItemResponse) {
		for itemResponse := range responses {
			keyword := itemResponse.Request.Parameters.Keyword
			log.Println(keyword)
			for _, item := range itemResponse.Result.Items {
				items <- item
			}
		}
		close(items)
	}(responses)

	for item := range items {
		bytes, _ := json.Marshal(item)

		outputDir := "/mnt/temp/dmm/"
		filename := outputDir + fmt.Sprintf("%s.json", item.ContentID)
		// log.Printf("success to save file: %s", filename)
		err := ioutil.WriteFile(filename, bytes, os.ModePerm)
		if err != nil {
			log.Fatalf("failed to write file: %s", filename)
		}
	}
}
