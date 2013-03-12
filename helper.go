// All the helper functions that don't direct functionality for the
// Graph, but help that functionality along. (Didn't want to call this
// file lib because that can get really confusing)
package go_graph

import (
	"crypto/rand"
	"crypto/sha1"
	"io"
	"sync"
)

type NodeError struct {
	msg string
}

func (ne *NodeError) Error() string {
	return ne.msg
}

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
