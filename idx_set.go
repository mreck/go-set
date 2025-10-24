package goset

import (
	"encoding/json"
)

// IdxSet holds the values for a set of type T in an array.
type IdxSet[T IdxSetItem] struct {
	maxIdx uint64
	items  []T
}

// NewIdxSet creates a new IdxSet.
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

// Clone creates a copy of the set with it's own array of items.
func (s IdxSet[T]) Clone() IdxSet[T] {
	var clone IdxSet[T]
	clone.maxIdx = s.maxIdx
	clone.items = make([]T, len(s.items))
	copy(clone.items, s.items)
	return clone
}

// Values returns the non-nil values in the set.
// The values are sorted by ascending index.
func (s IdxSet[T]) Values() []T {
	var values []T

	for _, it := range s.items {
		if !it.IsNil() {
			values = append(values, it)
		}
	}

	return values
}

// MarshalJSON implements the Marshaler interface.
func (set IdxSet[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(set.items)
}

// UnmarshalJSON implements the Unmarshaler interface.
func (set *IdxSet[T]) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &set.items)
	if err != nil {
		return err
	}

	set.maxIdx = getMaxIndexForIdxSetItems(set.items)

	return nil
}

// Add adds the item to the set.
// It will override an exiting item with the same index.
func (s *IdxSet[T]) Add(item T) {
	idx := item.Index()
	s.maybeGrow(idx)
	s.items[idx] = item
}

// AddMany adds the items to the set.
// It will override an exiting items with the same index.
func (s *IdxSet[T]) AddMany(items []T) {
	maxIdx := getMaxIndexForIdxSetItems(items)
	s.maybeGrow(maxIdx)

	for _, it := range items {
		s.items[it.Index()] = it
	}
}

// Remove removes the item from the set.
// It will override it with the zero value of the items type.
func (s *IdxSet[T]) Remove(item T) {
	idx := item.Index()
	if idx > s.maxIdx {
		return
	}

	var nilItem T
	s.items[idx] = nilItem
}

// Remove removes the items from the set.
// It will override them with the zero value of the items type.
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

// RemoveIndex remove the item at the index from the set.
// It will override it with the zero value of the items type.
func (s *IdxSet[T]) RemoveIndex(idx uint64) {
	if idx > s.maxIdx {
		return
	}

	var nilItem T
	s.items[idx] = nilItem
}

// RemoveManyIndexes removes the items at the indexes from the set.
// It will override them with the zero value of the items type.
func (s *IdxSet[T]) RemoveManyIndexes(idxes []uint64) {
	var nilItem T

	for _, idx := range idxes {
		if idx > s.maxIdx {
			continue
		}

		s.items[idx] = nilItem
	}
}

// GetAt returns the item at the index.
// It will the zero value of the items type by default.
func (s *IdxSet[T]) GetAt(idx uint64) T {
	if idx > s.maxIdx {
		var item T
		return item
	}

	return s.items[idx]
}

func (s *IdxSet[T]) maybeGrow(maxIdx uint64) {
	if len(s.items) == 0 || maxIdx > s.maxIdx {
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
