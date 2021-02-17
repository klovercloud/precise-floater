package PreciseFloater

import (
	"github.com/klovercloud/precise-floater/enum"
	"testing"
)

type float32TestStruct struct {
	input    float32
	mode     enum.PreciseMode
	width    int
	expected float32
}
type float64TestStruct struct {
	input    float64
	mode     enum.PreciseMode
	width    int
	expected float64
}

var float32TestArray = []float32TestStruct{
	//Floor Test Cases
	{12.664, enum.Floor, 0, 12},
	{12.0125, enum.Floor, 0, 12},
	{12.664, enum.Floor, 0, 12},
	{12.0101, enum.Floor, 0, 12},
	{12.6134, enum.Floor, 1, 12.6},
	{12.6632, enum.Floor, 1, 12.6},
	{78.599, enum.Floor, 1, 78.5},
	{78.99, enum.Floor, 1, 78.9},
	{12.45743, enum.Floor, 2, 12.45},
	{12.9999, enum.Floor, 2, 12.99},
	{12.1111, enum.Floor, 2, 12.11},
	{12.2364, enum.Floor, 2, 12.23},
	{12.05, enum.Floor, 2, 12.05},
	{12.699, enum.Floor, 2, 12.69},

	//Round Test Cases
	{12.664, enum.Round, 0, 13},
	{12.0125, enum.Round, 0, 12},
	{12.664, enum.Round, 0, 13},
	{12.0101, enum.Round, 0, 12},
	{12.6134, enum.Round, 1, 12.6},
	{12.6632, enum.Round, 1, 12.7},
	{78.599, enum.Round, 1, 78.6},
	{78.99, enum.Round, 1, 79.0},
	{12.45743, enum.Round, 2, 12.46},
	{12.9999, enum.Round, 2, 13.00},
	{12.1111, enum.Round, 2, 12.11},
	{12.2364, enum.Round, 2, 12.24},
	{12.05, enum.Round, 2, 12.05},
	{12.699, enum.Round, 2, 12.70},

	//Ceil Test Cases
	//Recommended to use in billing with WidthPrecision : 2
	{12.664, enum.Ceil, 0, 13},
	{12.0125, enum.Ceil, 0, 13},
	{12.664, enum.Ceil, 0, 13},
	{12.0101, enum.Ceil, 0, 13},
	{12.6134, enum.Ceil, 1, 12.7},
	{12.6632, enum.Ceil, 1, 12.7},
	{78.599, enum.Ceil, 1, 78.6},
	{78.99, enum.Ceil, 1, 79.0},
	{12.45743, enum.Ceil, 2, 12.46},
	{12.9999, enum.Ceil, 2, 13.00},
	{12.1111, enum.Ceil, 2, 12.12},
	{12.2364, enum.Ceil, 2, 12.24},
	{12.05, enum.Ceil, 2, 12.05},
	{12.699, enum.Ceil, 2, 12.70},
}
var float64TestArray = []float64TestStruct{
	//Floor Test Cases
	{12.664, enum.Floor, 0, 12},
	{12.0125, enum.Floor, 0, 12},
	{12.664, enum.Floor, 0, 12},
	{12.0101, enum.Floor, 0, 12},
	{12.6134, enum.Floor, 1, 12.6},
	{12.6632, enum.Floor, 1, 12.6},
	{78.599, enum.Floor, 1, 78.5},
	{78.99, enum.Floor, 1, 78.9},
	{12.45743, enum.Floor, 2, 12.45},
	{12.9999, enum.Floor, 2, 12.99},
	{12.1111, enum.Floor, 2, 12.11},
	{12.2364, enum.Floor, 2, 12.23},
	{12.05, enum.Floor, 2, 12.05},
	{12.699, enum.Floor, 2, 12.69},

	//Round Test Cases
	{12.664, enum.Round, 0, 13},
	{12.0125, enum.Round, 0, 12},
	{12.664, enum.Round, 0, 13},
	{12.0101, enum.Round, 0, 12},
	{12.6134, enum.Round, 1, 12.6},
	{12.6632, enum.Round, 1, 12.7},
	{78.599, enum.Round, 1, 78.6},
	{78.99, enum.Round, 1, 79.0},
	{12.45743, enum.Round, 2, 12.46},
	{12.9999, enum.Round, 2, 13.00},
	{12.1111, enum.Round, 2, 12.11},
	{12.2364, enum.Round, 2, 12.24},
	{12.05, enum.Round, 2, 12.05},
	{12.699, enum.Round, 2, 12.70},

	//Ceil Test Cases
	//Recommended to use in billing with WidthPrecision : 2
	{12.664, enum.Ceil, 0, 13},
	{12.0125, enum.Ceil, 0, 13},
	{12.664, enum.Ceil, 0, 13},
	{12.0101, enum.Ceil, 0, 13},
	{12.6134, enum.Ceil, 1, 12.7},
	{12.6632, enum.Ceil, 1, 12.7},
	{78.599, enum.Ceil, 1, 78.6},
	{78.99, enum.Ceil, 1, 79.0},
	{12.45743, enum.Ceil, 2, 12.46},
	{12.9999, enum.Ceil, 2, 13.00},
	{12.1111, enum.Ceil, 2, 12.12},
	{12.2364, enum.Ceil, 2, 12.24},
	{12.05, enum.Ceil, 2, 12.05},
	{12.699, enum.Ceil, 2, 12.70},
}

func TestFloat32(t *testing.T) {

	for _, test := range float32TestArray {
		preciseFloater := PreciseFloater{}
		output, err := preciseFloater.Float32(test.input, test.mode, test.width)
		if err != nil {
			t.Error("Error Thrown: " + err.Error())
		} else if output != test.expected {
			t.Errorf("Output %f not equal to expected %f. Mode: %s Width: %d", output, test.expected, test.mode, test.width)
		}

	}
}
func TestFloat64(t *testing.T) {

	for _, test := range float64TestArray {
		preciseFloater := PreciseFloater{}
		output, err := preciseFloater.Float64(test.input, test.mode, test.width)
		if err != nil {
			t.Error("Error Thrown: " + err.Error())
		} else if output != test.expected {
			t.Errorf("Output %f not equal to expected %f. Mode: %s Width: %d", output, test.expected, test.mode, test.width)
		}

	}
}
