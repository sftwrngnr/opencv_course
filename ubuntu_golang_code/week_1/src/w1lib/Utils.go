package w1lib

import (
	"errors"
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

func GetRoi(img *gocv.Mat, inch, sx, ex, sy, ey int) ([][]uint8, error) {
	bgr := gocv.Split(*img)
	nic := img.Channels()
	nr := img.Rows()
	nc := img.Cols()

	if (sx < 0) || (sx > nc) {
		return nil, errors.New("Start column value is invalid")
	}
	if (ex < sx) || (ex > nc) {
		return nil, errors.New("End column value is invalid")
	}
	if (sy < 0) || (sy > nr) {
		return nil, errors.New("Start row value is invalid")
	}
	if (ey < sy) || (ey > nr) {
		return nil, errors.New("End row value is invalid")
	}
	if inch < 0 || inch > nic {
		return nil, errors.New("Channel is invalid")
	}
	// Allocate return channels
	rval := make([][]uint8, (ey - sy))
	cury := 0
	for y := sy; y < ey; y++ {
		rval[cury] = make([]uint8, (ex - sx))
		curx := 0
		for x := sx; x < ex; x++ {
			rval[cury][curx] = bgr[inch].GetUCharAt(y, x)
			curx++
		}
		cury++
	}
	return rval, nil
}

func PrintRoiValues(inRoi [][]uint8) {
	nc := len(inRoi)
	for i := 0; i < nc; i++ {
		fmt.Printf("[")
		for j := 0; j < len(inRoi[i]); j++ {
			fmt.Printf("%3d ", inRoi[i][j])
		}
		fmt.Printf("]\n")
	}
}

func SetRoi(img *gocv.Mat, inch, sx, ex, sy, ey int, nValue uint8) error {
	bgr := gocv.Split(*img)
	nic := img.Channels()
	nr := img.Rows()
	nc := img.Cols()

	if (sx < 0) || (sx > nc) {
		return errors.New("Start column value is invalid")
	}
	if (ex < sx) || (ex > nc) {
		return errors.New("End column value is invalid")
	}
	if (sy < 0) || (sy > nr) {
		return errors.New("Start row value is invalid")
	}
	if (ey < sy) || (ey > nr) {
		return errors.New("End row value is invalid")
	}
	if inch < 0 || inch > nic {
		return errors.New("Channel is invalid")
	}
	// Allocate return channels
	cury := 0
	for y := sy; y < ey; y++ {
		curx := 0
		for x := sx; x < ex; x++ {
			bgr[inch].SetUCharAt(y, x, nValue)
			curx++
		}
		cury++
	}
	gocv.Merge(bgr, img)
	return nil
}
