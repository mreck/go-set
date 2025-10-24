package goset

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewIdxSet(t *testing.T) {
	testcases := []struct {
		items  []uint64
		expect []IdxSetItemUint64
	}{
		{
			nil,
			[]IdxSetItemUint64(nil),
		},
		{
			[]uint64{},
			[]IdxSetItemUint64(nil),
		},
		{
			[]uint64{1, 2, 4, 7},
			[]IdxSetItemUint64{1, 2, 0, 4, 0, 0, 7},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("cast:%d", i), func(t *testing.T) {
			items := NewIdxSetItemUint64Slice(tc.items)
			iset := NewIdxSet(items)
			assert.Equal(t, tc.expect, iset.items)
		})
	}
}

func Test_IdxSet_JSON(t *testing.T) {
	testcases := [][]uint64{
		nil,
		{},
		{1, 2, 3},
		{1, 3, 7},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case:%d", i), func(t *testing.T) {
			items := NewIdxSetItemUint64Slice(tc)
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

func Test_IdxSet_Clone(t *testing.T) {
	testcases := []struct {
		values []uint64
	}{
		{[]uint64{1, 2, 3}},
		{[]uint64{1, 3, 7}},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case:%d", i), func(t *testing.T) {
			set := NewIdxSet(NewIdxSetItemUint64Slice(tc.values))
			clone := set.Clone()
			assert.Equal(t, set, clone)
			assert.NotSame(t, &set, &clone)
			assert.NotSame(t, &set.items, &clone.items)
		})
	}
}

func Test_IdxSet_Add(t *testing.T) {
	var iset IdxSet[IdxSetItemUint64]
	assert.Equal(t, []IdxSetItemUint64(nil), iset.items)

	testcases := []struct {
		item   uint64
		expect []IdxSetItemUint64
	}{
		{3, []IdxSetItemUint64{0, 0, 3}},
		{2, []IdxSetItemUint64{0, 2, 3}},
		{3, []IdxSetItemUint64{0, 2, 3}},
		{7, []IdxSetItemUint64{0, 2, 3, 0, 0, 0, 7}},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case:%d", i), func(t *testing.T) {
			iset.Add(IdxSetItemUint64(tc.item))
			assert.Equal(t, tc.expect, iset.items)
		})
	}
}

func Test_IdxSet_AddMany(t *testing.T) {
	var iset IdxSet[IdxSetItemUint64]
	assert.Equal(t, []IdxSetItemUint64(nil), iset.items)

	testcases := []struct {
		items  []uint64
		expect []IdxSetItemUint64
	}{
		{[]uint64{3, 1}, []IdxSetItemUint64{1, 0, 3}},
		{[]uint64{7, 5}, []IdxSetItemUint64{1, 0, 3, 0, 5, 0, 7}},
		{[]uint64{7, 3}, []IdxSetItemUint64{1, 0, 3, 0, 5, 0, 7}},
		{[]uint64{1, 2}, []IdxSetItemUint64{1, 2, 3, 0, 5, 0, 7}},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case:%d", i), func(t *testing.T) {
			items := NewIdxSetItemUint64Slice(tc.items)
			iset.AddMany(items)
			assert.Equal(t, tc.expect, iset.items)
		})
	}
}

func Test_IdxSet_Remove(t *testing.T) {
	iset := NewIdxSet([]IdxSetItemUint64{1, 2, 4, 7})
	assert.Equal(t, []IdxSetItemUint64{1, 2, 0, 4, 0, 0, 7}, iset.items)

	testcases := []struct {
		item   uint64
		expect []IdxSetItemUint64
	}{
		{1, []IdxSetItemUint64{0, 2, 0, 4, 0, 0, 7}},
		{9, []IdxSetItemUint64{0, 2, 0, 4, 0, 0, 7}},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case:%d", i), func(t *testing.T) {
			iset.Remove(IdxSetItemUint64(tc.item))
			assert.Equal(t, tc.expect, iset.items)
		})
	}
}

func Test_IdxSet_RemoveMany(t *testing.T) {
	iset := NewIdxSet([]IdxSetItemUint64{1, 2, 4, 7})
	assert.Equal(t, []IdxSetItemUint64{1, 2, 0, 4, 0, 0, 7}, iset.items)

	testcases := []struct {
		items  []uint64
		expect []IdxSetItemUint64
	}{
		{[]uint64{1, 2}, []IdxSetItemUint64{0, 0, 0, 4, 0, 0, 7}},
		{[]uint64{9, 7}, []IdxSetItemUint64{0, 0, 0, 4, 0, 0, 0}},
		{[]uint64{7, 7}, []IdxSetItemUint64{0, 0, 0, 4, 0, 0, 0}},
		{[]uint64{4, 1}, []IdxSetItemUint64{0, 0, 0, 0, 0, 0, 0}},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case:%d", i), func(t *testing.T) {
			items := NewIdxSetItemUint64Slice(tc.items)
			iset.RemoveMany(items)
			assert.Equal(t, tc.expect, iset.items)
		})
	}
}

func Test_IdxSet_RemoveIndex(t *testing.T) {
	iset := NewIdxSet([]IdxSetItemUint64{1, 2, 4, 7})
	assert.Equal(t, []IdxSetItemUint64{1, 2, 0, 4, 0, 0, 7}, iset.items)

	testcases := []struct {
		index  uint64
		expect []IdxSetItemUint64
	}{
		{0, []IdxSetItemUint64{0, 2, 0, 4, 0, 0, 7}},
		{9, []IdxSetItemUint64{0, 2, 0, 4, 0, 0, 7}},
		{1, []IdxSetItemUint64{0, 0, 0, 4, 0, 0, 7}},
		{1, []IdxSetItemUint64{0, 0, 0, 4, 0, 0, 7}},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case:%d", i), func(t *testing.T) {
			iset.RemoveIndex(tc.index)
			assert.Equal(t, tc.expect, iset.items)
		})
	}
}

func Test_IdxSet_RemoveManyIndexes(t *testing.T) {
	iset := NewIdxSet([]IdxSetItemUint64{1, 2, 4, 7})
	assert.Equal(t, []IdxSetItemUint64{1, 2, 0, 4, 0, 0, 7}, iset.items)

	testcases := []struct {
		indexes []uint64
		expect  []IdxSetItemUint64
	}{
		{[]uint64{0, 3}, []IdxSetItemUint64{0, 2, 0, 0, 0, 0, 7}},
		{[]uint64{3, 0}, []IdxSetItemUint64{0, 2, 0, 0, 0, 0, 7}},
		{[]uint64{1, 1}, []IdxSetItemUint64{0, 0, 0, 0, 0, 0, 7}},
		{[]uint64{9, 9}, []IdxSetItemUint64{0, 0, 0, 0, 0, 0, 7}},
		{[]uint64{9, 6}, []IdxSetItemUint64{0, 0, 0, 0, 0, 0, 0}},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case:%d", i), func(t *testing.T) {
			iset.RemoveManyIndexes(tc.indexes)
			assert.Equal(t, tc.expect, iset.items)
		})
	}
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
