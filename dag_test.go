package dag

import (
	"reflect"
	"testing"
)

func TestConnectedComponentRoots(t *testing.T) {
	testConnectedComponentRoots1 := Graph{
		Nodes: []Node{0, 1, 2, 3},
		Edges: []Edge{
			{Depender: 2, Dependee: 0},
			{Depender: 3, Dependee: 1},
		},
	}
	want1 := testConnectedComponentRoots1.ConnectedComponentRoots()
	got1 := []Node{0, 1}
	if !reflect.DeepEqual(
		want1, got1) {
		t.Errorf("testConnectedComponentRoots1: wanted: %v, got: %v", want1, got1)
	}
	testConnectedComponentRoots2 := Graph{
		Nodes: []Node{0, 1, 2, 3, 4, 5, 6},
		Edges: []Edge{
			{Depender: 2, Dependee: 0},
			{Depender: 3, Dependee: 1},
			{Depender: 4, Dependee: 2},
			{Depender: 5, Dependee: 2},
			{Depender: 6, Dependee: 5},
			},
	}
	want2 := testConnectedComponentRoots2.ConnectedComponentRoots()
	got2 := []Node{0, 1}
	if !reflect.DeepEqual(
		want2, got2) {
		t.Errorf("testConnectedComponentRoots2: wanted: %v, got: %v", want2, got2)
	}
}

func TestDirectDependers(t *testing.T) {
	testDirectDependers1 := Graph{
		Nodes: []Node{0, 1, 2, 3, 4, 5, 6},
		Edges: []Edge{
			{Depender: 2, Dependee: 0},
			{Depender: 3, Dependee: 1},
			{Depender: 4, Dependee: 2},
			{Depender: 5, Dependee: 2},
			{Depender: 6, Dependee: 5},
		},
	}
	want1 := testDirectDependers1.DirectDependers(2)
	got1 := []Node{4, 5}
	if !reflect.DeepEqual(
		want1, got1) {
		t.Errorf("testDirectDependers1: wanted: %v, got: %v", want1, got1)
	}
}

func TestDirectDependees(t *testing.T) {
	testDirectDependees := Graph{
		Nodes: []Node{0, 1, 2, 3, 4, 5, 6},
		Edges: []Edge{
			{Depender: 2, Dependee: 0},
			{Depender: 3, Dependee: 1},
			{Depender: 4, Dependee: 2},
			{Depender: 5, Dependee: 2},
			{Depender: 6, Dependee: 5},
		},
	}
	want1 := testDirectDependees.DirectDependees(2)
	got1 := []Node{0}
	if !reflect.DeepEqual(
		want1, got1) {
		t.Errorf("testDirectDependees1: wanted: %v, got: %v", want1, got1)
	}
}
