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

package go_graph

type GraphNode struct {
	value       dataNode
	connectFrom []GraphEdge // The GraphEdges that use this node as a starting point
	connectTo   []GraphEdge // The GraphEdges the use this node as an end point
}

// returns the value that the GraphNode has
func (gn *GraphNode) Value() (*data, error) {
	val := gn.value.data
	return &val, nil
}

// Set the value this GraphNode stores. Can be called many times but onyl sets a value the first time it has been called.
func (gn *GraphNode) SetValue(input data) error {
	gn.value.setValue.Do(func() {
		gn.value.data = input
	})
	return nil
}

func (gn *GraphNode) Init(nt *nodeType, input data, from, to GraphEdge) error {
	dn, err := CreateDataNode(nt, input.data)
	if err != nil {
		return err
	}
	conTo := append(gn.connectTo, to)
	if err != nil {
		return err
	}
	conFrom := append(gn.connectFrom, from)
	if err != nil {
	}

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

func New(nt *nodeType, input []byte, from, to GraphEdge) (GraphNode, error) {
	gn := GraphNode{}
	dn, err := CreateDataNode(nt, input)
	if err != nil {
		return gn, err
	}
	gn.value = dn
	gn.AddToEdge(to)
	gn.AddFromEdge(from)
	return gn, nil
}

// // A.ConnectTo(B) Means A->From-GraphEdge-To->B
// // If the types allow.
// func (g *GraphNode) ConnectTo(node GraphNode) error {
// 	g.connectsTo = append(g.connectsTo, node.dataNode)
// 	node.connectedFrom = append(node.connectedFrom, g.dataNode)
// 	return nil
// }
