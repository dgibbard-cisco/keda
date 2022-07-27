package util

import (
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	t.Errorf("%v != %v", a, b)
}

func TestFindStringInSlice(t *testing.T) {
	inputSlice := []string{"this", "looks", "for", "strings"}
	inputValue := "looks"
	expectedIndex, expectedBool := int(1), bool(true)

	outputIndex, outputBool := FindStringInSlice(inputSlice, inputValue)
	assertEqual(t, outputIndex, expectedIndex)
	assertEqual(t, outputBool, expectedBool)
}

func TestMaxFloatFromSlice(t *testing.T) {
	input := []float64{1.0, 2.0, 3.0, 4.0}
	expectedOutput := float64(4.0)

	output := MaxFloatFromSlice(input)

	assertEqual(t, output, expectedOutput)
}

func TestAvgFloatFromSlice(t *testing.T) {
	input := []float64{1.0, 2.0, 3.0, 4.0}
	expectedOutput := float64(2.5)

	output := AvgFloatFromSlice(input)

	assertEqual(t, output, expectedOutput)
}
