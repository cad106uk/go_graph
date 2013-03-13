package go_graph

import "testing"

func TestMatchEdgeType(t *testing.T) {
	find, _ := GetOrCreateNodeType("find", "Your")
	invalid, _ := GetOrCreateNodeType("invalid", "Moma")
	junk, _ := GetOrCreateNodeType("junk", "junk")
	options := []*nodeType{find, junk}
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
	failFromArgs := []*nodeType{&nodeType{"FailFrom1", "Your Moma"}, &nodeType{"FailFrom2", "Your Moma"}}
	failToArgs := []*nodeType{&nodeType{"FailTo1", "Your Moma"}, &nodeType{"FailTo2", "Your Moma"}}
	et := edgeType{"Your Moma", failFromArgs, failToArgs}

	nt, _ := GetOrCreateNodeType("Your", "Moma")
	dn, _ := CreateDataNode(nt, []byte("Your Moma"))
	gn := GraphNode{}
	gn.SetValue(dn.data)
	if et.ValidToNode(gn) {
		t.Error("Failed this data and edge node don't match")
	}
	if et.ValidFromNode(gn) {
		t.Error("Failed this data and edge node don't match")
	}

	nt, _ = GetOrCreateNodeType("Correct Edge", "Moma")
	et = edgeType{"Your Moma", []*nodeType{nt}, []*nodeType{nt}}
	dn, _ = CreateDataNode(nt, []byte("Your Moma"))
	if !et.ValidToNode(gn) {
		t.Error("Failed this data and edge node do match")
	}
	if !et.ValidFromNode(gn) {
		t.Error("Failed this data and edge node do match")
	}
}

func TestGetEdgeType(t *testing.T) {
	_, err := GetEdgeType("Your Skinny Moma")
	if err == nil {
		t.Error("Your moma is not skinny", err)
	}
}

func TestCreateEdgeType(t *testing.T) {
	validFromArgs := []*nodeType{&nodeType{"ValidFrom1", "Your Moma"}, &nodeType{"ValidFrom2", "Your Moma"}}
	validToArgs := []*nodeType{&nodeType{"ValidTo1", "Your Moma"}, &nodeType{"ValidTo2", "Your Moma"}}
	et, err := CreateEdgeType("Your Moma", validFromArgs, validToArgs)
	if err != nil {
		t.Error(err)
	}

	_, err = CreateEdgeType("Your Moma", validFromArgs, validToArgs)
	if err == nil {
		t.Error("2 Momas. No wait, times have changes now. Ah what the hell, I'm cracking Your Moma jokes all over my testing, what do you expect")
	}

	test, err := GetEdgeType("Your Moma")
	if err != nil {
		t.Error("This call should get Your Moma")
	}

	if et.edgeTypeName != test.edgeTypeName {
		t.Error("These are supposed to the same Moma")
	}
}
