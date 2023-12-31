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
