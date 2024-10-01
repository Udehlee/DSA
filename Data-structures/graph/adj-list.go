package main

import "fmt"

//Graph hold the vertex
type Graph struct {
	vertex []*Vertex
}

//vertex holds the key and its relationship
type Vertex struct {
	Key  int
	Edge []*EdgeList
}

type EdgeList struct {
	Weight    int
	Dstvertex *Vertex
}

func (g *Graph) AddVertex(key int) {
	g.vertex = append(g.vertex, &Vertex{Key: key})

}

func (g *Graph) Print() {
	for _, v := range g.vertex {
		fmt.Printf("Vertex %v: \n", v.Key)
		for _, v := range v.Edge {
			fmt.Printf("%v", v.Dstvertex.Key)
		}
	}
}

func main() {
	test := &Graph{}
	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}
	test.AddVertex(0)
	test.AddVertex(0)
	test.AddVertex(0)
	test.Print()

}
