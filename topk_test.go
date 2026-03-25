package main

import (
	"reflect"
	"testing"
)

func TestTopK(t *testing.T) {
	// Test case 1: k = 0 -> return error
	_, err := NewTopK(0, func(a, b int) bool { return a < b })
	if err == nil {
		t.Error("Expected error for k=0")
	}

	// Test case 2: k > total items, max top k
	tk, _ := NewTopK(5, func(a, b int) bool { return a < b })
	data := []int{1, 2, 3}
	for _, v := range data {
		tk.Add(v)
	}
	result := tk.Get()
	expected := []int{3, 2, 1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 3: Simple increasing, min top k
	tk2, _ := NewTopK(3, func(a, b int) bool { return a > b })
	data2 := []int{1, 2, 3, 4, 5}
	for _, v := range data2 {
		tk2.Add(v)
	}
	result2 := tk2.Get()
	expected2 := []int{1, 2, 3}
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Expected %v, got %v", expected2, result2)
	}

	// Test case 4: Duplicate values, max top k
	tk3, _ := NewTopK(3, func(a, b int) bool { return a < b })
	data3 := []int{5, 5, 5, 1, 4}
	for _, v := range data3 {
		tk3.Add(v)
	}
	result3 := tk3.Get()
	expected3 := []int{5, 5, 5}
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("Expected %v, got %v", expected3, result3)
	}

	// Test case 5: Random order, max top k
	tk4, _ := NewTopK(3, func(a, b int) bool { return a < b })
	data4 := []int{3, 1, 5, 2, 4}
	for _, v := range data4 {
		tk4.Add(v)
	}
	result4 := tk4.Get()
	expected4 := []int{5, 4, 3}
	if !reflect.DeepEqual(result4, expected4) {
		t.Errorf("Expected %v, got %v", expected4, result4)
	}

	// Test case 6: All equals, max top k
	tk5, _ := NewTopK(2, func(a, b int) bool { return a < b })
	data5 := []int{2, 2, 2, 2, 2, 2}
	for _, v := range data5 {
		tk5.Add(v)
	}
	result5 := tk5.Get()
	expected5 := []int{2, 2}
	if !reflect.DeepEqual(result5, expected5) {
		t.Errorf("Expected %v, got %v", expected5, result5)
	}

	// Test case 7: Negative num, max top k
	tk6, _ := NewTopK(3, func(a, b int) bool { return a < b })
	data6 := []int{-1, -5, -3, -2}
	for _, v := range data6 {
		tk6.Add(v)
	}
	result6 := tk6.Get()
	expected6 := []int{-1, -2, -3}
	if !reflect.DeepEqual(result6, expected6) {
		t.Errorf("Expected %v, got %v", expected6, result6)
	}

	// Test case 8: Custom struct
	type Item struct {
		val int
	}
	tk7, _ := NewTopK(2, func(a, b Item) bool { return a.val < b.val })
	data7 := []Item{{1}, {3}, {2}}
	for _, v := range data7 {
		tk7.Add(v)
	}
	result7 := tk7.Get()
	expected7 := []Item{{3}, {2}}
	if !reflect.DeepEqual(result7, expected7) {
		t.Errorf("Expected %v, got %v", expected7, result7)
	}

	// Test case 9: Large input, max top k
	tk8, _ := NewTopK(2, func(a, b int) bool { return a < b })
	for i := 0; i <= 1000000; i++ {
		tk8.Add(i)
	}
	result8 := tk8.Get()
	expected8 := []int{1000000, 999999}
	if !reflect.DeepEqual(result8, expected8) {
		t.Errorf("Expected %v, got %v", expected8, result8)
	}
}
