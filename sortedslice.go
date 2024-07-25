// Copyright © 2024 Mark Summerfield. All rights reserved.

// This package provides a sorted slice, maintaining order using binary
// search.
package sortedslice

import (
	"cmp"
	_ "embed"
	"iter"
	"slices"
)

//go:embed Version.dat
var Version string

type SortedSlice[E cmp.Ordered] struct{ slice []E }

func New[E cmp.Ordered](capacity int) SortedSlice[E] {
	return SortedSlice[E]{make([]E, 0, capacity)}
}

func (me *SortedSlice[E]) Len() int { return len(me.slice) }

func (me *SortedSlice[E]) Equal(other SortedSlice[E]) bool {
	return slices.Equal(me.slice, other.slice)
}

func (me *SortedSlice[E]) Insert(element E) bool {
	if i, ok := slices.BinarySearch(me.slice, element); ok {
		return false // already present
	} else {
		me.slice = slices.Insert(me.slice, i, element)
		return true // inserted
	}
}

func (me *SortedSlice[E]) Delete(element E) bool {
	if i, ok := slices.BinarySearch(me.slice, element); ok {
		me.slice = slices.Delete(me.slice, i, i+1)
		return true // deleted
	}
	return false // not present so not deleted
}

func (me *SortedSlice[E]) Find(element E) int {
	if i, ok := slices.BinarySearch(me.slice, element); ok {
		return i // found at position i
	}
	return -1 // not found
}

func (me *SortedSlice[E]) All() iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		for i, element := range me.slice {
			if !yield(i, element) {
				return
			}
		}
	}
}
