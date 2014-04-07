package go_graph

import (
	"go_graph/data_types"
	"go_graph/node_edges"
	"sort"
	"testing"
)

func TestInit(t *testing.T) {
	nodeType1, _ := data_types.GetOrCreateNodeType("nodeType1", "nodeType1")
	node1, _ := node_edges.NewGraphNode(nodeType1, []byte("node1"))
	nodeType2, _ := data_types.GetOrCreateNodeType("nodeType2", "nodeType2")
	node_edges.CreateEdgeType("1-2", []data_types.NodeType{*nodeType1}, []data_types.NodeType{*nodeType2})
	node2, _ := node_edges.NewGraphNode(nodeType2, []byte("node2"))
	nodeType3, _ := data_types.GetOrCreateNodeType("nodeType3", "nodeType3")
	node_edges.CreateEdgeType("1-3", []data_types.NodeType{*nodeType1}, []data_types.NodeType{*nodeType3})
	node3, _ := node_edges.NewGraphNode(nodeType3, []byte("node3"))
	nodeType4, _ := data_types.GetOrCreateNodeType("nodeType4", "nodeType4")
	node_edges.CreateEdgeType("2-4", []data_types.NodeType{*nodeType2}, []data_types.NodeType{*nodeType4})
	node4, _ := node_edges.NewGraphNode(nodeType4, []byte("node4"))
	nodeType5, _ := data_types.GetOrCreateNodeType("nodeType5", "nodeType5")
	node_edges.CreateEdgeType("2-5", []data_types.NodeType{*nodeType2}, []data_types.NodeType{*nodeType5})
	node5, _ := node_edges.NewGraphNode(nodeType5, []byte("node5"))

	edge1, _ := node_edges.NewGraphEdge("1-2", &node1, &node2)
	edge2, _ := node_edges.NewGraphEdge("1-3", &node1, &node3)
	edge3, _ := node_edges.NewGraphEdge("2-4", &node2, &node4)
	edge4, _ := node_edges.NewGraphEdge("2-5", &node2, &node5)

	node1.AddToEdge(edge1)
	node1.AddToEdge(edge2)

	node2.AddFromEdge(edge1)
	node2.AddToEdge(edge3)
	node2.AddToEdge(edge4)

	node3.AddFromEdge(edge2)
	node4.AddFromEdge(edge3)
	node5.AddFromEdge(edge4)

	arrayPath := [][]string{[]string{"1-2"}, []string{"2-5"}}
	output := make(chan node_edges.GraphNode, 10)
	StartArrayStringWalkingPath(arrayPath, output, &node1)

	// Because string are easily sorted.
	nodeValue := func(local_node node_edges.GraphNode) []byte {
		tmp1 := local_node.Value()
		return tmp1.GetData()
	}
	expected := []string{string(nodeValue(node2)), string(nodeValue(node5))}
	actual := make([]string, 0)
stringLabel:
	for {
		select {
		case node, ok := <-output:
			if ok {
				actual = append(actual, string(nodeValue(node)))
			} else {
				break stringLabel
			}
		}
	}
	sort.Strings(expected)
	sort.Strings(actual)
	if len(expected) != len(actual) {
		t.Error("Different number of results from path search", len(expected), len(actual))
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Error("The Patch search is not returned the correct nodes")
		}
	}

	output = make(chan node_edges.GraphNode, 10)
	arrayPath = [][]string{[]string{"2"}, []string{"5"}}
	StartArrayRegExWalkingPath(arrayPath, output, &node1)
	actual = make([]string, 0)
regexLabel:
	for {
		select {
		case node, ok := <-output:
			if ok {
				actual = append(actual, string(nodeValue(node)))
			} else {
				break regexLabel
			}
		}
	}
	sort.Strings(expected)
	sort.Strings(actual)
	if len(expected) != len(actual) {
		t.Error("Different number of results from path search", len(expected), len(actual))
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Error("The Patch search is not returned the correct nodes")
		}
	}
}
