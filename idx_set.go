package goset

import (
	"encoding/json"
)

type IdxSet[T IdxSetItem] struct {
	maxIdx uint64
	items  []T
}

func NewIdxSet[T IdxSetItem](items []T) IdxSet[T] {
	var s IdxSet[T]

	if len(items) > 0 {
		s.maxIdx = getMaxIndexForIdxSetItems(items)
		s.items = make([]T, s.maxIdx+1)

		for _, it := range items {
			s.items[it.Index()] = it
		}
	}

	return s
}

// MarshalJSON implements the Marshaler interface
func (set IdxSet[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(set.items)
}

// UnmarshalJSON implements the Unmarshaler interface
func (set *IdxSet[T]) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &set.items)
	if err != nil {
		return err
	}

	set.maxIdx = getMaxIndexForIdxSetItems(set.items)

	return nil
}

func (s *IdxSet[T]) Add(item T) {
	idx := item.Index()
	s.maybeGrow(idx)
	s.items[idx] = item
}

func (s *IdxSet[T]) AddMany(items []T) {
	maxIdx := getMaxIndexForIdxSetItems(items)
	s.maybeGrow(maxIdx)

	for _, it := range items {
		s.items[it.Index()] = it
	}
}

func (s *IdxSet[T]) Remove(item T) {
	idx := item.Index()
	if idx > s.maxIdx {
		return
	}

	var nilItem T
	s.items[idx] = nilItem
}

func (s *IdxSet[T]) RemoveMany(items []T) {
	var (
		nilItem T
		idx     uint64
	)

	for _, it := range items {
		idx = it.Index()
		if idx > s.maxIdx {
			continue
		}

		s.items[idx] = nilItem
	}
}

func (s *IdxSet[T]) RemoveIndex(idx uint64) {
	if idx > s.maxIdx {
		return
	}

	var nilItem T
	s.items[idx] = nilItem
}

func (s *IdxSet[T]) RemoveManyIndexes(idxes []uint64) {
	var nilItem T

	for _, idx := range idxes {
		if idx > s.maxIdx {
			continue
		}

		s.items[idx] = nilItem
	}
}

func (s *IdxSet[T]) GetAt(idx uint64) T {
	if idx > s.maxIdx {
		var item T
		return item
	}

	return s.items[idx]
}

func (s *IdxSet[T]) maybeGrow(maxIdx uint64) {
	if maxIdx > s.maxIdx {
		newItems := make([]T, maxIdx+1)
		copy(newItems, s.items)

		s.items = newItems
		s.maxIdx = maxIdx
	}
}

func getMaxIndexForIdxSetItems[T IdxSetItem](items []T) uint64 {
	var maxIdx uint64

	for _, it := range items {
		if it.IsNil() {
			continue
		}

		idx := it.Index()
		if idx > maxIdx {
			maxIdx = idx
		}
	}

	return maxIdx
}
