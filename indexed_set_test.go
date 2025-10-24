package goset

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewIndexedSet(t *testing.T) {
	iset := NewIndexedSet[IndexedSetItemUint64](nil)
	assert.Equal(t, []IndexedSetItemUint64(nil), iset.items)

	iset = NewIndexedSet([]IndexedSetItemUint64{})
	assert.Equal(t, []IndexedSetItemUint64(nil), iset.items)

	iset = NewIndexedSet([]IndexedSetItemUint64{1, 2, 4, 7})
	assert.Equal(t, []IndexedSetItemUint64{0, 1, 2, 0, 4, 0, 0, 7}, iset.items)
}

func Test_IndexedSet_JSON(t *testing.T) {
	testcases := [][]uint64{
		[]uint64(nil),
		{},
		{1, 2, 3},
		{1, 3, 7},
	}

	for i, tc := range testcases {
		t.Run(strconv.FormatInt(int64(i), 10), func(t *testing.T) {
			var items []IndexedSetItemUint64
			for _, v := range tc {
				items = append(items, IndexedSetItemUint64(v))
			}

			set := NewIndexedSet(items)

			b, err := json.Marshal(set)
			assert.NoError(t, err)

			var uset IndexedSet[IndexedSetItemUint64]
			err = json.Unmarshal(b, &uset)
			assert.NoError(t, err)
			assert.Equal(t, set, uset)
		})
	}
}

func Test_IndexedSet_Add(t *testing.T) {
	var iset IndexedSet[IndexedSetItemUint64]
	assert.Equal(t, []IndexedSetItemUint64(nil), iset.items)

	iset.Add(IndexedSetItemUint64(3))
	assert.Equal(t, []IndexedSetItemUint64{0, 0, 0, 3}, iset.items)

	iset.Add(IndexedSetItemUint64(2))
	assert.Equal(t, []IndexedSetItemUint64{0, 0, 2, 3}, iset.items)
}

func Test_IndexedSet_Remove(t *testing.T) {
	iset := NewIndexedSet([]IndexedSetItemUint64{1, 2, 4, 7})
	assert.Equal(t, []IndexedSetItemUint64{0, 1, 2, 0, 4, 0, 0, 7}, iset.items)

	iset.Remove(1)
	assert.Equal(t, []IndexedSetItemUint64{0, 0, 2, 0, 4, 0, 0, 7}, iset.items)

	iset.Remove(9)
	assert.Equal(t, []IndexedSetItemUint64{0, 0, 2, 0, 4, 0, 0, 7}, iset.items)
}

func Test_IndexedSet_Get(t *testing.T) {
	iset := NewIndexedSet([]IndexedSetItemUint64{1, 2, 4, 7})
	assert.Equal(t, []IndexedSetItemUint64{0, 1, 2, 0, 4, 0, 0, 7}, iset.items)

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
