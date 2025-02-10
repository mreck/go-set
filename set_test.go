package goset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDifference(t *testing.T) {
	a := Create([]int{1, 2, 3, 4, 5, 6})
	b := Create([]int{2, 4, 6})

	assert.Equal(t, Create([]int{1, 3, 5}), Difference(a, b))
}

func TestIntersection(t *testing.T) {
	a := Create([]int{1, 2, 3, 4, 5, 6})
	b := Create([]int{2, 4, 6, 8, 10})

	assert.Equal(t, Create([]int{2, 4, 6}), Intersection(a, b))
}

func TestSymmetricDifference(t *testing.T) {
	a := Create([]int{1, 2, 3, 4, 5, 6})
	b := Create([]int{4, 5, 6, 7, 8, 9})

	assert.Equal(t, Create([]int{1, 2, 3, 7, 8, 9}), SymmetricDifference(a, b))
}

func TestUnion(t *testing.T) {
	a := Create([]int{1, 2, 3, 4, 5, 6})
	b := Create([]int{4, 5, 6, 7, 8, 9})

	assert.Equal(t, Create([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}), Union(a, b))
}

func TestAreDisjoint(t *testing.T) {
	a := Create([]int{1, 2, 3})
	b := Create([]int{4, 5, 6})

	assert.True(t, AreDisjoint(a, b))

	a = Create([]int{1, 2, 3})
	b = Create([]int{3, 4, 5})

	assert.False(t, AreDisjoint(a, b))
}

func TestIsSubset(t *testing.T) {
	a := Create([]int{1, 2, 3})
	b := Create([]int{1, 2, 3, 4, 5, 6})

	assert.True(t, IsSubset(a, b))

	a = Create([]int{1, 2, 3})
	b = Create([]int{2, 3, 4, 5, 6})

	assert.False(t, IsSubset(a, b))
}

func TestIsSuperset(t *testing.T) {
	a := Create([]int{1, 2, 3, 4, 5, 6})
	b := Create([]int{1, 2, 3})

	assert.True(t, IsSuperset(a, b))

	a = Create([]int{2, 3, 4, 5, 6})
	b = Create([]int{1, 2, 3})

	assert.False(t, IsSuperset(a, b))
}
