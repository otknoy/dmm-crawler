package repository

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/otknoy/dmm-crawler/domain/model"
)

type ItemRepository interface {
	Search(keyword string, hits int, offset int) (model.ItemResponse, error)
}

type DmmItemRepository struct {
	dmmAPIID       string
	dmmAffiliateID string
}

func NewDmmItemRepository(dmmAPIID string, dmmAffiliateID string) ItemRepository {
	r := &DmmItemRepository{
		dmmAPIID,
		dmmAffiliateID,
	}

	return r
}

func (r *DmmItemRepository) Search(keyword string, hits int, offset int) (model.ItemResponse, error) {
	u := r.buildURL(keyword, hits, offset)
	log.Println(u.String())

	res, err := http.Get(u.String())
	if err != nil {
		return model.ItemResponse{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return model.ItemResponse{}, err
	}

	response := model.ItemResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return model.ItemResponse{}, err
	}

	return response, nil
}

func (r *DmmItemRepository) buildURL(keyword string, hits int, offset int) *url.URL {
	q := url.Values{}
	q.Add("api_id", r.dmmAPIID)
	q.Add("affiliate_id", r.dmmAffiliateID)
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
