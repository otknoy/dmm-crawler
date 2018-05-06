package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/otknoy/dmm-crawler/domain/interfaces"
	"github.com/otknoy/dmm-crawler/domain/model"
)

type ItemSaveService struct {
	outputDir string
}

func NewItemSaveService(outputDir string) (interfaces.ItemSaveService, error) {
	return &ItemSaveService{outputDir}, nil
}

func (iss *ItemSaveService) SaveItem(filename string, item model.DmmItem) error {
	filepath := filepath.Join(iss.outputDir, filename)

	bytes, _ := json.Marshal(item)
	err := ioutil.WriteFile(filepath, bytes, os.ModePerm)
	if err != nil {
		log.Fatalf("failed to write file: %s", filepath)
		return err
	}
	log.Printf("success to save file: %s", filepath)

	return nil
}
