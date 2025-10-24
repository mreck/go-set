package goset

type IdxSetItem interface {
	Index() uint64
	IsNil() bool
}

type IdxSetItemUint64 uint64

func (it IdxSetItemUint64) Index() uint64 { return uint64(it - 1) }
func (it IdxSetItemUint64) IsNil() bool   { return it == 0 }
