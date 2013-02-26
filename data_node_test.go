package go_graph

import "testing"

var nt, _ = GetOrCreateNodeType("Your", "Moma")

func TestCreateDataNode(t *testing.T) {
	node, err := CreateDataNode(nt, "Your Moma")
	if err != nil {
		t.Error("Failed to create a data node ", err)
	}
	if node.dataType != nt {
		t.Error("The node was not created with the correct type")
	}
} 

