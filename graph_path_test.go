package go_graph

import (

	"testing"
)

func generateTestNodes1() []GraphNode {
	nodeType1, _ := GetOrCreateNodeType("nodeType1", "description of nodeType1")
	node1, _ := NewGraphNode(nodeType1, []byte("node1"))

	nodeType2, _ := GetOrCreateNodeType("nodeType2", "description of nodeType2")
	CreateEdgeType("1-2", []NodeType{*nodeType1}, []NodeType{*nodeType2})
	node2, _ := NewGraphNode(nodeType2, []byte("node2"))

	nodeType3, _ := GetOrCreateNodeType("nodeType3", "description of nodeType3")
	CreateEdgeType("1-3", []NodeType{*nodeType1}, []NodeType{*nodeType3})
	node3, _ := NewGraphNode(nodeType3, []byte("node3"))

	nodeType4, _ := GetOrCreateNodeType("nodeType4", "description of nodeType4")
	CreateEdgeType("2-4", []NodeType{*nodeType2}, []NodeType{*nodeType4})
	node4, _ := NewGraphNode(nodeType4, []byte("node4"))

	nodeType5, _ := GetOrCreateNodeType("nodeType5", "description of nodeType5")
	CreateEdgeType("2-5", []NodeType{*nodeType2}, []NodeType{*nodeType5})
	node5, _ := NewGraphNode(nodeType5, []byte("node5"))

	edge1, _ := NewGraphEdge("1-2", node1, node2)
	edge2, _ := NewGraphEdge("1-3", node1, node3)
	edge3, _ := NewGraphEdge("2-4", node2, node4)
	edge4, _ := NewGraphEdge("2-5", node2, node5)

	node1.AddToEdge(edge1)
	node1.AddToEdge(edge2)

	node2.AddFromEdge(edge1)
	node2.AddToEdge(edge3)
	node2.AddToEdge(edge4)

	node3.AddFromEdge(edge2)
	node4.AddFromEdge(edge3)
	node5.AddFromEdge(edge4)

	return []GraphNode{*node1, *node2, *node3, *node4, *node5}
}

func nodeValue(local_node GraphNode) []byte {
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
	output := make(chan GraphNode, 10)
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

	output = make(chan GraphNode, 10)
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

func generateTestNodes2() []GraphNode {
	nodeType1, _ := GetOrCreateNodeType("nodeType1", "description of nodeType1")
	node1, _ := NewGraphNode(nodeType1, []byte("data set 1"))
	node2, _ := NewGraphNode(nodeType1, []byte("Your Moma"))
	nodeType2, _ := GetOrCreateNodeType("nodeType2", "description of nodeType2")
	node3, _ := NewGraphNode(nodeType2, []byte("middle bit"))

	CreateEdgeType("1-2", []NodeType{*nodeType1}, []NodeType{*nodeType2})
	CreateEdgeType("2-1", []NodeType{*nodeType2}, []NodeType{*nodeType1})

	edge1, _ := NewGraphEdge("1-2", node1, node3)
	edge2, _ := NewGraphEdge("2-1", node3, node2)

	node1.AddToEdge(edge1)
	node3.AddToEdge(edge2)

	return []GraphNode{*node1}
}

func TestData(t *testing.T) {
	expected := "Your Moma"

	nodes := generateTestNodes2()

	output := make(chan GraphNode, 10)
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
