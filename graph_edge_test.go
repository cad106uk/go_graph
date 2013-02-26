package go_graph

import "testing"

func TestMatchEdgeType(t *testing.T) {
	find := "find"
	invalid := "invalid"
	options := []string{"find", "Your", "Mother"}
	t1 := matchEdgeType(find, options)
	if !t1 {
		t.Error("Can't find find!")
	}

	t2 := matchEdgeType(invalid, options)
	if t2 {
		t.Error("Found invalid string!")
	}
}