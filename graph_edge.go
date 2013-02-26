package go_graph

type edgeType struct {
	EdgeTypeName   string
	ValidFromNodes []string // A list of node types
	ValidToNodes   []string // A list of node types
}

func matchEdgeType(name string, validSlice []string) bool {
	match := false
	for _, val := range validSlice {
		if name == val {
			match = true
			break
		}
	}
	return match
}

func (et *edgeType) ValidToNode(to dataNode) bool {
	return matchEdgeType(to.dataType.name, et.ValidToNodes)
}

func (et *edgeType) ValidFromNode(from dataNode) bool {
	return matchEdgeType(from.dataType.name, et.ValidFromNodes)
}

var all_edge_types map[string]edgeType

func GetEdgeType(name string) (edgeType, error) {
	val, present := all_edge_types[name]
	if !present {
		return edgeType{}, error(&NodeError{"This edgeType does not exist"})
	}
	return val, nil
}

func CreateEdgeType(name string, validFrom []string, validTo []string) error {
	_, present := all_edge_types[name]
	if present {
		return error(&NodeError{"An EdgeType with this name has already been created"})
	}

	new_edge := edgeType{name, validFrom, validTo}
	all_edge_types[name] = new_edge
	return nil
}

type GraphEdge struct {
	edgeType    edgeType
	connectFrom dataNode
	connectTo   dataNode
}

// Create new edge. An edge is allowed to link to the same node
func (ge *GraphEdge) Init(edTy string, from, to dataNode) error {
	et, err := GetEdgeType(edTy)
	if err != nil {
		return err
	}

	match := et.ValidFromNode(from)
	if !match {
		return error(&NodeError{"The from dataNode is invalid for this edge type"})
	}
	tmp := dataNode{}
	if to.id == tmp.id {
		to = from
	}

	match = et.ValidToNode(to)
	if !match {
		return error(&NodeError{"The to dataNode is invalid for this edge type"})
	}
	ge.edgeType = et
	ge.connectFrom = from
	ge.connectTo = to
	return nil
}
