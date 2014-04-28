package go_graph

import (
	"github.com/cad106uk/go_graph/data_types"
	"github.com/cad106uk/go_graph/node_edges"
	"testing"
)

func generateTestNodes1() []node_edges.GraphNode {
	nodeType1, _ := data_types.GetOrCreateNodeType("nodeType1", "description of nodeType1")
	node1, _ := node_edges.NewGraphNode(nodeType1, []byte("node1"))

	nodeType2, _ := data_types.GetOrCreateNodeType("nodeType2", "description of nodeType2")
	node_edges.CreateEdgeType("1-2", []data_types.NodeType{*nodeType1}, []data_types.NodeType{*nodeType2})
	node2, _ := node_edges.NewGraphNode(nodeType2, []byte("node2"))

	nodeType3, _ := data_types.GetOrCreateNodeType("nodeType3", "description of nodeType3")
	node_edges.CreateEdgeType("1-3", []data_types.NodeType{*nodeType1}, []data_types.NodeType{*nodeType3})
	node3, _ := node_edges.NewGraphNode(nodeType3, []byte("node3"))

	nodeType4, _ := data_types.GetOrCreateNodeType("nodeType4", "description of nodeType4")
	node_edges.CreateEdgeType("2-4", []data_types.NodeType{*nodeType2}, []data_types.NodeType{*nodeType4})
	node4, _ := node_edges.NewGraphNode(nodeType4, []byte("node4"))

	nodeType5, _ := data_types.GetOrCreateNodeType("nodeType5", "description of nodeType5")
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

	return []node_edges.GraphNode{node1, node2, node3, node4, node5}
}

func nodeValue(local_node node_edges.GraphNode) []byte {
	tmp1 := local_node.Value()
	return tmp1.GetData()
}

func TestInit(t *testing.T) {
	nodes := generateTestNodes1()

	// Because string are easily sorted.
	expected1 := []string{"node2", "node5"}
	expected2 := []string{string(nodeValue(nodes[1])), string(nodeValue(nodes[4]))}
	actual := make([]string, 0)

	arrayPath := [][]string{[]string{"1-2"}, []string{"2-5"}}
	output := make(chan node_edges.GraphNode, 10)
	StartArrayStringWalkingPath(arrayPath, output, &nodes[0])

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

	if len(expected1) != len(actual) {
		t.Error("Different number of results from path search", len(expected1), len(actual))
	}
	for i := 0; i < len(expected1); i++ {
		if expected1[i] != actual[i] {
			t.Error("The Patch search is not returned the correct nodes", expected1[i], actual[i])
		}
		if expected2[i] != actual[i] {
			t.Error("The Patch search is not returned the correct nodes", expected2[i], actual[i])
		}
	}

	output = make(chan node_edges.GraphNode, 10)
	arrayPath = [][]string{[]string{"2"}, []string{"5"}}
	StartArrayRegExWalkingPath(arrayPath, output, &nodes[0])
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

	if len(expected1) != len(actual) {
		t.Error("Different number of results from path search", len(expected1), len(actual))
	}
	for i := 0; i < len(expected1); i++ {
		if expected1[i] != actual[i] {
			t.Error("The Patch search is not returned the correct nodes", expected1[i], actual[i])
		}
		if expected2[i] != actual[i] {
			t.Error("The Patch search is not returned the correct nodes", expected2[i], actual[i])
		}
	}
}

func generateTestNodes2() []node_edges.GraphNode {
	nodeType1, _ := data_types.GetOrCreateNodeType("nodeType1", "description of nodeType1")
	node1, _ := node_edges.NewGraphNode(nodeType1, []byte("data set 1"))
	node2, _ := node_edges.NewGraphNode(nodeType1, []byte("Your Moma"))
	nodeType2, _ := data_types.GetOrCreateNodeType("nodeType2", "description of nodeType2")
	node3, _ := node_edges.NewGraphNode(nodeType2, []byte("middle bit"))

	node_edges.CreateEdgeType("1-2", []data_types.NodeType{*nodeType1}, []data_types.NodeType{*nodeType2})
	node_edges.CreateEdgeType("2-1", []data_types.NodeType{*nodeType2}, []data_types.NodeType{*nodeType1})

	edge1, _ := node_edges.NewGraphEdge("1-2", &node1, &node3)
	edge2, _ := node_edges.NewGraphEdge("2-1", &node3, &node2)

	node1.AddToEdge(edge1)
	node3.AddToEdge(edge2)

	return []node_edges.GraphNode{node1}
}

func TestData(t *testing.T) {
	expected := "Your Moma"

	nodes := generateTestNodes2()

	output := make(chan node_edges.GraphNode, 10)
	arrayPath := [][]string{[]string{"1"}, []string{"1"}}
	StartArrayRegExWalkingPath(arrayPath, output, &nodes[0])
	actual := make([]string, 0)

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

	res := actual[len(actual)-1]
	if res != expected {
		t.Error("result not expected", res, expected)
	}
}
