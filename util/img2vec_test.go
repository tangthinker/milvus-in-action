package util

import (
	"testing"
)

func TestImage2Vector(t *testing.T) {

	imgPath := "test.png"
	vec, err := Image2Vector(imgPath)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(vec)

}
