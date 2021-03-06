package stringset_test

import (
	"testing"

	"github.com/jw4/x/stringset"
)

func TestOrderedStringSet(t *testing.T) {
	t.Parallel()

	var coll stringset.OrderedStringSet

	ok := coll.Add("one", "two", "three", "four")
	if !ok {
		t.Fatal("could not add strings")
	}

	ok = coll.Remove("one", "four", "two")
	if !ok {
		t.Fatal("could not remove strings")
	}

	v := coll.Values()
	if len(v) != 1 {
		t.Fatal("incorrect number of Values()")
	}

	if v[0] != "three" {
		t.Fatalf("wrong values removed: %v", v)
	}
}
