// Package listops implements basic operations over lists.package ops
package listops

type binFunc func(x, y int) int

type predFunc func(x int) bool

type unaryFunc func(x int) int

// IntList represents a list of integers.
type IntList []int

// Append a second list of integers to the list.
func (l *IntList) Append(a IntList) IntList {
	return append(*l, a...)
}

// Concat returns a new list of integers with all received lists concatenated to it.
func (l *IntList) Concat(a []IntList) IntList {
	r := append(*l, []int{}...)
	for _, b := range a {
		r = append(r, b...)
	}
	return r
}

// Filter applies a predicate (or filter) to the list and returns a list with all the elements that matches it.
func (l *IntList) Filter(fn predFunc) IntList {
	r := []int{}
	for _, e := range *l {
		if fn(e) {
			r = append(r, e)
		}
	}
	return IntList(r)
}

// Foldl applies a binary function to elements of the list, starting from the initial value and passing the next element on the list.
func (l *IntList) Foldl(fn binFunc, initial int) int {
	r := initial
	for _, e := range *l {
		r = fn(r, e)
	}
	return r
}

// Foldr applies a binary function to elements of the list, starting from the initial value and passing the next value on the reverse order of the list.
func (l *IntList) Foldr(fn binFunc, initial int) int {
	r := initial
	a := make([]int, len(*l))
	copy(a, *l)
	for i := len(a) - 1; i >= 0; i-- {
		r = fn(a[i], r)
	}
	return r
}

// Length returns the number of elements of the list.
func (l *IntList) Length() int {
	return len(*l)
}

// Map applies a unary function to every element of the list and returns a new list with the result.
func (l *IntList) Map(fn unaryFunc) IntList {
	r := []int{}
	for _, e := range *l {
		r = append(r, fn(e))
	}
	return IntList(r)
}

// Reverse the order of the list and returns a new list with the result.
func (l *IntList) Reverse() IntList {
	r := make(IntList, len(*l))
	for i, e := range *l {
		r[len(r)-i-1] = e
	}
	return IntList(r)
}
