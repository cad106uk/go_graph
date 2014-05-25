// All the helper functions that don't direct functionality for the
// Graph, but help that functionality along. (Didn't want to call this
// file lib because that can get really confusing)
package helpers

import (
	"crypto/rand"
	"crypto/sha1"
	"io"
	"sync"
)

type nodeError struct {
	msg string
}

func (ne *nodeError) Error() string {
	return ne.msg
}

func NodeError(val string) *nodeError {
	return &nodeError{val}
}

func genRandNum() []byte {
	count := 1024
	randStore := make([]byte, count)
	_, err := io.ReadFull(rand.Reader, randStore)
	if err != nil {
		// It is either this panic. Though panic might be better
		genRandNum()
	}
	return randStore
}

var idBuffer chan []byte = make(chan []byte, 10)
var endBuffer chan struct{} = make(chan struct{}, 1)
var genIdsOnce sync.Once

func bufferNewIds(done <-chan struct{}) {
	for {
		randStore := genRandNum()
		hasher := sha1.New()
		hasher.Write(randStore)
		select {
		case <-done:
			//End now
			return
		case idBuffer <- hasher.Sum(nil)[0:20]:
		}
	}
}

func GetId() []byte {
	genIdsOnce.Do(func() {
		go bufferNewIds(endBuffer)
	})
	return <-idBuffer
}
