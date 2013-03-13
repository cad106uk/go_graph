package go_graph

import "testing"

func TestBasicGraphNode(t *testing.T) {
	/*
	 data is kept in dataNode
	 dataNode has a nodeType
	 nodeType has a name and a description

	 GraphNode 1 and only 1 value which is a dataNode
	 GraphNode has 2 lists of GraphEdges 1 for all the GraphEdges that connect from this GraphNode and 1 for all the GraphEdges that connect to this GraphNode

	 GraphEdge has an edgeType and connects from 1 GraphNode to another GraphNode

	 A GraphEdge can only connect 2 GraphNodes together. A GraphNode can have many GraphEdges.

	 Also the GraphNode and GraphEdge are public and for the rest of the system to use. The dataNode and nodeType are private internal structs
	*/

	// nt, _ := GetOrCreateNodeType("Your", "Moma")
	// ge, _ := NewGraphEdge("Your Moma", *nt, *nt)
	// gn := GraphNode{}
	// err := gn.Init(nt, []byte("Your Moma"), ge, ge)
	// if err != nil {
	// 	t.Error("Couldn't init a graph node", err)
	// }

	// val := gn.Value()
	// if gn.value.data.id != val.id {
	// 	t.Error("GraphNode Value does not match value.data")
	// }
}
