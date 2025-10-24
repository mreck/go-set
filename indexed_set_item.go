package goset

type IndexedSetItem interface {
	Index() uint64
}

type IndexedSetItemUint64 uint64

func (it IndexedSetItemUint64) Index() uint64 { return uint64(it) }
