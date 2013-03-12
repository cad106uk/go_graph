package go_graph

import "sync"

type data struct {
	data []byte
}

type dataNode struct {
	dataType *nodeType
	data               // The data stored at this node
	setValue sync.Once // The value can only be set once
	id       string
}

func (dh *dataNode) GetType() *nodeType {
	return dh.dataType
}

func (dh *dataNode) GetValue() data {
	return dh.data
}

func (dh *dataNode) GetId() string {
	return dh.id
}

func CreateDataNode(t *nodeType, d []byte) (dataNode, error) {
	genIds()
	newData := data{d}
	newNode := dataNode{}
	empty := nodeType{}
	if *t == empty {
		return dataNode{}, error(&NodeError{"The nodeType is blank. Must have a valid node type"})
	}
	newNode.dataType = t
	newNode.setValue.Do(func() {
		newNode.data = newData
		newNode.id = string(<-id_buffer)
	})

	return newNode, nil
}
