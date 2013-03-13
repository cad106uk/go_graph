package go_graph

type GraphEdge struct {
	EdgeType    *edgeType
	ConnectFrom *GraphNode
	ConnectTo   *GraphNode
}

// Create new edge. An edge is allowed to link to the same node
func NewGraphEdge(edTy string, from, to *GraphNode) (GraphEdge, error) {
	ge := GraphEdge{}
	et, err := GetEdgeType(edTy)
	if err != nil {
		return ge, err
	}

	if from.value.id == "" {
		return ge, error(&NodeError{"The from dataNode has not been initialised"})
	}
	match := et.ValidFromNode(*from)
	if !match {
		return ge, error(&NodeError{"The from dataNode is invalid for this edge type"})
	}
	tmp := dataNode{}
	if to.value.id == tmp.id {
		to = from
	}

	if to.value.id == "" {
		return ge, error(&NodeError{"The to dataNode hsa not been initialised"})
	}
	match = et.ValidToNode(*to)
	if !match {
		return ge, error(&NodeError{"The to dataNode is invalid for this edge type"})
	}
	ge.EdgeType = et
	ge.ConnectFrom = from
	ge.ConnectTo = to
	return ge, nil
}

// To handle a set of relation eg. Famly would holed brother, sister etc
type RelationSet struct {
	edgeTypes []*edgeType
	name      string
}

func nodeInList(list []*nodeType, ele *nodeType) bool {
	for _, val := range list {
		if val == ele {
			return true
		}
	}
	return false
}

func handleRelSetValid(validList, exploreList []*nodeType) (bool, *nodeType) {
	for _, nt := range validList {
		if nodeInList(exploreList, nt) {
			return true, nt
		}
	}
	return false, &nodeType{}
}

func (rs *RelationSet) ValidFromNode(et *edgeType) (*nodeType, error) {
	for _, edge := range rs.edgeTypes {
		success, output := handleRelSetValid(edge.validFromNodes, et.validFromNodes)
		if success {
			return output, nil
		}
	}
	return &nodeType{}, error(&NodeError{"This edge type is not valid for this RelationSet"})
}

func (rs *RelationSet) ValidToNode(et *edgeType) (*nodeType, error) {
	for _, edge := range rs.edgeTypes {
		success, output := handleRelSetValid(edge.validToNodes, et.validToNodes)
		if success {
			return output, nil
		}
	}
	return &nodeType{}, error(&NodeError{"This edge type is not valid for this RelationSet"})
}
