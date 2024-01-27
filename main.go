package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    merged := mergeArray(nums1, nums2)
    Quicksort(merged, 0, len(merged) - 1)
    if len(merged) % 2 == 0 {
       return float64((merged[(len(merged) / 2 - 1)] + merged[(len(merged) / 2)])) / 2
    }
    
    return float64(merged[len(merged) / 2])
}

 func mergeArray(list1 []int, list2[]int) []int {
    merged := append(list1, list2...)
    return merged
}

func Quicksort(a []int, low, high int) {
	if(low < high) {
		pivotLocation := Partition(a, low, high)
		Quicksort(a, low, pivotLocation - 1)
		Quicksort(a, pivotLocation + 1, high) 
	}
}

func Partition(a []int, low, high int) int {
	pivot := a[high]
	leftwall := low - 1

	for i := low; i <= high - 1; i++ {
		if(a[i] < pivot) {
			leftwall += 1
			Swap(&a[i], &a[leftwall])
		}
	}

	Swap(&a[high], &a[leftwall + 1])

	return leftwall + 1
}

func Swap(r, l *int) {
	*r, *l = *l, *r
}