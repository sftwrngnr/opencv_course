package main

import (
	"gocv.io/x/gocv"
)

func main() {
	image := gocv.IMRead("boy.jpg", 0)
	gocv.IMWrite("boyGray.jpg", image)
}
