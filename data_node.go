package go_graph

import (
	"sync"
	"crypto/sha1"
)

type data struct {
	data []byte
}

type document interface{}

type NodeError struct {
	msg string
}

func (ne *NodeError) Error() string {
	return ne.msg
}

type dataNode struct {
	dataType nodeType
	data           // The data stored at this node
	setValue sync.Once // The value can only be set once
	id string
}

func (dh *dataNode) GetType() nodeType {
	return dh.dataType
}

func (dh *dataNode) GetValue() data {
	return dh.data
}

func (dh *dataNode) GetId() string {
	return dh.id
}

func CreateDataNode(t nodeType, d []byte) (dataNode, error) {
	newData := data{d}
	newNode := dataNode{}
	empty := nodeType{}
	if t == empty {
		return dataNode{}, error(&NodeError{"The nodeType is blank. Must have a valid node type"})
	}
	newNode.dataType = t
	newNode.setValue.Do(func() {
		newNode.data = newData
		hasher := sha1.New()
		hasher.Write([]byte(t.name))
		hasher.Write(d)
		newNode.id = string(hasher.Sum(nil))
	})

	return newNode, nil
}
