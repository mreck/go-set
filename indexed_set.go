package goset

import (
	"encoding/json"
)

type IndexedSet[T IndexedSetItem] struct {
	maxIndex uint64
	items    []T
}

func NewIndexedSet[T IndexedSetItem](items []T) IndexedSet[T] {
	var (
		maxIndex uint64
		values   []T
	)

	if len(items) > 0 {
		for _, it := range items {
			idx := it.Index()
			if idx > maxIndex {
				maxIndex = idx
			}
		}

		values = make([]T, maxIndex+1)

		for _, it := range items {
			values[it.Index()] = it
		}
	}

	return IndexedSet[T]{maxIndex, values}
}

// MarshalJSON implements the Marshaler interface
func (set IndexedSet[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(set.items)
}

// UnmarshalJSON implements the Unmarshaler interface
func (set *IndexedSet[T]) UnmarshalJSON(b []byte) error {
	var (
		maxIndex uint64
		items    []T
	)

	err := json.Unmarshal(b, &items)
	if err != nil {
		return err
	}

	if len(items) > 0 {
		for _, it := range items {
			if it.IsNil() {
				continue
			}
			idx := it.Index()
			if idx > maxIndex {
				maxIndex = idx
			}
		}
	}

	set.maxIndex = maxIndex
	set.items = items

	return nil
}

func (s *IndexedSet[T]) maybeGrow(maxIndex uint64) {
	if maxIndex > s.maxIndex {
		newValues := make([]T, maxIndex+1)
		newValues = append(newValues, s.items...)
		s.items = newValues
		s.maxIndex = maxIndex
	}
}

func (s *IndexedSet[T]) Add(item T) {
	idx := item.Index()
	s.maybeGrow(idx)
	s.items[item.Index()] = item
}

// @TODO: AddMany()

func (s *IndexedSet[T]) Remove(index uint64) {
	if index > s.maxIndex {
		return
	}
	var item T
	s.items[index] = item
}

// @TODO: RemoveMany()

func (s *IndexedSet[T]) Get(index uint64) T {
	if index > s.maxIndex {
		var item T
		return item
	}
	return s.items[index]
}
