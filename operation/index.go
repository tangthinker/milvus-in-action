package operation

import (
	"context"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
	"github.com/tangthinker/milvus-in-action/milvus"
	"log"
)

func CreateIndex() {

	err := milvus.GlobalMilvusClient.CreateIndex(context.Background(), collectionName, "img_vector", &entity.IndexHNSW{}, false)

	if err != nil {
		log.Println("create index error", err)
		panic(err)
	}

}
