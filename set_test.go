package goset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDifference(t *testing.T) {
	var a, b Set[int]

	assert.Equal(t, Create([]int(nil)), Difference(a, b))
	assert.Equal(t, Create([]int(nil)), Difference(b, a))

	a = Create([]int{1, 2, 3, 4, 5, 6})

	assert.Equal(t, Create([]int{1, 2, 3, 4, 5, 6}), Difference(a, b))
	assert.Equal(t, Create([]int(nil)), Difference(b, a))

	b = Create([]int{2, 4, 6})

	assert.Equal(t, Create([]int{1, 3, 5}), Difference(a, b))
	assert.Equal(t, Create([]int(nil)), Difference(b, a))
}

func TestIntersection(t *testing.T) {
	var a, b Set[int]

	assert.Equal(t, Create([]int(nil)), Intersection(a, b))
	assert.Equal(t, Create([]int(nil)), Intersection(b, a))

	a = Create([]int{1, 2, 3, 4, 5, 6})

	assert.Equal(t, Create([]int(nil)), Intersection(a, b))
	assert.Equal(t, Create([]int(nil)), Intersection(b, a))

	b = Create([]int{2, 4, 6, 8, 10})

	assert.Equal(t, Create([]int{2, 4, 6}), Intersection(a, b))
	assert.Equal(t, Create([]int{2, 4, 6}), Intersection(b, a))
}

func TestSymmetricDifference(t *testing.T) {
	var a, b Set[int]

	assert.Equal(t, Create([]int(nil)), SymmetricDifference(a, b))
	assert.Equal(t, Create([]int(nil)), SymmetricDifference(b, a))

	a = Create([]int{1, 2, 3, 4, 5, 6})
	assert.Equal(t, Create([]int{1, 2, 3, 4, 5, 6}), SymmetricDifference(a, b))
	assert.Equal(t, Create([]int{1, 2, 3, 4, 5, 6}), SymmetricDifference(b, a))

	b = Create([]int{4, 5, 6, 7, 8, 9})

	assert.Equal(t, Create([]int{1, 2, 3, 7, 8, 9}), SymmetricDifference(a, b))
	assert.Equal(t, Create([]int{1, 2, 3, 7, 8, 9}), SymmetricDifference(b, a))
}

func TestUnion(t *testing.T) {
	var a, b Set[int]

	assert.Equal(t, Create([]int(nil)), Union(a, b))
	assert.Equal(t, Create([]int(nil)), Union(b, a))

	a = Create([]int{1, 2, 3, 4, 5, 6})

	assert.Equal(t, Create([]int{1, 2, 3, 4, 5, 6}), Union(a, b))
	assert.Equal(t, Create([]int{1, 2, 3, 4, 5, 6}), Union(b, a))

	b = Create([]int{4, 5, 6, 7, 8, 9})

	assert.Equal(t, Create([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}), Union(a, b))
	assert.Equal(t, Create([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}), Union(b, a))
}

func TestAreDisjoint(t *testing.T) {
	var a, b Set[int]

	assert.True(t, AreDisjoint(a, b))
	assert.True(t, AreDisjoint(b, a))

	a = Create([]int{1, 2, 3})

	assert.True(t, AreDisjoint(a, b))
	assert.True(t, AreDisjoint(b, a))

	b = Create([]int{4, 5, 6})

	assert.True(t, AreDisjoint(a, b))
	assert.True(t, AreDisjoint(b, a))

	b = Create([]int{3, 4, 5})

	assert.False(t, AreDisjoint(a, b))
	assert.False(t, AreDisjoint(b, a))
}

func TestIsSubset(t *testing.T) {
	var a, b Set[int]

	assert.True(t, IsSubset(a, b))
	assert.True(t, IsSubset(b, a))

	a = Create([]int{1, 2, 3})

	assert.False(t, IsSubset(a, b))
	assert.True(t, IsSubset(b, a))

	b = Create([]int{1, 2, 3, 4, 5, 6})

	assert.True(t, IsSubset(a, b))
	assert.False(t, IsSubset(b, a))

	b = Create([]int{2, 3, 4, 5, 6})

	assert.False(t, IsSubset(a, b))
	assert.False(t, IsSubset(b, a))
}

func TestIsSuperset(t *testing.T) {
	var b, a Set[int]

	assert.True(t, IsSuperset(a, b))
	assert.True(t, IsSuperset(b, a))

	a = Create([]int{1, 2, 3})

	assert.True(t, IsSuperset(a, b))
	assert.False(t, IsSuperset(b, a))

	b = Create([]int{1, 2, 3, 4, 5, 6})

	assert.True(t, IsSuperset(b, a))
	assert.False(t, IsSuperset(a, b))

	b = Create([]int{1, 2, 3, 4, 5, 6, 7, 8})

	assert.False(t, IsSuperset(a, b))
	assert.True(t, IsSuperset(b, a))
}
