/*
 data is kept in dataNode
 dataNode has a nodeType
 nodeType has a name and a description

 GraphNode has 1 and only 1 value which is a dataNode
 GraphNode has 2 lists of GraphEdges 1 for all the GraphEdges that connect from this GraphNode and 1 for all the GraphEdges that connect to this GraphNode

 GraphEdge has an edgeType and connects from 1 GraphNode to another GraphNode

 A GraphEdge can only connect 2 GraphNodes together. A GraphNode can have many GraphEdges.

 Also the GraphNode and GraphEdge are public and for the rest of the system to use. The dataNode and nodeType are private internal structs
*/
package go_graph

import (
	"testing"
)

/*
Create an Edge Type, so we can have edges of this type.
Then create an actual edge of this type.
So we know the types of relations before we create them.
*/
func TestNew(t *testing.T) {
	// Clean the data
	allEdgeTypes.m = make(map[string]edgeType)

	fromEdge, _ := GetOrCreateNodeType("ValidFrom", "Moma")
	fromData, _ := CreateDataNode(fromEdge, []byte("Your Moma"))
	toEdge, _ := GetOrCreateNodeType("ValidTo", "Moma")
	toData, _ := CreateDataNode(toEdge, []byte("Your Moma"))
	failNode := GraphNode{}
	et, _ := CreateEdgeType("Your Moma", []NodeType{*fromEdge}, []NodeType{*toEdge})
	fromGN := GraphNode{*fromData, make([]GraphEdge, 0), make([]GraphEdge, 0)}
	toGN := GraphNode{*toData, make([]GraphEdge, 0), make([]GraphEdge, 0)}

	//No edge type called GAIL
	_, err := NewGraphEdge("FAIL", &fromGN, &toGN)
	if err == nil {
		t.Error("This edge was not made. This should fail")
	}
	//Have created an edge type called "Your Moma" with the CreateEdgeType function
	ge, err := NewGraphEdge("Your Moma", &fromGN, &toGN)
	if err != nil {
		t.Error(err) // This should fail
	}

	if ge.EdgeType.edgeTypeName != et.edgeTypeName {
		t.Error("edgeType problem")
	}
	if ge.ConnectFrom.value.GetId() != fromData.GetId() {
		t.Error("The from data shold match")
	}
	if ge.ConnectTo.value.GetId() != toData.GetId() {
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
	allEdgeTypes.m = make(map[string]edgeType)
}
