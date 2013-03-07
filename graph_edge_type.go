package go_graph

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

type edgeType struct {
	edgeTypeName   string
	validFromNodes []string // A list of node types
	validToNodes   []string // A list of node types
}

// The name of this edgeType
func (et *edgeType) GetName() string {
	return et.edgeTypeName
}

// The list of nodes this edge can connect from
func (et *edgeType) GetValidFromNodes() []string {
	return et.validFromNodes
}

// The list of  nodes this edge can connect to
func (et *edgeType) GetValidToNode() []string {
	return et.validToNodes
}

// Can this edge connect to spcific node?
func (et *edgeType) ValidToNode(to dataNode) bool {
	return matchEdgeType(to.dataType.name, et.validToNodes)
}

// Can this edge connect from a specific node?
func (et *edgeType) ValidFromNode(from dataNode) bool {
	return matchEdgeType(from.dataType.name, et.validFromNodes)
}

var all_edge_types map[string]edgeType = make(map[string]edgeType)

func GetEdgeType(name string) (edgeType, error) {
	val, present := all_edge_types[name]
	if !present {
		return edgeType{}, error(&NodeError{"This edgeType does not exist"})
	}
	return val, nil
}

func CreateEdgeType(name string, validFrom []string, validTo []string) (edgeType, error) {
	_, present := all_edge_types[name]
	if present {
		return edgeType{}, error(&NodeError{"An EdgeType with this name has already been created"})
	}

	new_edge := edgeType{name, validFrom, validTo}
	all_edge_types[name] = new_edge
	return new_edge, nil
}
