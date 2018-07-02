package model

import "testing"

func TestGetOffset(t *testing.T) {
	keyword := "test-keyword"
	sort := "test-sort"

	var r CrawlRequest

	r = NewCrawlRequest(keyword, 0, 0, sort)
	if r.GetOffset() != 1 {
		t.Errorf("fail: {hits=0, page=0}, offset to be 1, but %d\n", r.GetOffset())
	}

	r = NewCrawlRequest(keyword, 1, 1, sort)
	if r.GetOffset() != 1 {
		t.Errorf("fail: {hits=1, page=1}, offset to be 1, but %d\n", r.GetOffset())
	}

	r = NewCrawlRequest(keyword, 5, 1, sort)
	if r.GetOffset() != 1 {
		t.Errorf("fail: {hits=5, page=1}, offset to be 1, but %d\n", r.GetOffset())
	}

	r = NewCrawlRequest(keyword, 10, 2, sort)
	if r.GetOffset() != 11 {
		t.Errorf("fail: {hits=10, page=2}, offset to be 11, but %d\n", r.GetOffset())
	}

	r = NewCrawlRequest(keyword, 10, 13, sort)
	if r.GetOffset() != 121 {
		t.Errorf("fail: {hits=10, page=13}, offset to be 121, but %d\n", r.GetOffset())
	}
}
