package go_graph

import	"testing"

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

func TestGetEdgeType(t *testing.T) {
	_, err := GetEdgeType("Your Skinny Moma")
	if err == nil {
		t.Error("Your moma is not skinny", err)
	}
}

func TestCreateEdgeType(t *testing.T) {
	et, err := CreateEdgeType(
		"Your Moma", []string{"ValidFrom1", "ValidFrom2"},
		[]string{"ValidTo1", "ValidTo2"})
	if err != nil {
		t.Error(err)
	}

	_, err = CreateEdgeType(
		"Your Moma", []string{"ValidFrom1", "ValidFrom2"},
		[]string{"ValidTo1", "ValidTo2"})
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