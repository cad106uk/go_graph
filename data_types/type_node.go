package data_types

import (
	"github.com/cad106uk/go_graph/helpers"
	"sync"
)

type NodeType struct {
	name        string
	description string
}

func TempUnstoredNodeType(nam, desc string) NodeType {
	return NodeType{nam, desc}
}

var allNodeTypes = struct {
	sync.RWMutex
	m map[string]NodeType
}{m: make(map[string]NodeType)}

func CreateNewNodeType(name, desc string) error {
	allNodeTypes.Lock()
	defer allNodeTypes.Unlock()

	_, present := allNodeTypes.m[name]
	if present {
		return error(helpers.NodeError("A NodeType with this name has already been created"))
	}

	allNodeTypes.m[name] = NodeType{name, desc}
	return nil
}

func GetNodeType(name string) (*NodeType, error) {
	allNodeTypes.RLock()
	defer allNodeTypes.RUnlock()

	val, present := allNodeTypes.m[name]
	if !present {
		return nil, error(helpers.NodeError("No NodeType with this name exists"))
	}
	return &val, nil
}

func GetOrCreateNodeType(name, desc string) (*NodeType, error) {
	val, err := GetNodeType(name)
	if err == nil {
		return val, nil
	}

	err = CreateNewNodeType(name, desc)
	if err == nil {
		return GetNodeType(name)
	}

	return nil, err
}
