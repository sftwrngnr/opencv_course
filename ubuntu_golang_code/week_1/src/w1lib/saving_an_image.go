package w1lib

import (
	"gocv.io/x/gocv"
)

func SavingAnImage() {
	imagePath := DATA_PATH + "/images/number_zero.jpg"

	// Read image in Grayscale format
	testImage := gocv.IMRead(imagePath, 0)
	gocv.IMWrite("results/test.jpg", testImage)
}
