package main

import (
	"testing"
)

func TestEvenArray(t *testing.T) {
	arr1, arr2 := [2]int{1, 3}, [2]int{2, 4}
	median := findMedianSortedArrays(arr1[:], arr2[:])
	
	m := float64(arr1[1] + arr2[0]) / 2
	if median != m {
		t.Errorf("A mediana esta incorreta, o resultado era para ser: %f, mas o recebido foi: %f", m, median)
	}
	
}

func TestOddArray(t *testing.T) {
	arr1, arr2 := [2]int{1, 3}, [1]int{2}
	median := findMedianSortedArrays(arr1[:], arr2[:])

	if median != float64(arr2[0]) {
		t.Errorf("mediana deveria ser: %f, mas o valor devolvido foi: %f", float64(arr2[0]), median)
	}
}