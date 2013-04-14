package go_graph

import "sync"

type nodeType struct {
	name        string
	description string
}

var allNodeTypes = struct {
	sync.RWMutex
	m map[string]nodeType
}{m: make(map[string]nodeType)}

func CreateNewNodeType(name, desc string) error {
	allNodeTypes.Lock()
	defer allNodeTypes.Unlock()
	_, present := allNodeTypes.m[name]
	if present {
		return error(&NodeError{"A NodeType with this name has already been created"})
	}

	allNodeTypes.m[name] = nodeType{name, desc}
	return nil
}

func GetNodeType(name string) (*nodeType, error) {
	allNodeTypes.RLock()
	defer allNodeTypes.RUnlock()
	val, present := allNodeTypes.m[name]
	if !present {
		return &nodeType{}, error(&NodeError{"No NodeType with this name exists"})
	}
	return &val, nil
}

func GetOrCreateNodeType(name, desc string) (*nodeType, error) {
	val, err := GetNodeType(name)
	if err == nil {
		return val, nil
	}

	err = CreateNewNodeType(name, desc)
	if err == nil {
		return GetNodeType(name)
	}

	return &nodeType{}, err
}
