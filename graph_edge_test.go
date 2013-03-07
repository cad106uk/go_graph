package go_graph

import "testing"

func TestNew(t *testing.T) {
	// Clean the data
	all_edge_types = make(map[string]edgeType)

	fromEdge, _ := GetOrCreateNodeType("ValidFrom", "Moma")
	fromData, _ := CreateDataNode(fromEdge, []byte("Your Moma"))
	toEdge, _ := GetOrCreateNodeType("ValidTo", "Moma")
	toData, _ := CreateDataNode(toEdge, []byte("Your Moma"))
	failNode := dataNode{}
	et, _ := CreateEdgeType("Your Moma", []*nodeType{fromEdge}, []*nodeType{toEdge})

	ge, err := NewGraphEdge("FAIL", &fromData, &toData)
	if err == nil {
		t.Error("This edge was not made. This should fail")
	}
	ge, err = NewGraphEdge("Your Moma", &fromData, &toData)
	if err != nil {
		t.Error(err) // This should fail
	}

	if ge.edgeType.edgeTypeName != et.edgeTypeName {
		t.Error("edgeType problem")
	}
	if ge.connectFrom.id != fromData.id {
		t.Error("The from data shold match")
	}
	if ge.connectTo.id != toData.id {
		t.Error("The To data should match")
	}

	ge, err = NewGraphEdge("Your Moma", &failNode, &toData)
	if err == nil {
		t.Error("Bad from node should have stopped this")
	}

	ge, err = NewGraphEdge("Your Moma", &fromData, &failNode)
	if err == nil {
		t.Error("Bad to node should have stopped this")
	}

	// Clean the data
	all_edge_types = make(map[string]edgeType)
}