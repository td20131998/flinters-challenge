package main

import (
	"errors"
	"sort"
)

type itemHeap[T any] struct {
	items []T
	less  func(T, T) bool
}

func (h itemHeap[T]) Len() int           { return len(h.items) }
func (h itemHeap[T]) Less(i, j int) bool { return h.less(h.items[i], h.items[j]) }
func (h itemHeap[T]) Swap(i, j int)      { h.items[i], h.items[j] = h.items[i], h.items[j] }
func (h *itemHeap[T]) Push(x any) {
	h.items = append(h.items, x.(T))
	h.up(len(h.items) - 1)
}
func (h *itemHeap[T]) Pop() any {
	if len(h.items) == 0 {
		return nil
	}
	n := len(h.items) - 1
	h.items[0], h.items[n] = h.items[n], h.items[0]
	item := h.items[n]
	h.items = h.items[:n]
	if n > 0 {
		h.down(0, n)
	}
	return item
}

func (h *itemHeap[T]) up(i int) {
	for i > 0 {
		j := (i - 1) / 2
		if !h.less(h.items[i], h.items[j]) {
			break
		}
		h.items[i], h.items[j] = h.items[j], h.items[i]
		i = j
	}
}

func (h *itemHeap[T]) down(i, n int) {
	for {
		j1 := 2*i + 1
		if j1 >= n {
			break
		}
		j := j1
		if j2 := j1 + 1; j2 < n && h.less(h.items[j2], h.items[j1]) {
			j = j2
		}
		if !h.less(h.items[j], h.items[i]) {
			break
		}
		h.items[i], h.items[j] = h.items[j], h.items[i]
		i = j
	}
}

type TopK[T any] struct {
	heap *itemHeap[T]
	k    int
	less func(T, T) bool
}

func NewTopK[T any](k int, less func(T, T) bool) (*TopK[T], error) {
	if k <= 0 {
		return nil, errors.New("k must be positive")
	}
	return &TopK[T]{
		heap: &itemHeap[T]{less: less},
		k:    k,
		less: less,
	}, nil
}

func (tk *TopK[T]) Add(item T) {
	if tk.heap.Len() < tk.k {
		tk.heap.Push(item)
	} else if tk.less(tk.heap.items[0], item) {
		tk.heap.Pop()
		tk.heap.Push(item)
	}
}

func (tk *TopK[T]) Get() []T {
	items := make([]T, 0, tk.heap.Len())
	for tk.heap.Len() > 0 {
		items = append(items, tk.heap.Pop().(T))
	}
	sort.Slice(items, func(i, j int) bool { return tk.less(items[i], items[j]) })
	// Reverse to get the correct order
	for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
		items[i], items[j] = items[j], items[i]
	}
	return items
}
