package goset

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewIdxSet(t *testing.T) {
	iset := NewIdxSet[IdxSetItemUint64](nil)
	assert.Equal(t, []IdxSetItemUint64(nil), iset.items)

	iset = NewIdxSet([]IdxSetItemUint64{})
	assert.Equal(t, []IdxSetItemUint64(nil), iset.items)

	iset = NewIdxSet([]IdxSetItemUint64{1, 2, 4, 7})
	assert.Equal(t, []IdxSetItemUint64{1, 2, 0, 4, 0, 0, 7}, iset.items)
}

func Test_IdxSet_JSON(t *testing.T) {
	testcases := [][]uint64{
		[]uint64(nil),
		{},
		{1, 2, 3},
		{1, 3, 7},
	}

	for i, tc := range testcases {
		t.Run(strconv.FormatInt(int64(i), 10), func(t *testing.T) {
			var items []IdxSetItemUint64
			for _, v := range tc {
				items = append(items, IdxSetItemUint64(v))
			}

			set := NewIdxSet(items)

			b, err := json.Marshal(set)
			assert.NoError(t, err)

			var uset IdxSet[IdxSetItemUint64]
			err = json.Unmarshal(b, &uset)
			assert.NoError(t, err)
			assert.Equal(t, set, uset)
		})
	}
}

func Test_IdxSet_Add(t *testing.T) {
	var iset IdxSet[IdxSetItemUint64]
	assert.Equal(t, []IdxSetItemUint64(nil), iset.items)

	iset.Add(IdxSetItemUint64(3))
	assert.Equal(t, []IdxSetItemUint64{0, 0, 3}, iset.items)

	iset.Add(IdxSetItemUint64(2))
	assert.Equal(t, []IdxSetItemUint64{0, 2, 3}, iset.items)
}

func Test_IdxSet_AddMany(t *testing.T) {
	var iset IdxSet[IdxSetItemUint64]
	assert.Equal(t, []IdxSetItemUint64(nil), iset.items)

	iset.AddMany([]IdxSetItemUint64{3, 1})
	assert.Equal(t, []IdxSetItemUint64{1, 0, 3}, iset.items)

	iset.AddMany([]IdxSetItemUint64{7, 5})
	assert.Equal(t, []IdxSetItemUint64{1, 0, 3, 0, 5, 0, 7}, iset.items)
}

func Test_IdxSet_Remove(t *testing.T) {
	iset := NewIdxSet([]IdxSetItemUint64{1, 2, 4, 7})
	assert.Equal(t, []IdxSetItemUint64{1, 2, 0, 4, 0, 0, 7}, iset.items)

	iset.Remove(IdxSetItemUint64(1))
	assert.Equal(t, []IdxSetItemUint64{0, 2, 0, 4, 0, 0, 7}, iset.items)

	iset.Remove(IdxSetItemUint64(9))
	assert.Equal(t, []IdxSetItemUint64{0, 2, 0, 4, 0, 0, 7}, iset.items)
}

func Test_IdxSet_RemoveMany(t *testing.T) {
	iset := NewIdxSet([]IdxSetItemUint64{1, 2, 4, 7})
	assert.Equal(t, []IdxSetItemUint64{1, 2, 0, 4, 0, 0, 7}, iset.items)

	iset.RemoveMany([]IdxSetItemUint64{1, 2})
	assert.Equal(t, []IdxSetItemUint64{0, 0, 0, 4, 0, 0, 7}, iset.items)

	iset.RemoveMany([]IdxSetItemUint64{9, 7, 7})
	assert.Equal(t, []IdxSetItemUint64{0, 0, 0, 4, 0, 0, 0}, iset.items)
}

func Test_IdxSet_RemoveIndex(t *testing.T) {
	iset := NewIdxSet([]IdxSetItemUint64{1, 2, 4, 7})
	assert.Equal(t, []IdxSetItemUint64{1, 2, 0, 4, 0, 0, 7}, iset.items)

	iset.RemoveIndex(IdxSetItemUint64(1).Index())
	assert.Equal(t, []IdxSetItemUint64{0, 2, 0, 4, 0, 0, 7}, iset.items)

	iset.RemoveIndex(9)
	assert.Equal(t, []IdxSetItemUint64{0, 2, 0, 4, 0, 0, 7}, iset.items)
}

func Test_IdxSet_RemoveManyIndexes(t *testing.T) {
	iset := NewIdxSet([]IdxSetItemUint64{1, 2, 4, 7})
	assert.Equal(t, []IdxSetItemUint64{1, 2, 0, 4, 0, 0, 7}, iset.items)

	iset.RemoveManyIndexes([]uint64{0, 3})
	assert.Equal(t, []IdxSetItemUint64{0, 2, 0, 0, 0, 0, 7}, iset.items)

	iset.RemoveManyIndexes([]uint64{9, 8, 6})
	assert.Equal(t, []IdxSetItemUint64{0, 2, 0, 0, 0, 0, 0}, iset.items)
}

func Test_IdxSet_GetAt(t *testing.T) {
	iset := NewIdxSet([]IdxSetItemUint64{1, 2, 4, 7})
	assert.Equal(t, []IdxSetItemUint64{1, 2, 0, 4, 0, 0, 7}, iset.items)

	assert.Equal(t, IdxSetItemUint64(1), iset.GetAt(0))
	assert.Equal(t, IdxSetItemUint64(2), iset.GetAt(1))
	assert.Equal(t, IdxSetItemUint64(0), iset.GetAt(2))
	assert.Equal(t, IdxSetItemUint64(4), iset.GetAt(3))
	assert.Equal(t, IdxSetItemUint64(0), iset.GetAt(4))
	assert.Equal(t, IdxSetItemUint64(0), iset.GetAt(5))
	assert.Equal(t, IdxSetItemUint64(7), iset.GetAt(6))
	assert.Equal(t, IdxSetItemUint64(0), iset.GetAt(7))
	assert.Equal(t, IdxSetItemUint64(0), iset.GetAt(8))
}
