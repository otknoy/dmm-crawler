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

	responses := make(chan dmm.ItemResponse, 10)
	go func(keywords []string) {
		for _, keyword := range actresses {
			res, _ := itemSearchClient.Search(keyword)
			responses <- res
		}
		close(responses)
	}(actresses)

	for itemResponse := range responses {
		keyword := itemResponse.Request.Parameters.Keyword
		log.Println(keyword)
		for i, item := range itemResponse.Result.Items {
			bytes, _ := json.Marshal(item)

			outputDir := "/mnt/temp/dmm/"
			filename := outputDir + fmt.Sprintf("%04d_%s.json", i, keyword)
			// log.Printf("success to save file: %s", filename)
			err := ioutil.WriteFile(filename, bytes, os.ModePerm)
			if err != nil {
				log.Fatalf("failed to write file: %s", filename)
			}
		}
	}
}
