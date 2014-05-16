package go_graph

/*
 We need to let user build er there own custom search algorithims. There are predefined algos in graph_path.go that maybe useful here.

 These need to be stored so they can be used again. (And deleted)

 ---These can either run indefinately returning relivant new data when is it added to that graph, or have a given termination criteria--- Don't know how to have a usfull running return code yet.

 These must be able to handle loops within the graph. There must be a default catch all that can be over ridden by the user

 The results will be return on a channel these explorations will be run in goroutines.

 All edges are unidirectional. Each node has 2 list of edges, everything that connects to it and everything it connects to. Although it is possible to go backwards down an edge, it is (should) always be obvious that this is what you are doing.

 For a search to work we will need need a starting node, a list of valid edges to follow, these edges will take us to a new set of nodes. Once at this second set of nodes we will either have a list of edge types to follow or we will stop.

 Each search will have 1 output channel

 Each step will be done in a seperate goroutine. That is each time we go down a graph edge we will do so in a different gorourine. Each goroutine wil also be passed the search criteria.

 The first draft of the search criteria will be a simple list of arrays to consists of node - edge - node - edge so we can fan out our search. This will probably need to be improved on.
*/
