package go_graph

import "testing"
import "sort"

func TestInit(t *testing.T) {
	nodeType1, _ := GetOrCreateNodeType("nodeType1", "nodeType1")
	node1, _ := NewGraphNode(nodeType1, []byte("node1"))
	nodeType2, _ := GetOrCreateNodeType("nodeType2", "nodeType2")
	CreateEdgeType("1-2", []nodeType{*nodeType1}, []nodeType{*nodeType2})
	node2, _ := NewGraphNode(nodeType2, []byte("node2"))
	nodeType3, _ := GetOrCreateNodeType("nodeType3", "nodeType3")
	CreateEdgeType("1-3", []nodeType{*nodeType1}, []nodeType{*nodeType3})
	node3, _ := NewGraphNode(nodeType3, []byte("node3"))
	nodeType4, _ := GetOrCreateNodeType("nodeType4", "nodeType4")
	CreateEdgeType("2-4", []nodeType{*nodeType2}, []nodeType{*nodeType4})
	node4, _ := NewGraphNode(nodeType4, []byte("node4"))
	nodeType5, _ := GetOrCreateNodeType("nodeType5", "nodeType5")
	CreateEdgeType("2-5", []nodeType{*nodeType2}, []nodeType{*nodeType5})
	node5, _ := NewGraphNode(nodeType5, []byte("node5"))

	edge1, _ := NewGraphEdge("1-2", &node1, &node2)
	edge2, _ := NewGraphEdge("1-3", &node1, &node3)
	edge3, _ := NewGraphEdge("2-4", &node2, &node4)
	edge4, _ := NewGraphEdge("2-5", &node2, &node5)

	node1.AddToEdge(edge1)
	node1.AddToEdge(edge2)

	node2.AddFromEdge(edge1)
	node2.AddToEdge(edge3)
	node2.AddToEdge(edge4)

	node3.AddFromEdge(edge2)
	node4.AddFromEdge(edge3)
	node5.AddFromEdge(edge4)

	arrayPath := [][]string{[]string{"1-2"}, []string{"2-5"}}
	output := make(chan GraphNode, 10)
	StartArrayStringWalkingPath(arrayPath, output, &node1)

	// Because string are easily sorted.
	expected := []string{string(node2.value.data.data), string(node5.value.data.data)}
	actual := make([]string, 0)
stringLabel:
	for {
		select {
		case node, ok := <-output:
			if ok {
				actual = append(actual, string(node.value.data.data))
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

	output = make(chan GraphNode, 10)
	arrayPath = [][]string{[]string{"2"}, []string{"5"}}
	StartArrayRegExWalkingPath(arrayPath, output, &node1)
	actual = make([]string, 0)
regexLabel:
	for {
		select {
		case node, ok := <-output:
			if ok {
				actual = append(actual, string(node.value.data.data))
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
