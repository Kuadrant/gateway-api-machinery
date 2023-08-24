package common

import (
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []string
		target   string
		expected bool
	}{
		{
			name:     "when slice has one target item then return true",
			slice:    []string{"test-gw"},
			target:   "test-gw",
			expected: true,
		},
		{
			name:     "when slice is empty then return false",
			slice:    []string{},
			target:   "test-gw",
			expected: false,
		},
		{
			name:     "when target is in a slice then return true",
			slice:    []string{"test-gw1", "test-gw2", "test-gw3"},
			target:   "test-gw2",
			expected: true,
		},
		{
			name:     "when no target in a slice then return false",
			slice:    []string{"test-gw1", "test-gw2", "test-gw3"},
			target:   "test-gw4",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if Contains(tc.slice, tc.target) != tc.expected {
				t.Errorf("when slice=%v and target=%s, expected=%v, but got=%v", tc.slice, tc.target, tc.expected, !tc.expected)
			}
		})
	}
}

func TestContainsWithInts(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []int
		target   int
		expected bool
	}{
		{
			name:     "when slice has one target item then return true",
			slice:    []int{1},
			target:   1,
			expected: true,
		},
		{
			name:     "when slice is empty then return false",
			slice:    []int{},
			target:   2,
			expected: false,
		},
		{
			name:     "when target is in a slice then return true",
			slice:    []int{1, 2, 3},
			target:   2,
			expected: true,
		},
		{
			name:     "when no target in a slice then return false",
			slice:    []int{1, 2, 3},
			target:   4,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if Contains(tc.slice, tc.target) != tc.expected {
				t.Errorf("when slice=%v and target=%d, expected=%v, but got=%v", tc.slice, tc.target, tc.expected, !tc.expected)
			}
		})
	}
}

func TestSameElements(t *testing.T) {
	testCases := []struct {
		name     string
		slice1   []string
		slice2   []string
		expected bool
	}{
		{
			name:     "when slice1 and slice2 contain the same elements then return true",
			slice1:   []string{"test-gw1", "test-gw2", "test-gw3"},
			slice2:   []string{"test-gw1", "test-gw2", "test-gw3"},
			expected: true,
		},
		{
			name:     "when slice1 and slice2 contain unique elements then return false",
			slice1:   []string{"test-gw1", "test-gw2"},
			slice2:   []string{"test-gw1", "test-gw3"},
			expected: false,
		},
		{
			name:     "when both slices are empty then return true",
			slice1:   []string{},
			slice2:   []string{},
			expected: true,
		},
		{
			name:     "when both slices are nil then return true",
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if SameElements(tc.slice1, tc.slice2) != tc.expected {
				t.Errorf("when slice1=%v and slice2=%v, expected=%v, but got=%v", tc.slice1, tc.slice2, tc.expected, !tc.expected)
			}
		})
	}
}

func TestIntersect(t *testing.T) {
	testCases := []struct {
		name     string
		slice1   []string
		slice2   []string
		expected bool
	}{
		{
			name:     "when slice1 and slice2 have one common item then return true",
			slice1:   []string{"test-gw1", "test-gw2"},
			slice2:   []string{"test-gw1", "test-gw3", "test-gw4"},
			expected: true,
		},
		{
			name:     "when slice1 and slice2 have no common item then return false",
			slice1:   []string{"test-gw1", "test-gw2"},
			slice2:   []string{"test-gw3", "test-gw4"},
			expected: false,
		},
		{
			name:     "when slice1 is empty then return false",
			slice1:   []string{},
			slice2:   []string{"test-gw3", "test-gw4"},
			expected: false,
		},
		{
			name:     "when slice2 is empty then return false",
			slice1:   []string{"test-gw1", "test-gw2"},
			slice2:   []string{},
			expected: false,
		},
		{
			name:     "when both slices are empty then return false",
			slice1:   []string{},
			slice2:   []string{},
			expected: false,
		},
		{
			name:     "when slice1 is nil then return false",
			slice2:   []string{"test-gw3", "test-gw4"},
			expected: false,
		},
		{
			name:     "when slice2 is nil then return false",
			slice1:   []string{"test-gw1", "test-gw2"},
			expected: false,
		},
		{
			name:     "when both slices are nil then return false",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if Intersect(tc.slice1, tc.slice2) != tc.expected {
				t.Errorf("when slice1=%v and slice2=%v, expected=%v, but got=%v", tc.slice1, tc.slice2, tc.expected, !tc.expected)
			}
		})
	}
}

