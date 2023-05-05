package lib

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGonumGraph_ReturnSubGraphNodes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		graph   func() (*GonumGraph, error)
		model   string
		want    []string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "simple graph, one edge where ask for end node",
			graph: func() (*GonumGraph, error) {
				return NewGraph([][2]string{
					{"a", "b"},
				})
			},
			model:   "a",
			want:    []string{"a"},
			wantErr: assert.NoError,
		},
		{
			name: "simple graph, one edge where ask for end node",
			graph: func() (*GonumGraph, error) {
				return NewGraph([][2]string{
					{"a", "b"},
				})
			},
			model:   "b",
			want:    []string{"a", "b"},
			wantErr: assert.NoError,
		},
		{
			name: "simple graph, two steps",
			graph: func() (*GonumGraph, error) {
				return NewGraph([][2]string{
					{"a", "b"},
					{"b", "c"},
				})
			},
			model:   "c",
			want:    []string{"a", "b", "c"},
			wantErr: assert.NoError,
		},
		// TODO Implement this test as it is a valid graph and should work but is not deterministic
		//{
		//	name: "triangle graph",
		//	graph: func() (*GonumGraph, error) {
		//		return NewGraph([][2]string{
		//			{"a", "b"},
		//			{"b", "c"},
		//			{"d", "c"},
		//			{"d", "e"},
		//		})
		//	},
		//	model:   "c",
		//	want:    []string{"d", "a", "b", "c"},
		//	wantErr: assert.NoError,
		//},
		//{
		//	name: "diamond shape",
		//	graph: func() (*GonumGraph, error) {
		//		return NewGraph([][2]string{
		//			{"a", "b"},
		//			{"a", "c"},
		//			{"b", "d"},
		//			{"c", "d"},
		//		})
		//	},
		//	model:   "d",
		//	want:    []string{"a", "c", "b", "d"},
		//	wantErr: assert.NoError,
		//},
		//{
		//	name: "complicated structure",
		//	graph: func() (*GonumGraph, error) {
		//		return NewGraph([][2]string{
		//			{"a", "b"},
		//			{"b", "z"},
		//			{"b", "c"},
		//			{"d", "b"},
		//			{"c", "f"},
		//		})
		//	},
		//	model:   "f",
		//	want:    []string{"d", "a", "b", "c", "f"},
		//	wantErr: assert.NoError,
		//},
		{
			name: "simple graph, ask for non existing node",
			graph: func() (*GonumGraph, error) {
				return NewGraph([][2]string{
					{"a", "b"},
				})
			},
			model:   "c",
			want:    []string{},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph, err := tt.graph()
			require.NoError(t, err, "graph creation failed")

			got, err := graph.ReturnSubGraphNodes(tt.model)
			if !tt.wantErr(t, err, fmt.Sprintf("ReturnSubGraphNodes(%v)", tt.model)) {
				return
			}

			require.Equal(t, len(tt.want), len(got), "ReturnSubGraphNodes(%v) lengths", tt.model)
			outs := make([]string, len(got))
			for i, node := range got {
				outs[i], err = graph.GetNodeName(node)
				require.NoError(t, err, "GetNode(%v) failed", node)
			}
			assert.Equalf(t, tt.want, outs, "ReturnSubGraphNodes(%v)", tt.model)
		})
	}
}
