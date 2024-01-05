package goset

import "testing"

func TestDifference(t *testing.T) {
	a := Create([]int{1, 2, 3, 4, 5, 6})
	b := Create([]int{2, 4, 6})
	r := Difference(a, b)
	if !Identical(r, Create([]int{1, 3, 5})) {
		t.Error("There's a bug in Difference()")
	}
}

func TestIntersection(t *testing.T) {
	a := Create([]int{1, 2, 3, 4, 5, 6})
	b := Create([]int{2, 4, 6, 8, 10})
	r := Intersection(a, b)
	if !Identical(r, Create([]int{2, 4, 6})) {
		t.Error("There's a bug in Intersection()")
	}
}

func TestSymmetricDifference(t *testing.T) {
	a := Create([]int{1, 2, 3, 4, 5, 6})
	b := Create([]int{4, 5, 6, 7, 8, 9})
	r := SymmetricDifference(a, b)
	if !Identical(r, Create([]int{1, 2, 3, 7, 8, 9})) {
		t.Error("There's a bug in SymmetricDifference()")
	}
}

func TestUnion(t *testing.T) {
	a := Create([]int{1, 2, 3, 4, 5, 6})
	b := Create([]int{4, 5, 6, 7, 8, 9})
	r := Union(a, b)
	if !Identical(r, Create([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})) {
		t.Error("There's a bug in Union()")
	}
}

func TestAreDisjoint(t *testing.T) {
	a := Create([]int{1, 2, 3})
	b := Create([]int{4, 5, 6})
	r := AreDisjoint(a, b)
	if !r {
		t.Error("There's a bug in AreDisjoint()")
	}

	a = Create([]int{1, 2, 3})
	b = Create([]int{3, 4, 5})
	r = AreDisjoint(a, b)
	if r {
		t.Error("There's a bug in AreDisjoint()")
	}
}

func TestIsSubset(t *testing.T) {
	a := Create([]int{1, 2, 3})
	b := Create([]int{1, 2, 3, 4, 5, 6})
	r := IsSubset(a, b)
	if !r {
		t.Error("There's a bug in IsSubset()")
	}

	a = Create([]int{1, 2, 3})
	b = Create([]int{2, 3, 4, 5, 6})
	r = IsSubset(a, b)
	if r {
		t.Error("There's a bug in IsSubset()")
	}
}

func TestIsSuperset(t *testing.T) {
	a := Create([]int{1, 2, 3, 4, 5, 6})
	b := Create([]int{1, 2, 3})
	r := IsSuperset(a, b)
	if !r {
		t.Error("There's a bug in IsSuperset()")
	}

	a = Create([]int{2, 3, 4, 5, 6})
	b = Create([]int{1, 2, 3})
	r = IsSuperset(a, b)
	if r {
		t.Error("There's a bug in IsSuperset()")
	}
}
