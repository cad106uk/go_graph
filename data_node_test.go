package go_graph

import (
	"bytes"
	"testing"
)

var nt, _ = GetOrCreateNodeType("Your", "Moma")

func TestCreateDataNode(t *testing.T) {
	node, err := CreateDataNode(*nt, []byte("Your Moma"))
	if err != nil {
		t.Error("Failed to create a data node ", err)
	}
	if node.dataType != *nt {
		t.Error("The node was not created with the correct type")
	}

	t1 := node.GetType()
	if t1 != node.dataType {
		t.Error("GetType fail to return the correct dataType ", t1, node.dataType)
	}

	t2 := node.GetValue()
	if !bytes.Equal(t2.data, node.data.data) {
		t.Error("GetValue fail to return the correct data value", t2, node.data)
	}

	t3 := node.GetId()
	if t3 != node.id {
		t.Error("GetId fail to return the correct id ", t3, node.id)
	}

	// Yes I know this does not test randomness properly, but it
	// should catch me at my most stupid
	for i := 0; i < 100; i += 1 {
		temp, _ := CreateDataNode(*nt, []byte("Your Moma"))
		if node.id == temp.id {
			t.Error("MATCHING IDS! ", node.id, temp.id)
		}
	}

}