func TestIntersectWithInts(t *testing.T) {
	testCases := []struct {
		name     string
		slice1   []int
		slice2   []int
		expected bool
	}{
		{
			name:     "when slice1 and slice2 have one common item then return true",
			slice1:   []int{1, 2},
			slice2:   []int{1, 3, 4},
			expected: true,
		},
		{
			name:     "when slice1 and slice2 have no common item then return false",
			slice1:   []int{1, 2},
			slice2:   []int{3, 4},
			expected: false,
		},
		{
			name:     "when slice1 is empty then return false",
			slice1:   []int{},
			slice2:   []int{3, 4},
			expected: false,
		},
		{
			name:     "when slice2 is empty then return false",
			slice1:   []int{1, 2},
			slice2:   []int{},
			expected: false,
		},
		{
			name:     "when both slices are empty then return false",
			slice1:   []int{},
			slice2:   []int{},
			expected: false,
		},
		{
			name:     "when slice1 is nil then return false",
			slice2:   []int{3, 4},
			expected: false,
		},
		{
			name:     "when slice2 is nil then return false",
			slice1:   []int{1, 2},
			expected: false,
		},
		{
			name:     "when both slices are nil then return false",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if Intersect(tc.slice1, tc.slice2) != tc.expected {
				t.Errorf("when slice1=%v and slice2=%v, expected=%v, but got=%v", tc.slice1, tc.slice2, tc.expected, !tc.expected)
			}
		})
	}
}

func TestIntersection(t *testing.T) {
	testCases := []struct {
		name     string
		slice1   []string
		slice2   []string
		expected []string
	}{
		{
			name:     "when slice1 and slice2 have one common item then return that item",
			slice1:   []string{"test-gw1", "test-gw2"},
			slice2:   []string{"test-gw1", "test-gw3", "test-gw4"},
			expected: []string{"test-gw1"},
		},
		{
			name:     "when slice1 and slice2 have no common item then return nil",
			slice1:   []string{"test-gw1", "test-gw2"},
			slice2:   []string{"test-gw3", "test-gw4"},
			expected: nil,
		},
		{
			name:     "when slice1 is empty then return nil",
			slice1:   []string{},
			slice2:   []string{"test-gw3", "test-gw4"},
			expected: nil,
		},
		{
			name:     "when slice2 is empty then return nil",
			slice1:   []string{"test-gw1", "test-gw2"},
			slice2:   []string{},
			expected: nil,
		},
		{
			name:     "when both slices are empty then return nil",
			slice1:   []string{},
			slice2:   []string{},
			expected: nil,
		},
		{
			name:     "when slice1 is nil then return nil",
			slice2:   []string{"test-gw3", "test-gw4"},
			expected: nil,
		},
		{
			name:     "when slice2 is nil then return nil",
			slice1:   []string{"test-gw1", "test-gw2"},
			expected: nil,
		},
		{
			name:     "when both slices are nil then return nil",
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if r := Intersection(tc.slice1, tc.slice2); !reflect.DeepEqual(r, tc.expected) {
				t.Errorf("expected=%v; got=%v", tc.expected, r)
			}
		})
	}
}

func TestFind(t *testing.T) {
	s := []string{"a", "ab", "abc"}

	if r, found := Find(s, func(el string) bool { return el == "ab" }); !found || r == nil || *r != "ab" {
		t.Error("should have found 'ab' in the slice")
	}

	if r, found := Find(s, func(el string) bool { return len(el) <= 3 }); !found || r == nil || *r != "a" {
		t.Error("should have found 'a' in the slice")
	}

	if r, found := Find(s, func(el string) bool { return len(el) >= 3 }); !found || r == nil || *r != "abc" {
		t.Error("should have found 'abc' in the slice")
	}

	if r, found := Find(s, func(el string) bool { return len(el) == 4 }); found || r != nil {
		t.Error("should not have found anything in the slice")
	}

	i := []int{1, 2, 3}

	if r, found := Find(i, func(el int) bool { return el/3 == 1 }); !found || r == nil || *r != 3 {
		t.Error("should have found 3 in the slice")
	}

	if r, found := Find(i, func(el int) bool { return el == 75 }); found || r != nil {
		t.Error("should not have found anything in the slice")
	}
}

