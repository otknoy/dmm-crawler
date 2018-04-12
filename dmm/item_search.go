package dmm

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type ItemSearchClient interface {
	Search(string, int, int) (ItemResponse, error)
}

type ItemSearchClientImpl struct {
	dmmAPIID       string
	dmmAffiliateID string
}

func NewItemSearchClientImpl(dmmAPIID string, dmmAffiliateID string) ItemSearchClient {
	c := &ItemSearchClientImpl{
		dmmAPIID,
		dmmAffiliateID,
	}

	return c
}

func (c *ItemSearchClientImpl) Search(keyword string, hits int, offset int) (ItemResponse, error) {
	u := c.buildURL(keyword, hits, offset)
	log.Println(u.String())

	res, err := http.Get(u.String())
	if err != nil {
		return ItemResponse{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ItemResponse{}, err
	}

	r := ItemResponse{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return ItemResponse{}, err
	}

	return r, nil
}

func (c *ItemSearchClientImpl) buildURL(keyword string, hits int, offset int) *url.URL {
	q := url.Values{}
	q.Add("api_id", c.dmmAPIID)
	q.Add("affiliate_id", c.dmmAffiliateID)
	q.Add("site", "DMM.R18")
	q.Add("service", "digital")
	q.Add("floor", "videoa")
	q.Add("hits", strconv.Itoa(hits))
	q.Add("offset", strconv.Itoa(offset))
	q.Add("keyword", keyword)
	q.Add("sort", "date")

	u := &url.URL{}
	u.Scheme = "https"
	u.Host = "api.dmm.com"
	u.Path = "affiliate/v3/ItemList"
	u.RawQuery = q.Encode()

	return u
}
