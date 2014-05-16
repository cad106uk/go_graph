package = go_graph

build:
	go fmt
	go build github.com/cad106uk/go_graph/helpers
	go build github.com/cad106uk/go_graph/data_types
	go build github.com/cad106uk/go_graph/node_edges
	go build

test:
	go fmt
	go test github.com/cad106uk/go_graph/helpers
	go test github.com/cad106uk/go_graph/data_types
	go test github.com/cad106uk/go_graph/node_edges
	go test

clean:
	go clean
	git clean -xdf
