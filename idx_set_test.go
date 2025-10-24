package goset

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewIdxSet(t *testing.T) {
	testCases := []struct {
		items  []IdxSetItemUint64
		expect []IdxSetItemUint64
	}{
		{
			items:  nil,
			expect: []IdxSetItemUint64(nil),
		},
		{
			items:  []IdxSetItemUint64{},
			expect: []IdxSetItemUint64(nil),
		},
		{
			items:  []IdxSetItemUint64{1, 2, 4, 7},
			expect: []IdxSetItemUint64{1, 2, 0, 4, 0, 0, 7},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("cast:%d", i), func(t *testing.T) {
			iset := NewIdxSet(tc.items)
			assert.Equal(t, tc.expect, iset.items)
		})
	}
}

func Test_IdxSet_JSON(t *testing.T) {
	testcases := [][]uint64{
		[]uint64(nil),
		{},
		{1, 2, 3},
		{1, 3, 7},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case:%d", i), func(t *testing.T) {
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

	testcases := []struct {
		index  uint64
		expect uint64
	}{
		{0, 1},
		{1, 2},
		{2, 0},
		{3, 4},
		{4, 0},
		{5, 0},
		{6, 7},
		{7, 0},
		{8, 0},
	}
	for _, tc := range testcases {
		t.Run(fmt.Sprintf("index:%d", tc.index), func(t *testing.T) {
			assert.Equal(t, IdxSetItemUint64(tc.expect), iset.GetAt(tc.index))
		})
	}
}
