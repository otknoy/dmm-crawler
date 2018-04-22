package infrastructure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/otknoy/dmm-crawler/interfaces"
	"github.com/otknoy/dmm-crawler/model"
)

type SolrRepository struct{}

func NewSolrRepository() interfaces.SolrAdder {
	return &SolrRepository{}
}

func (r *SolrRepository) Add(item model.Item) error {
	doc := model.NewSolrDocument(item)

	jsonStr, err := json.Marshal(doc)
	if err != nil {
		return err
	}

	url := "http://localhost:8983/solr/dmm_items/update?commit=true"

	res, err := http.Post(url, "text/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	return nil
	// }
	// log.Println(string(body))

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("err, http-status-code: %d", res.StatusCode)
	}

	return nil
}
