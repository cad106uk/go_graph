package go_graph

import "testing"

/*
Create an Edge Type, so we can have edges of this type.
Then create an actual edge of this type.
So we know the types of relations before we create them.
*/
func TestNew(t *testing.T) {
	// Clean the data
	allEdgeTypes = make(map[string]edgeType)

	fromEdge, _ := GetOrCreateNodeType("ValidFrom", "Moma")
	fromData, _ := CreateDataNode(fromEdge, []byte("Your Moma"))
	toEdge, _ := GetOrCreateNodeType("ValidTo", "Moma")
	toData, _ := CreateDataNode(toEdge, []byte("Your Moma"))
	failNode := GraphNode{}
	et, _ := CreateEdgeType("Your Moma", []*nodeType{fromEdge}, []*nodeType{toEdge})
	fromGN := GraphNode{fromData, make([]GraphEdge, 0), make([]GraphEdge, 0)}
	toGN := GraphNode{toData, make([]GraphEdge, 0), make([]GraphEdge, 0)}

	_, err := NewGraphEdge("FAIL", &fromGN, &toGN)
	if err == nil {
		t.Error("This edge was not made. This should fail")
	}
	ge, err := NewGraphEdge("Your Moma", &fromGN, &toGN)
	if err != nil {
		t.Error(err) // This should fail
	}

	if ge.EdgeType.edgeTypeName != et.edgeTypeName {
		t.Error("edgeType problem")
	}
	if ge.ConnectFrom.value.id != fromData.id {
		t.Error("The from data shold match")
	}
	if ge.ConnectTo.value.id != toData.id {
		t.Error("The To data should match")
	}

	_, err = NewGraphEdge("Your Moma", &failNode, &toGN)
	if err == nil {
		t.Error("Bad from node should have stopped this")
	}

	_, err = NewGraphEdge("Your Moma", &fromGN, &failNode)
	if err == nil {
		t.Error("Bad to node should have stopped this")
	}

	// Clean the data
	allEdgeTypes = make(map[string]edgeType)
}
