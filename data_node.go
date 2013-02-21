package go_graph

import "sync"

type document interface{}

type NodeError struct {
	msg string
}

func (ne *NodeError) Error() string {
	return ne.msg
}

type nodeType struct {
	name        string
	description string
}

var allNodeTypes = make(map[string]nodeType)

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

func CreateNewNodeType(name, desc string) error {
	_, present := allNodeTypes[name]
	if present {
		return error(&NodeError{"A NodeType with this name has already been created"})
	}

	allNodeTypes[name] = nodeType{name, desc}
	return nil
}

func GetNodeType(name string) (nodeType, error) {
	val, present := allNodeTypes[name]
	if !present {
		return nodeType{}, error(&NodeError{"No NodeType with this name exists"})
	}
	return val, nil
}

func GetOrCreateNodeType(name, desc string) (nodeType, error) {
	val, err := GetNodeType(name)
	if err == nil {
		return val, nil
	}

	err = CreateNewNodeType(name, desc)
	if err == nil {
		return GetNodeType(name)
	}

	return nodeType{}, err
}
