package main

import (
	"milvus-in-action/operation"
)

func main() {
	operation.CreateCollection()
	//operation.CreateIndex()
	//
	//searchVector, err := util.Image2Vector(operation.SearchImageName)
	//if err != nil {
	//	panic(err)
	//}
	//
	//operation.Search(searchVector, 4)
}
