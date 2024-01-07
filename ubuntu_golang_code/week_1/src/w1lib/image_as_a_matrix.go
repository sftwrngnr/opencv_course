package w1lib

import (
	"fmt"

	"gocv.io/x/gocv"
)

func PrintImageChannelValues(img *gocv.Mat) {
	bgr := gocv.Split(*img)
	nr := img.Rows()
	nc := img.Cols()
	nic := img.Channels()
	fmt.Printf("[")
	for i := 0; i < nr; i++ {
		fmt.Printf("[")
		for j := 0; j < nc; j++ {
			if nic > 1 {
				fmt.Printf("(")
			}
			for k := 0; k < nic; k++ {
				fmt.Printf("%3d", bgr[k].GetUCharAt(i, j))
				if k < (nic - 1) {
					fmt.Printf(",")
				}
			}
			if nic > 1 {
				fmt.Printf(")")
			}
			fmt.Printf(" ")
		}
		if i < nr-1 {
			fmt.Print("]\n")
		} else {
			fmt.Printf("]")
		}
	}
	fmt.Printf("]\n")
	fmt.Printf("Data type = %T\n", bgr)

}

func ImageAsAMatrix() {
	imagePath := DATA_PATH + "/images/number_zero.jpg"

	// Read image in Grayscale format
	testImage := gocv.IMRead(imagePath, 0)
	PrintImageChannelValues(&testImage)

	//Image properties
	fmt.Printf("Object type = %s\n", testImage.Type())
	fmt.Printf("Image Dimensions = %d x %d\n", testImage.Rows(), testImage.Cols())

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
