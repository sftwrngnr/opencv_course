package w1lib

import (
	"fmt"

	"gocv.io/x/gocv"
)

func ImageAsAMatrix() {
	imagePath := DATA_PATH + "/images/number_zero.jpg"

	// Read image in Grayscale format
	testImage := gocv.IMRead(imagePath, 0)
	PrintImageChannelValues(&testImage) //In Utils

	//Image properties
	fmt.Printf("Object type = %s\n", testImage.Type())
	fmt.Printf("Image Dimensions = %d x %d\n", testImage.Rows(), testImage.Cols())

	window := gocv.NewWindow("Image as a Matrix")
	defer window.Close()
	window.IMShow(testImage)

	for {
		c := gocv.WaitKey(20)
		if c == 27 {
			break
		}
	}
}
