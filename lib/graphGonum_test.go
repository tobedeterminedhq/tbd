package lib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestNewGraph tests the NewGraph function. It is a simple test that only tests the length of the nodeDictionary.
func TestNewGraph(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                 string
		edges                [][2]string
		wantDictionaryLength int
		wantErr              bool
	}{
		{
			name:                 "empty",
			edges:                [][2]string{},
			wantDictionaryLength: 0,
			wantErr:              false,
		},
		{
			name:                 "simple",
			edges:                [][2]string{{"a", "b"}},
			wantDictionaryLength: 2,
			wantErr:              false,
		},
		{
			name: "diamond",
			edges: [][2]string{
				{"A", "B"},
				{"A", "C"},
				{"B", "D"},
				{"C", "D"},
			},
			wantDictionaryLength: 4,
			wantErr:              false,
		},
		// TODO need to catch cycles cleanly
		//{
		//	name: "cycle",
		//	edges: [][2]string{
		//		{"A", "B"},
		//		{"B", "A"},
		//	},
		//	wantErr:              true,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewGraph(tt.edges)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGraph() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got.nodeDictionary), tt.wantDictionaryLength) {
				t.Errorf("NewGraph() got = %v, want %v", len(got.nodeDictionary), tt.wantDictionaryLength)
			}
			assert.Equal(t, len(got.nodeDictionary), got.Nodes().Len())
			assert.Equal(t, len(tt.edges), got.Edges().Len())
		})
	}
}

func TestGonumGraph_ToDotViz(t *testing.T) {
	tests := []struct {
		name    string
		edges   [][2]string
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "empty",
			edges:   [][2]string{},
			want:    "strict digraph {\n}",
			wantErr: assert.NoError,
		},
		{
			name: "single edge",
			edges: [][2]string{
				{"A", "B"},
			},
			// TODO Need to add attributes to node definitions
			want: `strict digraph {
// Node definitions.
A;
B;

// Edge definitions.
A -> B;
}`,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := NewGraph(tt.edges)
			if err != nil {
				t.Fatalf("NewGraph() error = %v", err)
			}

			got, err := g.ToDotViz()
			if !tt.wantErr(t, err, fmt.Sprintf("ToDotViz()")) {
				return
			}
			assert.Equalf(t, tt.want, string(got), "ToDotViz()")
		})
	}
}

func TestGonumGraph_GetNodeSorted(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		edges   [][2]string
		want    []string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "empty",
			edges:   [][2]string{},
			want:    []string{},
			wantErr: assert.NoError,
		},
		{
			name: "simple diagram",
			edges: [][2]string{
				{"B", "C"},
				{"A", "B"},
			},
			want:    []string{"A", "B", "C"},
			wantErr: assert.NoError,
		},
		// TODO Add more complex tests
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := NewGraph(tt.edges)
			if err != nil {
				t.Fatal(err)
			}

			got, err := g.GetNodeSorted()

			if !tt.wantErr(t, err, fmt.Sprintf("GetNodeSorted()")) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetNodeSorted()")
		})
	}
}

func TestGonumGraph_GetNode(t *testing.T) {
	tests := []struct {
		name     string
		edges    [][2]string
		nodeName string
		want     int64
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			name:     "empty",
			edges:    [][2]string{},
			nodeName: "C",
			want:     0,
			wantErr:  assert.Error,
		},
		{
			name: "simple diagram",
			edges: [][2]string{
				{"B", "C"},
			},
			nodeName: "B",
			want:     0,
			wantErr:  assert.NoError,
		},
		{
			name: "simple diagram, target",
			edges: [][2]string{
				{"B", "C"},
			},
			nodeName: "C",
			want:     1,
			wantErr:  assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := NewGraph(tt.edges)
			require.NoError(t, err)

			got, err := g.GetNode(tt.nodeName)
			if !tt.wantErr(t, err, fmt.Sprintf("GetNode(%v)", tt.nodeName)) {
				return
			}
			fromDict := g.nodeDictionary[tt.nodeName]
			assert.Equalf(t, fromDict, got, "GetNode(%v)", tt.nodeName)
		})
	}
}

func TestGonumGraph_GetNodeName(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		graph   *GonumGraph
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "simple diagram",
			graph: func() *GonumGraph {
				g, err := NewGraph([][2]string{
					{"B", "C"},
				})
				require.NoError(t, err)
				return g
			}(),
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := tt.graph

			nodes := g.Nodes()
			for nodes.Next() {
				id := nodes.Node().ID()
				name, err := g.GetNodeName(id)
				require.NoError(t, err)
				returnedId, err := g.GetNode(name)
				require.NoError(t, err)
				assert.Equal(t, id, returnedId)
			}
		})
	}
}
