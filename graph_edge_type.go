package go_graph

import (
	"sync"
)

func matchEdgeType(name *NodeType, validSlice []NodeType) bool {
	match := false
	for _, val := range validSlice {
		if *name == val {
			match = true
			break
		}
	}
	return match
}

type edgeType struct {
	edgeTypeName   string
	validFromNodes []NodeType // A list of node types
	validToNodes   []NodeType // A list of node types
}

// The name of this edgeType
func (et *edgeType) GetName() string {
	return et.edgeTypeName
}

// The list of nodes this edge can connect from
func (et *edgeType) GetValidFromNodes() []NodeType {
	return et.validFromNodes
}

// The list of  nodes this edge can connect to
func (et *edgeType) GetValidToNode() []NodeType {
	return et.validToNodes
}

// Can this edge connect to spcific node?
func (et *edgeType) ValidToNode(to GraphNode) bool {
	return matchEdgeType(to.value.GetType(), et.validToNodes)
}

// Can this edge connect from a specific node?
func (et *edgeType) ValidFromNode(from GraphNode) bool {
	return matchEdgeType(from.value.GetType(), et.validFromNodes)
}

var allEdgeTypes = struct {
	sync.RWMutex
	m map[string]edgeType
}{m: make(map[string]edgeType)}

func GetEdgeType(name string) (*edgeType, error) {
	allEdgeTypes.RLock()
	defer allEdgeTypes.RUnlock()

	val, present := allEdgeTypes.m[name]
	if !present {
		return nil, error(NodeError("This edgeType does not exist"))
	}
	return &val, nil
}

func CreateEdgeType(name string, validFrom, validTo []NodeType) (*edgeType, error) {
	allEdgeTypes.Lock()
	defer allEdgeTypes.Unlock()

	_, present := allEdgeTypes.m[name]
	if present {
		return nil, error(NodeError("An EdgeType with this name has already been created"))
	}

	newEdge := edgeType{name, validFrom, validTo}
	allEdgeTypes.m[name] = newEdge
	return &newEdge, nil
}
