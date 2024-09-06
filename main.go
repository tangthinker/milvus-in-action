package main

import (
	"github.com/tangthinker/milvus-in-action/operation"
	"github.com/tangthinker/milvus-in-action/util"
)

func main() {
	//operation.CreateCollection()
	searchVector, err := util.Image2Vector(operation.SearchImageName)
	if err != nil {
		panic(err)
	}

	operation.Search(searchVector, 4)
}
