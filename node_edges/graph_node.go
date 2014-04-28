/*
 This package provides the basic structure for building a graph
 the graph node is ment encapsulate a data node and all the
 graph edges the come to and from this node

The current plan is that that date nodes, types and graph edges are passive.
It is the GraphNodes that 'know' about the graph. They link to 1 and only 1
data node. Thay contain all the edges that link to this noe and all the
edges that link from this node. The graph node also contains the logic to
say whether a connection can be made. (So a valid connection can only be
created after 2 GraphNodes have agreed the edge can be made).
*/

package node_edges

import "github.com/cad106uk/go_graph/data_types"

type GraphNode struct {
	value       data_types.DataNode
	connectFrom []GraphEdge // The GraphEdges that use this node as a starting point
	connectTo   []GraphEdge // The GraphEdges the use this node as an end point
}

func (gn *GraphNode) GetConnectFrom() []GraphEdge {
	return gn.connectFrom
}

func (gn *GraphNode) GetConnectTo() []GraphEdge {
	return gn.connectTo
}

// returns the value that the GraphNode has
func (gn *GraphNode) Value() data_types.NodeData {
	return gn.value.NodeData
}

// Set the value this GraphNode stores. Can be called many times but onyl sets a value the first time it has been called.
func (gn *GraphNode) SetValue(input data_types.NodeData) error {
	gn.value.SetValue.Do(func() {
		gn.value.NodeData = input
	})
	return nil
}

func (gn *GraphNode) Init(nt *data_types.NodeType, input data_types.NodeData, from, to GraphEdge) error {
	dn, err := data_types.CreateDataNode(nt, input.GetData())
	if err != nil {
		return err
	}

	conTo := append(gn.connectTo, to)
	conFrom := append(gn.connectFrom, from)
	gn.value = dn
	gn.connectTo = conTo
	gn.connectFrom = conFrom
	return nil
}

func (gn *GraphNode) AddFromEdge(from GraphEdge) error {
	new_edge := append(gn.connectFrom, from)
	gn.connectFrom = new_edge
	return nil
}

func (gn *GraphNode) AddToEdge(to GraphEdge) error {
	new_edge := append(gn.connectTo, to)
	gn.connectTo = new_edge
	return nil
}

func NewGraphNode(nt *data_types.NodeType, input []byte) (GraphNode, error) {
	gn := GraphNode{}
	dn, err := data_types.CreateDataNode(nt, input)
	if err != nil {
		return nil, err
	}
	gn.value = dn
	return gn, nil
}
