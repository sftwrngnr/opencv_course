package w1lib

import (
	"fmt"

	"gocv.io/x/gocv"
)

func ManipulatingPixels() {
	imagePath := DATA_PATH + "/images/number_zero.jpg"

	// Read image in Grayscale format
	testImage := gocv.IMRead(imagePath, 0)
	PrintImageChannelValues(&testImage) //In Utils
	bgr := gocv.Split(testImage)
	fmt.Printf("Currently %d\n", bgr[0].GetUCharAt(0, 0))
	for x := 0; x < 7; x++ {
		for y := 0; y < 5; y++ {
			bgr[0].SetUCharAt(y, x, 128)
		}
	}
	gocv.Merge(bgr, &testImage)
	PrintImageChannelValues(&testImage) //In Utils
	window := gocv.NewWindow("Image as a Matrix")
	window.IMShow(testImage)

	for {
		c := gocv.WaitKey(20)
		if c == 27 {
			break
		}
	}
	window.Close()

}
