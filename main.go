package main

import (
	"github.com/otknoy/dmm-crawler/application"
)

func main() {
	crawler := application.NewCrawler()

	crawler.Crawl()
}
