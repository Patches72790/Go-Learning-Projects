package main

import "testing"

func TestIntSet(t *testing.T) {
	var x, y IntSet

	x.Add(1)
	x.Add(2)
	x.Add(4)
	if x.Len() != 3 {
		t.Fatal("Expected length 3 for intset x")
	}

	y.Add(8)
	y.Add(16)
	y.Add(32)
	if y.Len() != 3 {
		t.Fatal("Expected length 3 for intset y")
	}

	x.UnionWith(&y)
	if x.Len() != 6 {
		t.Fatal("Expected length 6 for intset x")
	}

	if x.String() != "{1 2 4 8 16 32}" {
		t.Fatalf("Expected 6 items in string for intset x: %v", x.String())
	}
}

func TestRemove(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(2)

	x.Remove(2)
	if x.Has(2) || x.Len() != 1 {
		t.Fatalf("Expected to remove 2, but not: %v", x.String())
	}
}

func TestClear(t *testing.T) {
	var x IntSet
	x.AddAll(1, 2, 3)
	if x.Len() != 3 {
		t.Fatalf("Expected length 3: %v", x.String())
	}

	x.Clear()
	if x.Len() != 0 {
		t.Fatalf("Expected empty set: %v", x.String())
	}
}

func TestCopy(t *testing.T) {

	var x IntSet

	x.AddAll(1, 2, 4, 8, 16, 32)
	y := x.Copy()

	if x.String() != y.String() && x.Len() != y.Len() {
		t.Fatalf("Expected y to by copy of x: %v", y.String())
	}

	y.Remove(32)
	if !x.Has(32) {
		t.Fatalf("Expected x not to be affected by y copy removal: %v", x.String())
	}

}

func TestIntersect(t *testing.T) {
	var x, y IntSet

	x.AddAll(1, 2, 4, 8)
	y.AddAll(4, 8, 16, 32)

	x.IntersectWith(&y)

	if x.String() != "{4 8}" {
		t.Fatalf("Expected intersection of x to be {4 8}, but got: %v", x.String())
	}

	var a IntSet
	a.AddAll(1, 2, 3, 4, 5)
	b := a.Copy()

	a.IntersectWith(b)

	if a.String() != b.String() {
		t.Fatalf("Expected intersection to be equal for a and b: %v %v", a.String(), b.String())
	}
}
