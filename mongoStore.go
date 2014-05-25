package go_graph

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"sync"
)

type NodeData struct {
	data []byte
}

func (nd *NodeData) GetData() []byte {
	return nd.data
}

type NodeType struct {
	Id          bson.ObjectId `bson:"_id" json:"_id"`
	Name        string        `bson:"Name" json:"Name"`
	Description string        `bson:"Description" json:"Description"`
}

var allNodeTypes = struct {
	sync.RWMutex
	m map[string]NodeType
}{m: make(map[string]NodeType)}

type DataNode struct {
	Id       bson.ObjectId `bson:"_id" json:"_id"`
	Data     []byte        `bson:"Data" json:"Data"` // The data stored at this node
	DataType []NodeType    `bson:"DataType" json:"DataType"`
}

type GraphNode struct {
	Id          bson.ObjectId `bson:"_id" json:"_id"`
	Value       []DataNode    `bson:"Value" json:"Value"`
	ConnectFrom []GraphEdge   `bson:"ConnectFrom" json:"ConnectFrom"` // The GraphEdges that use this node as a starting point
	ConnectTo   []GraphEdge   `bson:"ConnectTo" json:"ConnectTo"`     // The GraphEdges the use this node as an end point
}

type EdgeType struct {
	Id             bson.ObjectId `bson:"_id" json:"_id"`
	EdgeTypeName   string        `bson:"EdgeTypeName" json:"EdgeTypeName"`
	ValidFromNodes []NodeType    `bson:"ValidFromNodes" json:"ValidFromNodes"` // A list of node types
	ValidToNodes   []NodeType    `bson:"ValidToNodes json:"ValidToNodes""`     // A list of node types
}

type GraphEdge struct {
	Id          bson.ObjectId `bson:"_id" json:"_id"`
	EdgeType    []EdgeType    `bson:"EdgeType" json:"EdgeType"`
	ConnectFrom []GraphNode   `bson:"ConnectFrom" json:"ConnectFrom"`
	ConnectTo   []GraphNode   `bson:"ConnectTo" json:"ConnectTo"`
}

func InitFromDB() {
	dbSess, err := mgo.Dial("localhost")
	if err != nil {
		panic("Cannot connect to local instance of MongoDB")
	}
	defer dbSess.Close()

	db := dbSess.DB("go_graph")
	nodeTypes := db.C("node_type")

	query := nodeTypes.Find(nil).Iter() // Find all
	if query == nil {
		return
	}

	var result NodeType
	allNodeTypes.Lock()
	defer allNodeTypes.Unlock()
	for query.Next(&result) {
		allNodeTypes.m[result.Name] = result
	}
}
