package common

// Contains returns true if the target is present in the slice, false otherwise.
func Contains[T comparable](slice []T, target T) bool {
	for idx := range slice {
		if slice[idx] == target {
			return true
		}
	}
	return false
}

// SameElements checks if the two slices contain the exact same elements. Order does not matter.
func SameElements[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}
	for _, v := range s1 {
		if !Contains(s2, v) {
			return false
		}
	}
	return true
}

func Intersect[T comparable](slice1, slice2 []T) bool {
	for _, item := range slice1 {
		if Contains(slice2, item) {
			return true
		}
	}
	return false
}

func Intersection[T comparable](slice1, slice2 []T) []T {
	smallerSlice := slice1
	largerSlice := slice2
	if len(slice1) > len(slice2) {
		smallerSlice = slice2
		largerSlice = slice1
	}
	var result []T
	for _, item := range smallerSlice {
		if Contains(largerSlice, item) {
			result = append(result, item)
		}
	}
	return result
}

func Find[T any](slice []T, match func(T) bool) (*T, bool) {
	for _, item := range slice {
		if match(item) {
			return &item, true
		}
	}
	return nil, false
}

// IndexOf returns the index of the first occurrence of the target in the slice, or -1 if the target is not present.
// TODO(@guicassolato): unit test
func IndexOf[T comparable](slice []T, target T) int {
	for idx := range slice {
		if slice[idx] == target {
			return idx
		}
	}
	return -1
}

// Map applies the given mapper function to each element in the input slice and returns a new slice with the results.
func Map[T, U any](slice []T, f func(T) U) []U {
	arr := make([]U, len(slice))
	for i, e := range slice {
		arr[i] = f(e)
	}
	return arr
}

// Filter filters the input slice using the given predicate function and returns a new slice with the results.
func Filter[T any](slice []T, f func(T) bool) []T {
	arr := make([]T, 0)
	for _, e := range slice {
		if f(e) {
			arr = append(arr, e)
		}
	}
	return arr
}

// SliceCopy copies the elements from the input slice into the output slice, and returns the output slice.
func SliceCopy[T any](s1 []T) []T {
	s2 := make([]T, len(s1))
	copy(s2, s1)
	return s2
}

// ReverseSlice creates a reversed copy of the input slice.
func ReverseSlice[T any](input []T) []T {
	inputLen := len(input)
	output := make([]T, inputLen)

	for i, n := range input {
		j := inputLen - i - 1

		output[j] = n
	}

	return output
}
