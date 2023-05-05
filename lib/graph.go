package lib

import (
	"fmt"

	servicev1 "github.com/benfdking/tbd/proto/gen/go/tbd/service/v1"
	"github.com/samber/lo"
)

type ProjectGraph struct {
	taken map[string]struct{}

	seeds   map[string]*servicev1.Seed
	models  map[string]*servicev1.Model
	tests   map[string]*servicev1.Test
	sources map[string]*servicev1.Source

	edges [][2]string

	graph *GonumGraph
}

func ProjectToGraph(p *servicev1.Project) (ProjectGraph, error) {
	taken := make(map[string]struct{})
	var edges [][2]string

	// seeds
	for k := range p.Seeds {
		err := errorMapAdder(taken, k, struct{}{})
		if err != nil {
			return ProjectGraph{}, err
		}
	}

	// sources
	for k := range p.Sources {
		err := errorMapAdder(taken, k, struct{}{})
		if err != nil {
			return ProjectGraph{}, err
		}
	}

	// models
	models := make(map[string]*servicev1.Model, len(p.Models))
	for _, m := range p.Models {
		err := errorMapAdder(taken, m.GetName(), struct{}{})
		if err != nil {
			return ProjectGraph{}, err
		}
		models[m.GetName()] = m
	}
	for _, m := range p.Models {
		for _, r := range m.GetReferences() {
			if _, ok := taken[r]; !ok {
				return ProjectGraph{}, fmt.Errorf("reference %s does not exist in keys %v", r, lo.Keys(taken))
			}
			edges = append(edges, [2]string{r, m.GetName()})
		}
	}

	tests := make(map[string]*servicev1.Test, len(p.Tests))
	for _, t := range p.Tests {
		switch {
		case t.GetSql() != nil:
			sqlTest := t.GetSql()
			err := errorMapAdder(taken, sqlTest.GetName(), struct{}{})
			if err != nil {
				return ProjectGraph{}, err
			}
			tests[sqlTest.GetName()] = t

			for _, r := range sqlTest.GetReferences() {
				if _, ok := taken[r]; !ok {
					return ProjectGraph{}, fmt.Errorf("reference %s does not exist in keys %v", r, lo.Keys(taken))
				}
				edges = append(edges, [2]string{r, sqlTest.GetName()})
			}
		case t.GetNotNull() != nil:
			nn := t.GetNotNull()
			name := GenerateTestNameNotNull(nn)
			err := errorMapAdder(taken, name, struct{}{})
			if err != nil {
				return ProjectGraph{}, err
			}
			edges = append(edges, [2]string{nn.GetModel(), name})
		case t.GetUnique() != nil:
			test := t.GetUnique()
			name := GenerateTestNameUnique(test)
			err := errorMapAdder(taken, name, struct{}{})
			if err != nil {
				return ProjectGraph{}, err
			}
			edges = append(edges, [2]string{test.GetModel(), name})
		case t.GetRelationship() != nil:
			test := t.GetRelationship()
			name := GenerateTestNameRelationship(test)
			err := errorMapAdder(taken, name, struct{}{})
			if err != nil {
				return ProjectGraph{}, err
			}
			// TODO Need to add the edges for the relationship test to both sources but also need to check this with a test
			edges = append(edges, [2]string{test.GetSourceModel(), name})
			edges = append(edges, [2]string{test.GetTargetModel(), name})
		case t.GetAcceptedValues() != nil:
			test := t.GetAcceptedValues()
			name := GenerateTestNameAcceptedValues(test)
			err := errorMapAdder(taken, name, struct{}{})
			if err != nil {
				return ProjectGraph{}, err
			}
			edges = append(edges, [2]string{test.GetModel(), name})
		case t.GetCustomColumn() != nil:
			test := t.GetCustomColumn()
			name := GenerateTestNameCustomColumn(test)
			err := errorMapAdder(taken, name, struct{}{})
			if err != nil {
				return ProjectGraph{}, err
			}
			edges = append(edges, [2]string{test.GetModel(), name})
		default:
			return ProjectGraph{}, fmt.Errorf("only sql/not-null/unique/accepted_values/custom tests are supported")
		}
	}

	graph, err := NewGraph(edges)
	if err != nil {
		return ProjectGraph{}, err
	}

	return ProjectGraph{
		taken:   taken,
		seeds:   p.Seeds,
		models:  models,
		sources: p.Sources,
		edges:   edges,
		graph:   graph,
	}, nil
}

func (g ProjectGraph) ToDotViz() ([]byte, error) {
	return g.graph.ToDotViz()
}

// ReturnSubGraphNodes returns the subgraph that is needed to compile a particular model, in the order that if applied
// should be ok.
//
// It will remove duplicate values if they are in output more than once.
func (g *GonumGraph) ReturnSubGraphNodes(model string) ([]int64, error) {
	v, ok := g.nodeDictionary[model]
	if !ok {
		return nil, fmt.Errorf("model %s not found in graph", model)
	}
	parents := g.recursiveReturnParents(v)
	return removeFollowingDuplicatesAnywhere(parents), nil
}

// removeFollowingDuplicatesAnywhere removes duplicates from a slice and keeps the position of the first appearance.
func removeFollowingDuplicatesAnywhere(s []int64) []int64 {
	seen := make(map[int64]struct{})
	j := 0
	for i, v := range s {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			s[j] = s[i]
			j++
		}
	}
	return s[:j]
}

func (g *GonumGraph) recursiveReturnParents(v int64) (order []int64) {
	newParents := g.To(v)
	for newParents.Next() {
		parents := g.recursiveReturnParents(newParents.Node().ID())
		order = append(parents, order...)
	}
	order = append(order, v)
	return order
}

// errorMapAdder adds a key to a map and returns an error if the key already exists
func errorMapAdder[T any](m map[string]T, k string, v T) error {
	if _, ok := m[k]; ok {
		return fmt.Errorf("value %s already exists in %v", k, lo.Keys(m))
	}
	m[k] = v
	return nil
}
