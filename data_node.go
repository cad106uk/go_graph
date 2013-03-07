package go_graph

import (
	"crypto/rand"
	"crypto/sha1"
	"io"
	"sync"
)

func genRandNum() *[]byte {
	count := 1024
	rand_store := make([]byte, count)
	io.ReadFull(rand.Reader, rand_store)
	return &rand_store
}

var id_buffer chan []byte = make(chan []byte, 10)
var gen_ids_once sync.Once

func bufferNewIds() {
	for {
		rand_store := genRandNum()
		hasher := sha1.New()
		hasher.Write(*rand_store)
		id_buffer <- hasher.Sum(nil)[0:20]
	}
}

func genIds() {
	gen_ids_once.Do(func() {
		go bufferNewIds()
	})
}

type data struct {
	data []byte
}

type NodeError struct {
	msg string
}

func (ne *NodeError) Error() string {
	return ne.msg
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
		newNode.id  = string(<- id_buffer)
	})

	return newNode, nil
}
