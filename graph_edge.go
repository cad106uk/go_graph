package go_graph

type GraphEdge struct {
	edgeType    edgeType
	connectFrom dataNode
	connectTo   dataNode
}

// Create new edge. An edge is allowed to link to the same node
func NewGraphEdge(edTy string, from, to dataNode) (GraphEdge, error) {
	ge := GraphEdge{}
	et, err := GetEdgeType(edTy)
	if err != nil {
		return ge, err
	}

	match := et.ValidFromNode(from)
	if !match {
		return ge, error(&NodeError{"The from dataNode is invalid for this edge type"})
	}
	tmp := dataNode{}
	if to.id == tmp.id {
		to = from
	}

	match = et.ValidToNode(to)
	if !match {
		return ge, error(&NodeError{"The to dataNode is invalid for this edge type"})
	}
	ge.edgeType = et
	ge.connectFrom = from
	ge.connectTo = to
	return ge, nil
}
