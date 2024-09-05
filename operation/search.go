package operation

import (
	"context"
	"fmt"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
	"github.com/tangthinker/milvus-in-action/milvus"
	"log"
)

func Search(vec []float32, topK int) {

	sp, err := entity.NewIndexHNSWSearchParam(10)
	if err != nil {
		panic(err)
	}

	searchResult, err := milvus.GlobalMilvusClient.Search(
		context.Background(),
		collectionName,
		nil,
		"",
		[]string{"id", "img_vector", "img_info"},
		[]entity.Vector{entity.FloatVector(vec)},
		"img_vector",
		entity.COSINE,
		topK,
		sp,
	)

	if err != nil {
		panic(err)
	}

	log.Println(searchResult)

	result := searchResult[0]

	for i := 0; i < result.ResultCount; i++ {
		fmt.Print("[")
		id, err := result.IDs.Get(i)
		if err != nil {
			panic(err)
		}
		fmt.Print("id: ", id)

		//imgVector, err := result.Fields.GetColumn("img_vector").Get(i)
		//if err != nil {
		//	panic(err)
		//}
		//log.Print("imgVector: ", imgVector)

		imgInfo, err := result.Fields.GetColumn("img_info").Get(i)
		if err != nil {
			panic(err)
		}
		fmt.Print(" imgInfo: ", string(imgInfo.([]byte)))

		score := result.Scores[i]
		fmt.Print(" score: ", score)
		fmt.Println("]")

	}

}
