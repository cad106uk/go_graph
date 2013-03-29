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

func genRandNum() []byte {
	count := 1024
	rand_store := make([]byte, count)
	_, err := io.ReadFull(rand.Reader, rand_store)
	if err != nil {
		// It is either this panic. Though panic might be better
		genRandNum()
	}
	return rand_store
}

var id_buffer chan []byte = make(chan []byte, 10)
var end_buffer chan struct{} = make(chan struct{}, 1)
var gen_ids_once sync.Once

func bufferNewIds(done <- chan struct {}) {
	for {
		rand_store := genRandNum()
		hasher := sha1.New()
		hasher.Write(rand_store)
		select {
		case <- done:
			//End now
			return
		case id_buffer <- hasher.Sum(nil)[0:20]:
		}
	}
}

func genIds() {
	gen_ids_once.Do(func() {
		go bufferNewIds(end_buffer)
	})
}
