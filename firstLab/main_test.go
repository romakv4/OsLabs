package main

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestReturnInt(t *testing.T) {
	if ReturnInt() != 1 {
		t.Error("expected 1")
	}
}

func ReturnInt() int {
	return 1
}
func TestReturnFloat(t *testing.T) {
	if ReturnFloat() != float32(1.1) {
		t.Error("expected 1.1")
	}
}

func ReturnFloat() float32 {
	return 1.1
}

func TestReturnIntArray(t *testing.T) {
	if ReturnIntArray() != [3]int{1, 3, 4} {
		t.Error("expected '[3]int{1, 3, 4}'")
	}
}

func ReturnIntArray() [3]int {
	return [3]int{1, 3, 4}
}

func TestReturnIntSlice(t *testing.T) {
	expected := []int{1, 2, 3}
	result := ReturnIntSlice()
	if !reflect.DeepEqual(result, expected) {
		t.Error("expected", expected, "have", result)
	}
}

func ReturnIntSlice() []int {
	return []int{1, 2, 3}
}

func TestIntSliceToString(t *testing.T) {
	expected := "1723100500"
	result := IntSliceToString([]int{17, 23, 100500})
	if expected != result {
		t.Error("expected", expected, "have", result)
	}
}

func IntSliceToString(slice []int) string {
	stringSlice := []string{}
	for i := range slice {
		value := slice[i]
		stringValue := strconv.Itoa(value)
		stringSlice = append(stringSlice, stringValue)
	}
	return strings.Join(stringSlice, "")
}

func TestMergeSlices(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5}
	result := MergeSlices([]float32{1.1, 2.1, 3.1}, []int32{4, 5})
	if !reflect.DeepEqual(result, expected) {
		t.Error("expected", expected, "have", result)
	}
}

func MergeSlices(floatSlice []float32, intSlice []int32) []int {
	var resultSlice []int
	for i := range floatSlice {
		resultSlice = append(resultSlice, int(floatSlice[i]))
	}
	for i := range intSlice {
		resultSlice = append(resultSlice, int(intSlice[i]))
	}
	return resultSlice
}
