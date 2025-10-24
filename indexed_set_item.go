package goset

type IndexedSetItem interface {
	Index() uint64
	IsNil() bool
}

type IndexedSetItemUint64 uint64

func (it IndexedSetItemUint64) Index() uint64 { return uint64(it - 1) }
func (it IndexedSetItemUint64) IsNil() bool   { return it == 0 }
