package node_edges

import (
	"github.com/cad106uk/go_graph/data_types"
	"github.com/cad106uk/go_graph/helpers"
	"sync"
)

func matchEdgeType(name *data_types.NodeType, validSlice []data_types.NodeType) bool {
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
	validFromNodes []data_types.NodeType // A list of node types
	validToNodes   []data_types.NodeType // A list of node types
}

// The name of this edgeType
func (et *edgeType) GetName() string {
	return et.edgeTypeName
}

// The list of nodes this edge can connect from
func (et *edgeType) GetValidFromNodes() []data_types.NodeType {
	return et.validFromNodes
}

// The list of  nodes this edge can connect to
func (et *edgeType) GetValidToNode() []data_types.NodeType {
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
		return nil, error(helpers.NodeError("This edgeType does not exist"))
	}
	return &val, nil
}

func CreateEdgeType(name string, validFrom, validTo []data_types.NodeType) (edgeType, error) {
	allEdgeTypes.Lock()
	defer allEdgeTypes.Unlock()

	_, present := allEdgeTypes.m[name]
	if present {
		return nil, error(helpers.NodeError("An EdgeType with this name has already been created"))
	}

	newEdge := edgeType{name, validFrom, validTo}
	allEdgeTypes.m[name] = newEdge
	return newEdge, nil
}
