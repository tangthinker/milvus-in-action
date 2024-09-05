package util

import (
	"errors"
	"gocv.io/x/gocv"
	"image"
	"log"
)

const (
	UniformedImgWidth  = 50
	UniformedImgHeight = 50
)

func Image2Vector(filename string) ([]float32, error) {

	log.Println("Loading image", filename)
	img := gocv.IMRead(filename, gocv.IMReadColor)
	if img.Empty() {
		return nil, errors.New("failed to read image")
	}

	defer img.Close()

	resizedImg := gocv.NewMat()
	defer resizedImg.Close()
	gocv.Resize(img, &resizedImg, image.Point{X: UniformedImgWidth, Y: UniformedImgHeight}, 0, 0, gocv.InterpolationLinear)

	grayImg := gocv.NewMat()
	defer grayImg.Close()

	gocv.CvtColor(resizedImg, &grayImg, gocv.ColorBGRToGray)

	size := grayImg.Rows() * grayImg.Cols()

	vector := make([]float32, grayImg.Rows()*grayImg.Cols())

	for i := 0; i < grayImg.Rows(); i++ {
		for j := 0; j < grayImg.Cols(); j++ {
			vector[i*grayImg.Cols()+j] = float32(grayImg.GetUCharAt(j, i))
		}
	}

	for i := 0; i < size; i++ {
		if vector[i] != 0 {
			vector[i] /= 255
		}
	}

	return vector, nil
}

//
//func Image2Vector(img image.Image) []float64 {
//	bounds := img.Bounds()
//	width, height := bounds.Max.X, bounds.Max.Y
//
//	vec := make([]float64, width*height)
//
//	cur := 0
//
//	for y := 0; y < height; y++ {
//		for x := 0; x < width; x++ {
//			r, g, b, _ := img.At(x, y).RGBA()
//			vec[cur] = float64(r >> 8)
//			cur++
//			vec[cur] = float64(g >> 8)
//			cur++
//			vec[cur] = float64(b >> 8)
//		}
//	}
//
//	return vec
//}
//
//func LoadImage(filename string) (image.Image, error) {
//
//	file, err := os.Open(filename)
//
//	if err != nil {
//		return nil, err
//	}
//
//	defer file.Close()
//
//	img, _, err := image.Decode(file)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return img, nil
//
//}
