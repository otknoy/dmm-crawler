package model

type CrawlRequest struct {
	keyword string
	n       uint
	page    uint
	sort    string
}

func NewCrawlRequest(keyword string, n uint, page uint, sort string) CrawlRequest {
	return CrawlRequest{keyword, n, page, sort}
}

func (r *CrawlRequest) GetKeyword() string {
	return r.keyword
}

func (r *CrawlRequest) GetN() uint {
	return r.n
}

func (r *CrawlRequest) GetOffset() uint {
	return r.n*(r.page-1) + 1
}

func (r *CrawlRequest) GetSort() string {
	return r.sort
}
