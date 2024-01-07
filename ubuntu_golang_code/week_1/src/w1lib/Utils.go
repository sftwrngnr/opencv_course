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
