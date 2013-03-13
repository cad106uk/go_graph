package go_graph

import "fmt"

func matchEdgeType(name *nodeType, validSlice []*nodeType) bool {
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
	validFromNodes []*nodeType // A list of node types
	validToNodes   []*nodeType // A list of node types
}

// The name of this edgeType
func (et *edgeType) GetName() string {
	return et.edgeTypeName
}

// The list of nodes this edge can connect from
func (et *edgeType) GetValidFromNodes() []*nodeType {
	return et.validFromNodes
}

// The list of  nodes this edge can connect to
func (et *edgeType) GetValidToNode() []*nodeType {
	return et.validToNodes
}

// Can this edge connect to spcific node?
func (et *edgeType) ValidToNode(to GraphNode) bool {
	fmt.Println("------------------------------")
	fmt.Println(to)
	fmt.Println(to.value)
	fmt.Println(to.value.dataType)
	fmt.Println("------------------------------")
	fmt.Println(et)
	fmt.Println(et.validToNodes)
	fmt.Println("------------------------------")
	return matchEdgeType(to.value.dataType, et.validToNodes)
}

// Can this edge connect from a specific node?
func (et *edgeType) ValidFromNode(from GraphNode) bool {
	return matchEdgeType(from.value.dataType, et.validFromNodes)
}

var all_edge_types map[string]edgeType = make(map[string]edgeType)

func GetEdgeType(name string) (*edgeType, error) {
	val, present := all_edge_types[name]
	if !present {
		return &edgeType{}, error(&NodeError{"This edgeType does not exist"})
	}
	return &val, nil
}

func CreateEdgeType(name string, validFrom, validTo []*nodeType) (edgeType, error) {
	_, present := all_edge_types[name]
	if present {
		return edgeType{}, error(&NodeError{"An EdgeType with this name has already been created"})
	}

	new_edge := edgeType{name, validFrom, validTo}
	all_edge_types[name] = new_edge
	return new_edge, nil
}
