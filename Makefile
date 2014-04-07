package = go_graph

build:
	go build go_graph/helpers
	go build go_graph/data_types
	go build go_graph/node_edges
	go build

test:
	go test go_graph/helpers
	go test go_graph/data_types
	go test go_graph/node_edges
	go test
