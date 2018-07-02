package model

import (
	"net/url"
	"reflect"
	"testing"
)

func TestToUlrlValues(t *testing.T) {
	r := NewSearchRequest("test-keyword", 10, 0, "test-sort")

	v := r.ToUrlValues()

	q := url.Values{}
	q.Add("keyword", "test-keyword")
	q.Add("hits", "10")
	q.Add("offset", "0")
	q.Add("sort", "test-sort")
	q.Add("site", "DMM.R18")
	q.Add("service", "digital")
	q.Add("floor", "videoa")

	eq := reflect.DeepEqual(q, v)

	if !eq {
		t.Error(v)
	}
}
