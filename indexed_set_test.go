package goset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testItem int64

func (i testItem) Index() uint64 {
	return uint64(i - 1)
}

func TestNewIndexedSet(t *testing.T) {
	iset := NewIndexedSet[testItem](nil)
	assert.Equal(t, []testItem(nil), iset.values)

	iset = NewIndexedSet([]testItem{})
	assert.Equal(t, []testItem(nil), iset.values)

	iset = NewIndexedSet([]testItem{1, 2, 4, 7})
	assert.Equal(t, []testItem{1, 2, 0, 4, 0, 0, 7}, iset.values)
}

func Test_IndexedSet_Add(t *testing.T) {
	var iset IndexedSet[testItem]
	assert.Equal(t, []testItem(nil), iset.values)

	iset.Add(testItem(3))
	assert.Equal(t, []testItem{0, 0, 3}, iset.values)

	iset.Add(testItem(2))
	assert.Equal(t, []testItem{0, 2, 3}, iset.values)
}

func Test_IndexedSet_Remove(t *testing.T) {
	iset := NewIndexedSet([]testItem{1, 2, 4, 7})
	assert.Equal(t, []testItem{1, 2, 0, 4, 0, 0, 7}, iset.values)

	iset.Remove(0)
	assert.Equal(t, []testItem{0, 2, 0, 4, 0, 0, 7}, iset.values)

	iset.Remove(9)
	assert.Equal(t, []testItem{0, 2, 0, 4, 0, 0, 7}, iset.values)
}

func Test_IndexedSet_Get(t *testing.T) {
	iset := NewIndexedSet([]testItem{1, 2, 4, 7})
	assert.Equal(t, []testItem{1, 2, 0, 4, 0, 0, 7}, iset.values)

	assert.Equal(t, testItem(1), iset.Get(0))
	assert.Equal(t, testItem(2), iset.Get(1))
	assert.Equal(t, testItem(0), iset.Get(2))
	assert.Equal(t, testItem(4), iset.Get(3))
	assert.Equal(t, testItem(0), iset.Get(4))
	assert.Equal(t, testItem(0), iset.Get(5))
	assert.Equal(t, testItem(7), iset.Get(6))
	assert.Equal(t, testItem(0), iset.Get(7))
	assert.Equal(t, testItem(0), iset.Get(8))
}
