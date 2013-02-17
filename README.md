go_graph
========

My little teach myself go project. A simple little graph DB engine written in go. I doubt it will have any use for anyone else.

I am writing a simple little graph DB engine. It will have nodes and edges, the nodes and edges will have types. So node will connect to edges and edges will connect to node. An edge will have a start and an end and will connect to 2 nodes (or loop back on to the node it started from). The edge will have a list of valid node types it can start from and a different list of node types it can end at.

So only preset valid connections are possible.

Traverse the graph the dataNodes are wrapped by a GraphNode. A GraphNode contains 1 dataNode and 1 list of all the connections this node starts at and 1 list of all the connections this node ends at. Each edge can at most have 2 different dataNodes attached to it. The dataNodes are pasive, they know nothing about edges other node or anything other than there own data. The GraphNode wraps the dataNode can keeps a list of everything that connects to and from the data.

Next step(s):
     UNIT TESTS!!!!!!
     Use the GraphNodes to write a means of traversing the graph
     The functions/methods to create nodes and edges are all currently blocking, must find a safe and useful way of wrapping them in goroutines.


This DB can only run on 1 machine (add culster? how?)

Everything is held in memory no permiant store (yet?)
