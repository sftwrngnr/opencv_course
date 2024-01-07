package w1lib

import (
	"fmt"
	"log"

	"gocv.io/x/gocv"
)

func ManipulatingGroupOfPixels() {
	imagePath := DATA_PATH + "/images/number_zero.jpg"

	// Read image in Grayscale format
	testImage := gocv.IMRead(imagePath, 0)
	PrintImageChannelValues(&testImage) //In Utils
	myRoi, err := GetRoi(&testImage, 0, 3, 5, 6, 8)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ROI is:\n")
	PrintRoiValues(myRoi)
	SetRoi(&testImage, 0, 3, 5, 6, 8, 111)

	PrintImageChannelValues(&testImage)
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
