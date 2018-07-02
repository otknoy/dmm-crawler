package infrastructure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/otknoy/dmm-crawler/domain/model"
	"github.com/otknoy/dmm-crawler/domain/repository"
)

type DmmItemRepository struct {
	dmmAPIID       string
	dmmAffiliateID string
}

func NewDmmItemRepository(dmmAPIID string, dmmAffiliateID string) repository.ItemSearcher {
	r := &DmmItemRepository{
		dmmAPIID,
		dmmAffiliateID,
	}

	return r
}

func (r *DmmItemRepository) Search(searchRequest model.SearchRequest) (model.ItemResponse, error) {
	q := searchRequest.ToUrlValues()
	q.Add("api_id", r.dmmAPIID)
	q.Add("affiliate_id", r.dmmAffiliateID)

	u := &url.URL{}
	u.Scheme = "https"
	u.Host = "api.dmm.com"
	u.Path = "affiliate/v3/ItemList"
	u.RawQuery = q.Encode()

	log.Println(u)

	res, err := http.Get(u.String())
	if err != nil {
		return model.ItemResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return model.ItemResponse{}, fmt.Errorf("err, http-status-code: %d", res.StatusCode)
	}

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
