package set

import (
	"encoding/json"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalJSON(t *testing.T) {
	data := []int{1, 2, 4, 5}
	input := New(data)

	b, err := json.Marshal(input)
	assert.Nil(t, err)

	var output []int

	err = json.Unmarshal(b, &output)
	assert.Nil(t, err)

	slices.Sort(output)
	assert.Equal(t, data, output)
}

func TestUnmarshalJSON(t *testing.T) {
	data := []int{1, 2, 4, 5}

	b, err := json.Marshal(data)
	assert.Nil(t, err)

	var output Set[int]

	err = json.Unmarshal(b, &output)
	assert.Nil(t, err)

	assert.Equal(t, New(data), output)
}

func TestRange(t *testing.T) {
	var empty Set[int]
	var collected []int
	for v := range empty.Range() {
		collected = append(collected, v)
	}
	assert.Nil(t, collected)

	data := []int{1, 2, 3, 4, 5}
	a := New(data)

	collected = nil
	for v := range a.Range() {
		collected = append(collected, v)
	}
	slices.Sort(collected)
	assert.Equal(t, data, collected)

	count := 0
	for range a.Range() {
		count++
		if count == 2 {
			break
		}
	}
	assert.Equal(t, 2, count)
}

func TestNewFromSeq(t *testing.T) {
	data := []int{1, 2, 4, 5}
	assert.Equal(t, New(data), NewFromSeq(slices.Values(data)))
	assert.Equal(t, New([]int(nil)), NewFromSeq(slices.Values([]int(nil))))
}

func TestDifference(t *testing.T) {
	var a, b Set[int]

	assert.Equal(t, New([]int(nil)), Difference(a, b))
	assert.Equal(t, New([]int(nil)), Difference(b, a))

	a = New([]int{1, 2, 3, 4, 5, 6})

	assert.Equal(t, New([]int{1, 2, 3, 4, 5, 6}), Difference(a, b))
	assert.Equal(t, New([]int(nil)), Difference(b, a))

	b = New([]int{2, 4, 6})

	assert.Equal(t, New([]int{1, 3, 5}), Difference(a, b))
	assert.Equal(t, New([]int(nil)), Difference(b, a))
}

func TestIntersection(t *testing.T) {
	var a, b Set[int]

	assert.Equal(t, New([]int(nil)), Intersection(a, b))
	assert.Equal(t, New([]int(nil)), Intersection(b, a))

	a = New([]int{1, 2, 3, 4, 5, 6})

	assert.Equal(t, New([]int(nil)), Intersection(a, b))
	assert.Equal(t, New([]int(nil)), Intersection(b, a))

	b = New([]int{2, 4, 6, 8, 10})

	assert.Equal(t, New([]int{2, 4, 6}), Intersection(a, b))
	assert.Equal(t, New([]int{2, 4, 6}), Intersection(b, a))
}

func TestSymmetricDifference(t *testing.T) {
	var a, b Set[int]

	assert.Equal(t, New([]int(nil)), SymmetricDifference(a, b))
	assert.Equal(t, New([]int(nil)), SymmetricDifference(b, a))

	a = New([]int{1, 2, 3, 4, 5, 6})
	assert.Equal(t, New([]int{1, 2, 3, 4, 5, 6}), SymmetricDifference(a, b))
	assert.Equal(t, New([]int{1, 2, 3, 4, 5, 6}), SymmetricDifference(b, a))

	b = New([]int{4, 5, 6, 7, 8, 9})

	assert.Equal(t, New([]int{1, 2, 3, 7, 8, 9}), SymmetricDifference(a, b))
	assert.Equal(t, New([]int{1, 2, 3, 7, 8, 9}), SymmetricDifference(b, a))
}

func TestUnion(t *testing.T) {
	var a, b Set[int]

	assert.Equal(t, New([]int(nil)), Union(a, b))
	assert.Equal(t, New([]int(nil)), Union(b, a))

	a = New([]int{1, 2, 3, 4, 5, 6})

	assert.Equal(t, New([]int{1, 2, 3, 4, 5, 6}), Union(a, b))
	assert.Equal(t, New([]int{1, 2, 3, 4, 5, 6}), Union(b, a))

	b = New([]int{4, 5, 6, 7, 8, 9})

	assert.Equal(t, New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}), Union(a, b))
	assert.Equal(t, New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}), Union(b, a))
}

func TestIsDisjoint(t *testing.T) {
	var a, b Set[int]

	assert.True(t, IsDisjoint(a, b))
	assert.True(t, IsDisjoint(b, a))

	a = New([]int{1, 2, 3})

	assert.True(t, IsDisjoint(a, b))
	assert.True(t, IsDisjoint(b, a))

	b = New([]int{4, 5, 6})

	assert.True(t, IsDisjoint(a, b))
	assert.True(t, IsDisjoint(b, a))

	b = New([]int{3, 4, 5})

	assert.False(t, IsDisjoint(a, b))
	assert.False(t, IsDisjoint(b, a))
}

func TestIsSubset(t *testing.T) {
	var a, b Set[int]

	assert.True(t, IsSubset(a, b))
	assert.True(t, IsSubset(b, a))

	a = New([]int{1, 2, 3})

	assert.False(t, IsSubset(a, b))
	assert.True(t, IsSubset(b, a))

	b = New([]int{1, 2, 3, 4, 5, 6})

	assert.True(t, IsSubset(a, b))
	assert.False(t, IsSubset(b, a))

	b = New([]int{2, 3, 4, 5, 6})

	assert.False(t, IsSubset(a, b))
	assert.False(t, IsSubset(b, a))
}

func TestIsSuperset(t *testing.T) {
	var b, a Set[int]

	assert.True(t, IsSuperset(a, b))
	assert.True(t, IsSuperset(b, a))

	a = New([]int{1, 2, 3})

	assert.True(t, IsSuperset(a, b))
	assert.False(t, IsSuperset(b, a))

	b = New([]int{1, 2, 3, 4, 5, 6})

	assert.True(t, IsSuperset(b, a))
	assert.False(t, IsSuperset(a, b))

	b = New([]int{1, 2, 3, 4, 5, 6, 7, 8})

	assert.False(t, IsSuperset(a, b))
	assert.True(t, IsSuperset(b, a))
}

func TestEqual(t *testing.T) {
	var a, b Set[int]

	assert.True(t, Equal(a, b))
	assert.True(t, Equal(b, a))

	a = New([]int{1, 2, 3})

	assert.False(t, Equal(a, b))
	assert.False(t, Equal(b, a))

	b = New([]int{1, 2, 3, 4})

	assert.False(t, Equal(a, b))
	assert.False(t, Equal(b, a))

	b = New([]int{1, 2, 3})

	assert.True(t, Equal(a, b))
	assert.True(t, Equal(b, a))
}
