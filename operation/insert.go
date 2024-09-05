package operation

import (
	"context"
	"encoding/json"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
	"io/fs"
	"log"
	"milvus-in-action/milvus"
	"milvus-in-action/util"
	"path/filepath"
)

func InsertData() {

	imgPaths := make([]string, 0)

	err := filepath.Walk(dataDirPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		imgPaths = append(imgPaths, path)
		return nil
	})

	if err != nil {
		panic(err)
	}

	for _, path := range imgPaths {

		log.Printf("insert data from %s", path)

		imgVector, err := util.Image2Vector(path)

		if err != nil {
			panic(err)
		}

		err = insertData(imgVector, path)
		if err != nil {
			panic(err)
		}
	}

	log.Printf("insert data success")

}

func insertData(vec []float32, path string) error {

	imgVector := entity.NewColumnFloatVector("img_vector", 2500, [][]float32{vec})

	imgInfo := &ImageInfo{
		Path: path,
	}

	imgInfoBytes, _ := json.Marshal(imgInfo)

	imgInfoCol := entity.NewColumnJSONBytes("img_info", [][]byte{imgInfoBytes})

	_, err := milvus.GlobalMilvusClient.Insert(
		context.Background(),
		collectionName,
		"",
		imgVector,
		imgInfoCol,
	)

	if err != nil {
		return err
	}

	return nil

}

type ImageInfo struct {
	Path string `json:"path"`
}
