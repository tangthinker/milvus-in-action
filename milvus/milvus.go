package milvus

import (
	"context"
	"github.com/milvus-io/milvus-sdk-go/v2/client"
)

var GlobalMilvusClient client.Client

func init() {
	cli, err := client.NewClient(context.Background(), client.Config{
		Address: "localhost:19530",
	})

	if err != nil {
		panic(err)
	}

	GlobalMilvusClient = cli
}
