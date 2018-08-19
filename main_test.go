package TwoNumberSum

import (
	"reflect"
	"testing"
)

func TestCase(t *testing.T) {
	expected := []int{-1, 11}

	x := []int{3, 5, -4, 8, 11, 1, -1, 6}
	y := 10

	output := twoNumberSum(x, y)

	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Expected: %v, Output: %v", expected, output)
	}
}