func TestMap(t *testing.T) {
	slice1 := []int{1, 2, 3, 4}
	f1 := func(x int) int { return x + 1 }
	expected1 := []int{2, 3, 4, 5}
	result1 := Map(slice1, f1)
	t.Run("when mapping an int slice with an increment function then return new slice with the incremented values", func(t *testing.T) {
		if !reflect.DeepEqual(result1, expected1) {
			t.Errorf("result1 = %v; expected %v", result1, expected1)
		}
	})

	slice2 := []string{"hello", "world", "buz", "a"}
	f2 := func(s string) int { return len(s) }
	expected2 := []int{5, 5, 3, 1}
	result2 := Map(slice2, f2)
	t.Run("when mapping a string slice with string->int mapping then return new slice with the mapped values", func(t *testing.T) {
		if !reflect.DeepEqual(result2, expected2) {
			t.Errorf("result2 = %v; expected %v", result2, expected2)
		}
	})

	slice3 := []int{}
	f3 := func(x int) float32 { return float32(x) / 2 }
	expected3 := []float32{}
	result3 := Map(slice3, f3)
	t.Run("when mapping an empty int slice then return an empty slice", func(t *testing.T) {
		if !reflect.DeepEqual(result3, expected3) {
			t.Errorf("result3 = %v; expected %v", result3, expected3)
		}
	})
}

func TestSliceCopy(t *testing.T) {
	input1 := []int{1, 2, 3}
	expected1 := []int{1, 2, 3}
	output1 := SliceCopy(input1)
	t.Run("when given slice of integers then return a copy of the input slice", func(t *testing.T) {
		if !reflect.DeepEqual(output1, expected1) {
			t.Errorf("SliceCopy(%v) = %v; expected %v", input1, output1, expected1)
		}
	})

	input2 := []string{"foo", "bar", "baz"}
	expected2 := []string{"foo", "bar", "baz"}
	output2 := SliceCopy(input2)
	t.Run("when given slice of strings then return a copy of the input slice", func(t *testing.T) {
		if !reflect.DeepEqual(output2, expected2) {
			t.Errorf("SliceCopy(%v) = %v; expected %v", input2, output2, expected2)
		}
	})

	type person struct {
		name string
		age  int
	}
	input3 := []person{{"Artem", 65}, {"DD", 18}, {"Charlie", 23}}
	expected3 := []person{{"Artem", 65}, {"DD", 18}, {"Charlie", 23}}
	output3 := SliceCopy(input3)
	t.Run("when given slice of structs then return a copy of the input slice", func(t *testing.T) {
		if !reflect.DeepEqual(output3, expected3) {
			t.Errorf("SliceCopy(%v) = %v; expected %v", input3, output3, expected3)
		}
	})

	input4 := []int{1, 2, 3}
	expected4 := []int{1, 2, 3}
	output4 := SliceCopy(input4)
	t.Run("when modifying the original input slice then does not affect the returned copy", func(t *testing.T) {
		if !reflect.DeepEqual(output4, expected4) {
			t.Errorf("SliceCopy(%v) = %v; expected %v", input4, output4, expected4)
		}
		input4[0] = 4
		if reflect.DeepEqual(output4, input4) {
			t.Errorf("modifying the original input slice should not change the output slice")
		}
	})
}

func TestReverseSlice(t *testing.T) {
	input1 := []int{1, 2, 3}
	expected1 := []int{3, 2, 1}
	output1 := ReverseSlice(input1)
	t.Run("when given slice of integers then return reversed copy of the input slice", func(t *testing.T) {
		if !reflect.DeepEqual(output1, expected1) {
			t.Errorf("ReverseSlice(%v) = %v; expected %v", input1, output1, expected1)
		}
	})

	input2 := []string{"foo", "bar", "baz"}
	expected2 := []string{"baz", "bar", "foo"}
	output2 := ReverseSlice(input2)
	t.Run("when given slice of strings then return reversed copy of the input slice", func(t *testing.T) {
		if !reflect.DeepEqual(output2, expected2) {
			t.Errorf("ReverseSlice(%v) = %v; expected %v", input2, output2, expected2)
		}
	})

	input3 := []int{}
	expected3 := []int{}
	output3 := ReverseSlice(input3)
	t.Run("when given an empty slice then return empty slice", func(t *testing.T) {
		if !reflect.DeepEqual(output3, expected3) {
			t.Errorf("ReverseSlice(%v) = %v; expected %v", input3, output3, expected3)
		}
	})
}
