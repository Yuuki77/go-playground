package algorithm

import "errors"

func binarySearchError(array []int8, lo int8, hi int8, target int8) (int8, error) {
	for lo <= hi {
		middle := (lo + hi) / 2
		if middle < 0 {
			return -1, errors.New("integer overflow")
		}
		if array[middle] == target {
			return middle, nil
		} else if target < array[middle] {
			hi = middle - 1
		} else {
			lo = middle + 1
		}
	}
	return -1, nil
}

func binarySearch(array []int8, lo int8, hi int8, target int8) (int8, error) {
	for lo <= hi {
		middle := lo + (hi-lo)/2
		if middle < 0 {
			return -1, errors.New("integer overflow")
		}
		if array[middle] == target {
			return middle, nil
		} else if target < array[middle] {
			hi = middle - 1
		} else {
			lo = middle + 1
		}
	}
	return -1, nil
}
