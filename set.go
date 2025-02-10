package goset

import (
	"fmt"
	"strings"
)

// Set holds the values for a set of type T
type Set[T comparable] struct {
	values map[T]struct{}
}

// String returns the string representation of the set
func (set Set[T]) String() string {
	var b strings.Builder
	first := true
	b.WriteString("{ ")
	for value := range set.values {
		if !first {
			b.WriteString(", ")
		} else {
			first = false
		}
		b.WriteString(fmt.Sprintf("%v", value))
	}
	b.WriteString(" }")
	return b.String()
}

// Add adds a value
func (set *Set[T]) Add(value T) {
	if set.values == nil {
		set.values = map[T]struct{}{}
	}
	set.values[value] = struct{}{}
}

// AddMany adds multiple values
func (set *Set[T]) AddMany(values []T) {
	if values == nil {
		return
	}
	if set.values == nil {
		set.values = map[T]struct{}{}
	}
	for _, value := range values {
		set.values[value] = struct{}{}
	}
}

// AddMany adds multiple values
func (set *Set[T]) AddSet(otherSet Set[T]) {
	if otherSet.values == nil {
		return
	}
	if set.values == nil {
		set.values = map[T]struct{}{}
	}
	for value := range otherSet.values {
		set.values[value] = struct{}{}
	}
}

// Remove removes the value from the set
func (set *Set[T]) Remove(value T) {
	if set.values == nil {
		return
	}
	delete(set.values, value)
}

// RemoveMany removes the values from the set
func (set *Set[T]) RemoveMany(values []T) {
	if set.values == nil {
		return
	}
	for _, value := range values {
		delete(set.values, value)
	}
}

// RemoveSet removes the values from the set
func (set *Set[T]) RemoveSet(otherSet Set[T]) {
	if set.values == nil {
		return
	}
	for value := range otherSet.values {
		delete(set.values, value)
	}
}

// Clear removes all values from the set
func (set *Set[T]) Clear() {
	set.values = nil
}

// Len returns the number of values in the set
func (set Set[T]) Len() int {
	return len(set.values)
}

// Contains returns true if the set contains the value
func (set Set[T]) Contains(value T) bool {
	if set.values == nil {
		return false
	}
	_, ok := set.values[value]
	return ok
}

// List returns a list of the Set's values
func (set Set[T]) List() []T {
	if set.values == nil {
		return nil
	}
	var list []T
	for value := range set.values {
		list = append(list, value)
	}
	return list
}

// Create returns a new Set based on the list
func Create[T comparable](list []T) Set[T] {
	var set Set[T]
	if list != nil {
		set.AddMany(list)
	}
	return set
}

// Difference returns: A \ B
func Difference[T comparable](a Set[T], b Set[T]) Set[T] {
	var result Set[T]
	for value := range a.values {
		if !b.Contains(value) {
			result.Add(value)
		}
	}
	return result
}

// Intersection returns: A ∩ B
func Intersection[T comparable](a Set[T], b Set[T]) Set[T] {
	var result Set[T]
	for value := range a.values {
		if b.Contains(value) {
			result.Add(value)
		}
	}
	return result
}

// SymmetricDifference returns: (A \ B) ∪ (B \ A)
func SymmetricDifference[T comparable](a Set[T], b Set[T]) Set[T] {
	return Union(Difference(a, b), Difference(b, a))
}

// Union returns: A ∪ B
func Union[T comparable](a Set[T], b Set[T]) Set[T] {
	var result Set[T]
	result.AddSet(a)
	result.AddSet(b)
	return result
}

// AreDisjoint returns: A ∩ B = 0
func AreDisjoint[T comparable](a Set[T], b Set[T]) bool {
	for value := range a.values {
		if b.Contains(value) {
			return false
		}
	}
	for value := range b.values {
		if a.Contains(value) {
			return false
		}
	}
	return true
}

// IsSubset returns: A ⊆ B
func IsSubset[T comparable](a Set[T], b Set[T]) bool {
	for value := range a.values {
		if !b.Contains(value) {
			return false
		}
	}
	return true
}

// IsSuperset returns: A ⊇ B
func IsSuperset[T comparable](a Set[T], b Set[T]) bool {
	return IsSubset(b, a)
}

// Identical returns: (A ⊆ B) & (A ⊇ B)
func Identical[T comparable](a Set[T], b Set[T]) bool {
	if a.Len() != b.Len() {
		return false
	}
	for value := range a.values {
		if !b.Contains(value) {
			return false
		}
	}
	return true
}
