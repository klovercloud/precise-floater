package main

import (
	"errors"
	"github.com/klovercloud/precise-floater/enum"
	"math"
)

type PreciseFloater struct {
}

func (preciseFloater PreciseFloater) Float32(input float32, mode enum.PreciseMode, precisionWidth int) (float32, error) {
	if mode != enum.Ceil && mode != enum.Floor && mode != enum.Round {
		return 0.0, errors.New("Precision Mode invalid ")
	}
	if precisionWidth < 0 || precisionWidth > 5 {
		return 0.0, errors.New("precisionWidth invalid")
	}
	for i := 0; i < precisionWidth; i++ {
		input = input * 10
	}
	if mode == enum.Ceil {
		input = float32(math.Ceil(float64(input)))
		input = input / float32(math.Pow10(precisionWidth))
	}
	if mode == enum.Round {
		input = float32(math.Round(float64(input)))
		input = input / float32(math.Pow10(precisionWidth))
	}
	if mode == enum.Floor {
		input = float32(math.Floor(float64(input)))
		input = input / float32(math.Pow10(precisionWidth))
	}
	return input, nil
}
func (preciseFloater PreciseFloater) Float64(input float64, mode enum.PreciseMode, precisionWidth int) (float64, error) {
	if mode != enum.Ceil && mode != enum.Floor && mode != enum.Round {
		return 0.0, errors.New("Precision Mode invalid ")
	}
	if precisionWidth < 0 || precisionWidth > 5 {
		return 0.0, errors.New("precisionWidth invalid")
	}
	for i := 0; i < precisionWidth; i++ {
		input = input * 10
	}
	if mode == enum.Ceil {
		input = math.Ceil(input) / math.Pow10(precisionWidth)
	}
	if mode == enum.Round {
		input = math.Round(input) / math.Pow10(precisionWidth)
	}
	if mode == enum.Floor {
		input = math.Floor(input) / math.Pow10(precisionWidth)
	}
	return input, nil
}
func (preciseFloater PreciseFloater) DefaultFloat64(input float64) float64 {
	rtn, _ := PreciseFloater{}.Float64(input, enum.Ceil, 2)
	return rtn
}
func (preciseFloater PreciseFloater) DefaultFloat32(input float32) float32 {
	rtn, _ := PreciseFloater{}.Float32(input, enum.Ceil, 2)
	return rtn
}
