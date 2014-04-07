package data_types

import (
	"go_graph/helpers"
	"sync"
)

type NodeData struct {
	data []byte
}

func (nd *NodeData) GetData() []byte {
	return nd.data
}

type DataNode struct {
	dataType *NodeType
	NodeData           // The data stored at this node
	SetValue sync.Once // The value can only be set once
	id       string
}

func (dh *DataNode) GetType() *NodeType {
	return dh.dataType
}

func (dh *DataNode) GetValue() NodeData {
	return dh.NodeData
}

func (dh *DataNode) GetId() string {
	return dh.id
}

func CreateDataNode(t *NodeType, d []byte) (DataNode, error) {
	helpers.GenIds()
	newData := NodeData{d}
	newNode := DataNode{}
	empty := NodeType{}
	if *t == empty {
		return DataNode{},
			error(helpers.NodeError("The NodeType is blank. Must have a valid node type"))
	}
	newNode.dataType = t
	newNode.SetValue.Do(func() {
		newNode.NodeData = newData
		newNode.id = string(helpers.GetId())
	})

	return newNode, nil
}
