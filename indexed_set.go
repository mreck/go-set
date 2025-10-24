package goset

type IndexedSetItem interface {
	Index() uint64
}

type IndexedSet[T IndexedSetItem] struct {
	maxIndex uint64
	values   []T
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

func (s *IndexedSet[T]) maybeGrow(maxIndex uint64) {
	if maxIndex > s.maxIndex {
		newValues := make([]T, maxIndex+1)
		newValues = append(newValues, s.values...)
		s.values = newValues
		s.maxIndex = maxIndex
	}
}

func (s *IndexedSet[T]) Add(item T) {
	idx := item.Index()
	s.maybeGrow(idx)
	s.values[item.Index()] = item
}

func (s *IndexedSet[T]) Remove(index uint64) {
	if index > s.maxIndex {
		return
	}
	var item T
	s.values[index] = item
}

func (s *IndexedSet[T]) Get(index uint64) T {
	if index > s.maxIndex {
		var item T
		return item
	}
	return s.values[index]
}
