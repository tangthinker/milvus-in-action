package operation

import (
	"context"
	"github.com/tangthinker/milvus-in-action/milvus"
	"log"
	"os"
)

func LoadCollection() {

	if err := milvus.GlobalMilvusClient.LoadCollection(context.Background(), collectionName, false); err != nil {
		log.Println("load collection failed:", err)
		os.Exit(1)
	}

	log.Println("load collection successfully")

}
