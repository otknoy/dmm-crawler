package model

import (
	"fmt"
	"net/url"
)

type SearchRequest struct {
	keyword string
	hits    uint
	offset  uint
	sort    string
}

func NewSearchRequest(keyword string, hits uint, offset uint, sort string) SearchRequest {
	return SearchRequest{keyword, hits, offset, sort}
}

func (r *SearchRequest) ToUrlValues() url.Values {
	q := url.Values{}
	q.Add("site", "DMM.R18")
	q.Add("service", "digital")
	q.Add("floor", "videoa")
	q.Add("hits", fmt.Sprint(r.hits))
	q.Add("offset", fmt.Sprint(r.offset))
	q.Add("keyword", r.keyword)
	q.Add("sort", r.sort)

	return q
}
