go_graph
========

My little teach myself go project. A simple little graph DB engine written in go.

The thing I want to explore with this project is the goroutines. My plan for this is write a basic graph DB engine that will be able to build the graph structure when given new data. The current plan is to only have graph edge of predifened types (as in you have to create the edge type before you can use it). Then the data nodes will have lists of valid edges that can connect to the node and another list of edges that can connect from the node. The final step in this process is to creat building logic, so I can tell the DB what connection a node of a given type should have.

EG:

node Person
a Person has from/to edge type sibling
a Person has from edge type child
a Person has To edge type Parent
a Person has to edge type Aunty
a person has From edage type Niece

So I have 3 sisters, when I have a daughter. I only have to tell the DB that I have had a child and the DB engine (in the background using goroutines) will use the give rules for a person to create all the Sibling, Aunt/Niece, Cousin relations automatically

To start with this will be an in memory process only.

Have:
    Creted data store
    edge types

Next step(s):
     UNIT TESTS!!!!!!
     Put data and edge nodes together
     use graph node to keep track of all the edges a node has
     traverse the graph.
     build logic to automatically create edges for a node (start with static example)
     make a system to create this logic as needed
     Write a standard set of graph searches.
     Make an API so other processes can use this graph engine
     Rewrite everything again from scratch using what I have larnt along the way



This DB can only run on 1 machine (add culstering? how?)

Everything is held in memory no permiant store (What file format works for a graph?)
