package lib

import (
	"errors"
	"fmt"

	"gonum.org/v1/gonum/graph/topo"

	"github.com/samber/lo"
	"gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/simple"
)

type GonumGraph struct {
	nodeDictionary map[string]int64
	*simple.DirectedGraph
}

type GonumGraphNode struct {
	name string
	simple.Node
}

func (g GonumGraphNode) DOTID() string {
	return g.name
}

func NewGraph(edges [][2]string) (*GonumGraph, error) {
	m := make(map[string]int64)
	count := int64(0)
	for _, edge := range edges {
		k := edge[0]
		v := edge[1]
		if _, ok := m[k]; !ok {
			m[k] = count
			count++
		}
		if _, ok := m[v]; !ok {
			m[v] = count
			count++
		}
	}
	inverted := lo.Invert(m)

	directedGraph := simple.NewDirectedGraph()
	for i := int64(0); i < count; i++ {
		directedGraph.AddNode(GonumGraphNode{
			name: inverted[i],
			Node: simple.Node(i),
		})
	}
	for _, e := range edges {
		from, ok := m[e[0]]
		if !ok {
			return nil, fmt.Errorf("could not find node %s", e[0])
		}
		to, ok := m[e[1]]
		if !ok {
			return nil, fmt.Errorf("could not find node %s", e[1])
		}
		e := directedGraph.NewEdge(directedGraph.Node(from), directedGraph.Node(to))
		directedGraph.SetEdge(e)
	}

	g := GonumGraph{
		nodeDictionary: m,
		DirectedGraph:  directedGraph,
	}

	cycles := topo.DirectedCyclesIn(g)
	if len(cycles) != 0 {
		return nil, errors.New("cycles in graph")
	}

	return &g, nil
}

func (g *GonumGraph) ToDotViz() ([]byte, error) {
	return dot.Marshal(g, "", "", "")
}

// GetNodeSorted returns the nodes in the graph in a sorted order so that they can be applied in the right order.
// The order is determined by the order of the nodes in the graph.
//
// TODO Make this deterministic by sorting the nodes by name.
func (g *GonumGraph) GetNodeSorted() ([]string, error) {
	nodes, err := topo.Sort(g)
	if err != nil {
		return nil, err
	}
	outs := make([]string, len(nodes))
	reverse := lo.Invert(g.nodeDictionary)
	for i, node := range nodes {
		out, ok := reverse[node.ID()]
		if !ok {
			return nil, fmt.Errorf("could not find node %d", node.ID())
		}
		outs[i] = out
	}
	return outs, nil
}

// GetNode returns node int64 for a given node name.
func (g *GonumGraph) GetNode(name string) (int64, error) {
	node, ok := g.nodeDictionary[name]
	if !ok {
		return 0, fmt.Errorf("node %s not found", name)
	}
	return node, nil
}

func (g *GonumGraph) GetNodeName(id int64) (string, error) {
	reverse := lo.Invert(g.nodeDictionary)
	node, ok := reverse[id]
	if !ok {
		return "", fmt.Errorf("node %d not found", id)
	}
	return node, nil
}
