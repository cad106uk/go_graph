package go_graph

import "testing"

func TestMatchEdgeType(t *testing.T) {
	find, _ := GetOrCreateNodeType("find", "Your")
	invalid, _ := GetOrCreateNodeType("invalid", "Moma")
	junk, _ := GetOrCreateNodeType("junk", "junk")
	options := []*nodeType{find, junk}
	t1 := matchEdgeType(find, options)
	if !t1 {
		t.Error("Can't find find!")
	}

	t2 := matchEdgeType(invalid, options)
	if t2 {
		t.Error("Found invalid string!")
	}
}

/*
Show that edge type restrict the types of data that can be
connected to and from. Part of the definition of an edge type
is to limit what connections it can have.
*/
func TestEdgeType(t *testing.T) {
	// Clean the data
	all_edge_types = make(map[string]edgeType)

	// Test the failing edge connections
	FailFrom1, _ := GetOrCreateNodeType("FailFrom1", "Your Moma")
	FailFrom2, _ := GetOrCreateNodeType("FailFrom2", "Your Moma")
	failFromArgs := []*nodeType{FailFrom1, FailFrom2}
	fromData, _ := CreateDataNode(FailFrom1, []byte("Your Moma"))

	FailTo1, _ := GetOrCreateNodeType("FailTo1", "Your Moma")
	FailTo2, _ := GetOrCreateNodeType("FailTo2", "Your Moma")
	failToArgs := []*nodeType{FailTo1, FailTo2}
	toData, _ := CreateDataNode(FailTo1, []byte("Your Moma"))

	CreateEdgeType("Your Moma", failFromArgs, failToArgs)
	nt, _ := GetOrCreateNodeType("Your", "Moma")
	dn, _ := CreateDataNode(nt, []byte("Your Moma"))
	fromGN := GraphNode{fromData, make([]GraphEdge, 0), make([]GraphEdge, 0)}
	toGN := GraphNode{toData, make([]GraphEdge, 0), make([]GraphEdge, 0)}
	et, _ := NewGraphEdge("Your Moma", &fromGN, &toGN)
	gn := GraphNode{}
	gn.Init(nt, dn.data, et, et)
	// This method will fail because we are comparing &{FailTo1 Your Moma} and &{Your Moma}
	if et.EdgeType.ValidToNode(gn) {
		t.Error("Failed this data and edge node don't match")
	}
	// This method will fail because we are comparing &{FailFrom1 Your Moma} and &{Your Moma}
	if et.EdgeType.ValidFromNode(gn) {
		t.Error("Failed this data and edge node don't match")
	}
	if len(fromGN.connectFrom) != 0 {
		t.Error("There should be no graph edge here.")
	}
	if len(fromGN.connectTo) != 0 {
		t.Error("There should be no graph edge here.")
	}
	if len(toGN.connectFrom) != 0 {
		t.Error("There should be no graph edge here.")
	}
	if len(toGN.connectTo) != 0 {
		t.Error("There should be no graph edge here.")
	}


	// Show the graph edges only building correct edges
	nt_correct_from, _ := GetOrCreateNodeType("Correct From Edge", "Moma")
	correctFrom, _ := CreateDataNode(nt_correct_from, []byte("Your Moma"))
	fromGNcorrect := GraphNode{correctFrom, make([]GraphEdge, 0), make([]GraphEdge, 0)}
	correctFromArgs := []*nodeType{nt_correct_from}

	nt_correct_to, _ := GetOrCreateNodeType("Correct To Edge", "Moma")
	correctTo, _ := CreateDataNode(nt_correct_to, []byte("Your Moma"))
	toGNcorrect := GraphNode{correctTo, make([]GraphEdge, 0), make([]GraphEdge, 0)}
	correctToArgs := []*nodeType{nt_correct_to}

	CreateEdgeType("Correct Edge", correctFromArgs, correctToArgs)

	//Fails to create a new edge because this edge type cannot connect from &fromGN
	_, err := NewGraphEdge("Correct Edge", &fromGN, &toGNcorrect)
	if err == nil {
		t.Error("the data node and the from edge type should not match")
	}
	//Fails to create a new edge because this edge type cannot connect to &toGN
	_, err = NewGraphEdge("Correct Edge", &fromGNcorrect, &toGN)
	if err == nil {
		t.Error("the data node and the from edge type should not match")
	}

	et_from, _ := NewGraphEdge("Correct Edge", &fromGNcorrect, &toGNcorrect)
	dn_from, _ := CreateDataNode(nt_correct_from, []byte("Your Moma"))
	gn = GraphNode{}
	gn.Init(nt_correct_from, dn_from.data, et_from, et)
	// The correct from node type and correct from edge
	if !et_from.EdgeType.ValidFromNode(gn) {
		t.Error("Failed this data and edge node do match")
	}
	// The to part of the edge to not set for this graph node
	if et_from.EdgeType.ValidToNode(gn) {
		t.Error("Failed this data and edge node don't match")
	}

	et_to, _ := NewGraphEdge("Correct Edge", &fromGNcorrect, &toGNcorrect)
	dn_to, _ := CreateDataNode(nt_correct_to, []byte("Your Moma"))
	gn = GraphNode{}
	gn.Init(nt_correct_to, dn_to.data, et, et_to)
	// The from part of the edge to not set for this graph node
	if et_to.EdgeType.ValidFromNode(gn) {
		t.Error("Failed this data and edge node don't match")
	}
	// The correct to node type and correct to edge
	if !et_to.EdgeType.ValidToNode(gn) {
		t.Error("Failed this data and edge node do match")
	}
}

func TestGetEdgeType(t *testing.T) {
	_, err := GetEdgeType("Your Skinny Moma")
	if err == nil {
		t.Error("Your moma is not skinny", err)
	}
}

func TestCreateEdgeType(t *testing.T) {
	all_edge_types = make(map[string]edgeType)
	validFromArgs := []*nodeType{
		&nodeType{"ValidFrom1", "Your Moma"},
		&nodeType{"ValidFrom2", "Your Moma"}}
	validToArgs := []*nodeType{
		&nodeType{"ValidTo1", "Your Moma"},
		&nodeType{"ValidTo2", "Your Moma"}}
	et, err := CreateEdgeType("Your Moma", validFromArgs, validToArgs)
	if err != nil {
		t.Error(err)
	}

	_, err = CreateEdgeType("Your Moma", validFromArgs, validToArgs)
	if err == nil {
		t.Error("2 Momas. No wait, times have changes now. Ah what the hell, I'm cracking Your Moma jokes all over my testing, what do you expect")
	}

	test, err := GetEdgeType("Your Moma")
	if err != nil {
		t.Error("This call should get Your Moma")
	}

	if et.edgeTypeName != test.edgeTypeName {
		t.Error("These are supposed to the same Moma")
	}
}
