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

func TestEdgeType(t *testing.T) {
	et := edgeType{
		"Your Moma", []string{"FailFrom1", "FailFrom2"},
		[]string{"FailTo1", "FailTo2"}}

	nt, _ := GetOrCreateNodeType("Your", "Moma")
	dn, _ := CreateDataNode(nt, []byte("Your Moma"))
	if et.ValidToNode(dn) {
		t.Error("Failed this data and edge node don't match")
	}
	if et.ValidFromNode(dn) {
		t.Error("Failed this data and edge node don't match")
	}

	et = edgeType{
		"Your Moma", []string{"Correct Edge"},
		[]string{"Correct Edge"}}
	nt, _ = GetOrCreateNodeType("Correct Edge", "Moma")
	dn, _ = CreateDataNode(nt, []byte("Your Moma"))
	if !et.ValidToNode(dn) {
		t.Error("Failed this data and edge node do match")
	}
	if !et.ValidFromNode(dn) {
		t.Error("Failed this data and edge node do match")
	}
}