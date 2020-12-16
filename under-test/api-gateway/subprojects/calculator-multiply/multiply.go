package main

import (
	"errors"
	"strconv"
)

func Multiply(x string, y string) (float64, error) {

	xInt, err1 := strconv.ParseFloat(x, 64)
	yInt, err2 := strconv.ParseFloat(y, 64)

	if err1 == nil && err2 == nil {
		return xInt*yInt, nil
	} else {
		return 0, errors.New("Invalid Parameters")
	}
}
