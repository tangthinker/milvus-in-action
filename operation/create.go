package operation

import (
	"context"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
	"log"
	"milvus-in-action/milvus"
	"os"
)

func CreateCollection() {

	log.Println("Creating collection " + collectionName)

	schema := entity.NewSchema().WithName(collectionName).
		WithField(entity.NewField().WithName("id").WithDataType(entity.FieldTypeInt64).WithIsPrimaryKey(true).WithIsAutoID(true)).
		WithField(entity.NewField().WithName("img_vector").WithDataType(entity.FieldTypeFloatVector).WithDim(2500)).
		WithField(entity.NewField().WithName("img_info").WithDataType(entity.FieldTypeJSON))

	if err := milvus.GlobalMilvusClient.CreateCollection(context.Background(), schema, entity.DefaultShardNumber); err != nil {
		log.Fatal("create collection fail ", err)
		os.Exit(1)
	}
	log.Println("Create collection" + collectionName + " successfully")
}
