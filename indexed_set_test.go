package goset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIndexedSet(t *testing.T) {
	iset := NewIndexedSet[IndexedSetItemUint64](nil)
	assert.Equal(t, []IndexedSetItemUint64(nil), iset.values)

	iset = NewIndexedSet([]IndexedSetItemUint64{})
	assert.Equal(t, []IndexedSetItemUint64(nil), iset.values)

	iset = NewIndexedSet([]IndexedSetItemUint64{1, 2, 4, 7})
	assert.Equal(t, []IndexedSetItemUint64{0, 1, 2, 0, 4, 0, 0, 7}, iset.values)
}

func Test_IndexedSet_Add(t *testing.T) {
	var iset IndexedSet[IndexedSetItemUint64]
	assert.Equal(t, []IndexedSetItemUint64(nil), iset.values)

	iset.Add(IndexedSetItemUint64(3))
	assert.Equal(t, []IndexedSetItemUint64{0, 0, 0, 3}, iset.values)

	iset.Add(IndexedSetItemUint64(2))
	assert.Equal(t, []IndexedSetItemUint64{0, 0, 2, 3}, iset.values)
}

func Test_IndexedSet_Remove(t *testing.T) {
	iset := NewIndexedSet([]IndexedSetItemUint64{1, 2, 4, 7})
	assert.Equal(t, []IndexedSetItemUint64{0, 1, 2, 0, 4, 0, 0, 7}, iset.values)

	iset.Remove(1)
	assert.Equal(t, []IndexedSetItemUint64{0, 0, 2, 0, 4, 0, 0, 7}, iset.values)

	iset.Remove(9)
	assert.Equal(t, []IndexedSetItemUint64{0, 0, 2, 0, 4, 0, 0, 7}, iset.values)
}

func Test_IndexedSet_Get(t *testing.T) {
	iset := NewIndexedSet([]IndexedSetItemUint64{1, 2, 4, 7})
	assert.Equal(t, []IndexedSetItemUint64{0, 1, 2, 0, 4, 0, 0, 7}, iset.values)

	assert.Equal(t, IndexedSetItemUint64(0), iset.Get(0))
	assert.Equal(t, IndexedSetItemUint64(1), iset.Get(1))
	assert.Equal(t, IndexedSetItemUint64(2), iset.Get(2))
	assert.Equal(t, IndexedSetItemUint64(0), iset.Get(3))
	assert.Equal(t, IndexedSetItemUint64(4), iset.Get(4))
	assert.Equal(t, IndexedSetItemUint64(0), iset.Get(5))
	assert.Equal(t, IndexedSetItemUint64(0), iset.Get(6))
	assert.Equal(t, IndexedSetItemUint64(7), iset.Get(7))
	assert.Equal(t, IndexedSetItemUint64(0), iset.Get(8))
}
