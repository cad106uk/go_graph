package node_edges

import (
	"github.com/cad106uk/go_graph/data_types"
	"github.com/cad106uk/go_graph/helpers"
)

type GraphEdge struct {
	EdgeType    *edgeType
	ConnectFrom *GraphNode
	ConnectTo   *GraphNode
}

// Create new edge. An edge is allowed to link to the same node
func NewGraphEdge(edTy string, from, to *GraphNode) (*GraphEdge, error) {
	et, err := GetEdgeType(edTy)
	if err != nil {
		return nil, err
	}

	if from.value.GetId() == "" {
		return nil, error(helpers.NodeError("The from dataNode has not been initialised"))
	}

	match := et.ValidFromNode(*from)
	if !match {
		return nil, error(helpers.NodeError("The from dataNode is invalid for this edge type"))
	}

	tmp := data_types.DataNode{}
	if to.value.GetId() == tmp.GetId() {
		to = from
	}

	if to.value.GetId() == "" {
		return nil, error(helpers.NodeError("The to dataNode hsa not been initialised"))
	}

	match = et.ValidToNode(*to)
	if !match {
		return nil, error(helpers.NodeError("The to dataNode is invalid for this edge type"))
	}

	ge := GraphEdge{}
	ge.EdgeType = et
	ge.ConnectFrom = from
	ge.ConnectTo = to
	return &ge, nil
}

// To handle a set of relation eg. Famly would holed brother, sister etc
type RelationSet struct {
	edgeTypes []edgeType
	name      string
}

func nodeInList(list []data_types.NodeType, ele *data_types.NodeType) bool {
	for _, val := range list {
		if &val == ele {
			return true
		}
	}
	return false
}

func handleRelSetValid(validList, exploreList []data_types.NodeType) (bool, *data_types.NodeType) {
	for _, nt := range validList {
		if nodeInList(exploreList, &nt) {
			return true, &nt
		}
	}
	return false, nil
}

func (rs *RelationSet) ValidFromNode(et *edgeType) (*data_types.NodeType, error) {
	for _, edge := range rs.edgeTypes {
		success, output := handleRelSetValid(edge.validFromNodes, et.validFromNodes)
		if success {
			return output, nil
		}
	}
	return nil, error(helpers.NodeError("This edge type is not valid for this RelationSet"))
}

func (rs *RelationSet) ValidToNode(et *edgeType) (*data_types.NodeType, error) {
	for _, edge := range rs.edgeTypes {
		success, output := handleRelSetValid(edge.validToNodes, et.validToNodes)
		if success {
			return output, nil
		}
	}
	return nil, error(helpers.NodeError("This edge type is not valid for this RelationSet"))
}
