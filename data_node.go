package go_graph

import "sync"

type document interface{}

type NodeError struct {
	msg string
}

func (ne *NodeError) Error() string {
	return ne.msg
}

type dataNode struct {
	dataType nodeType
	document           // The data stored at this node
	setValue sync.Once // The value can only be set once
}

func (dn *dataNode) Init(dt string, doc document) error {
	val, err := GetNodeType(dt)
	if err != nil {
		return err
	}

	dn.dataType = val
	dn.setValue.Do(func() {
		dn.document = doc
	})
	return nil
}