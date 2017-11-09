package _7_recursion

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arrayA := []int{23, 47, 81, 95}
	arrayB := []int{7, 14, 39, 55, 62, 74}
	arrayC := mergeSort(arrayA, arrayB)
	expected := []int{7, 14, 23, 39, 47, 55, 62, 74, 81, 95}
	if !reflect.DeepEqual(arrayC, expected) {
		t.Error("Want", expected, "have", arrayC)
	}
}