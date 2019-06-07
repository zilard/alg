package main

import (
	"fmt"
)



type AdjacencyRecord struct {
        adjacencyArray []int
        index int
}


type Graph struct {
	Adjacencies map[int]*AdjacencyRecord
}



// New Graph creates an empty Graph with v vertices
func NewGraph() *Graph {
        g := &Graph{}
        g.Adjacencies = make(map[int]*AdjacencyRecord)
        return g

}


//AddEdge add an edge v-w
func (g *Graph) AddEdge(v, w int) {
        if _, ok := g.Adjacencies[v]; !ok {
                g.Adjacencies[v] = NewAdjacencyRecord()
        }
        g.Adjacencies[v].Add(w)

        if _, ok := g.Adjacencies[w]; !ok {
                g.Adjacencies[w] = NewAdjacencyRecord()
        }
        g.Adjacencies[w].Add(v)
}


func checkIfPresent(arr []int, a int) bool {
	for _, v := range arr {
		if v == a {
			return true
		}
        }
	return false
}



func (b *AdjacencyRecord) Add(i int) {
        if ok := checkIfPresent(b.adjacencyArray, i); !ok {
		b.adjacencyArray = append(b.adjacencyArray, i)
		b.index++
        }
}


func NewAdjacencyRecord() *AdjacencyRecord {
        ar := &AdjacencyRecord{}
        ar.adjacencyArray = make([]int, 0)
        return ar
}


// #################################################################


type BFSPath struct {
        Source int
        DistTo map[int]int
        EdgeTo map[int]int
        Path   Queue
        G      *Graph
}




type Queue interface {
        Enqueue(obj interface{})
        Dequeue() interface{}
        IsEmpty() bool
        Size() int
        Iterate() <-chan interface{}
}


type queueLinkedList struct {
        First  *Node
        Last   *Node
        Length int
}

func NewQueueLinkedList() Queue {
        return &queueLinkedList{}
}


func (q *queueLinkedList) Enqueue(value interface{}) {
        oldLast := q.Last
        q.Last = &Node{}
        q.Last.Value = value

        if q.IsEmpty() {
                q.First = q.Last
        } else {
                oldLast.Next = q.Last
        }

        q.Length++
}


func (q *queueLinkedList) Dequeue() interface{} {
        if !q.IsEmpty() {
                item := q.First.Value
                q.Length--
                q.First = q.First.Next
                if q.Length == 0 {
                        q.Last = q.First
                }
                return item
        }

        return 0
}


func (q *queueLinkedList) IsEmpty() bool {
        return q.Size() == 0
}

func (q *queueLinkedList) Size() int {
        return q.Length
}

func (q *queueLinkedList) Iterate() <-chan interface{} {
        ch := make(chan interface{})
        go func() {
                for {
                        if q.IsEmpty() {
                                break
                        }
                        ch <- q.Dequeue()
                }
                close(ch)
        }()
        return ch
}

















func NewBFSPath(g *Graph, source int) *BFSPath {
        bfsPath := &BFSPath{
                DistTo: make(map[int]int),
                EdgeTo: make(map[int]int),
                G:      g,
                Path:   NewQueueLinkedList(),
                Source: source,
        }
        bfsPath.bfs(source)
        return bfsPath
}





func (b *BFSPath) bfs(v int) {
        queue := NewQueueLinkedList()
        b.DistTo[v] = 0
        queue.Enqueue(v)
        for {
                if queue.IsEmpty() {
                        break
                }
                d := queue.Dequeue().(int)
                b.Path.Enqueue(d)
                for r := range b.G.Adj(d) {
                        if _, ok := b.DistTo[r]; !ok {
                                queue.Enqueue(r)
                                b.EdgeTo[r] = d
                                b.DistTo[r] = 1 + b.DistTo[d]
                        }
                }
        }
}






func GetGraph() *Graph {
        g := NewGraph()

/*
        g.AddEdge(0, 1)
        g.AddEdge(0, 2)
        g.AddEdge(0, 6)
        g.AddEdge(0, 5)
        g.AddEdge(5, 3)
        g.AddEdge(3, 4)
        g.AddEdge(4, 6)
        g.AddEdge(7, 8)
        g.AddEdge(9, 10)
        g.AddEdge(9, 12)
        g.AddEdge(9, 11)
        g.AddEdge(11, 12)
        g.AddEdge(11, 11)
*/


      g.AddEdge(0, 1)
      g.AddEdge(1, 2)
      g.AddEdge(2, 0)
      g.AddEdge(1, 3)

      return g
}






func main() {

	g := GetGraph()

	for k, v := range g.Adjacencies {
		fmt.Printf("key: %v\n", k)
		fmt.Printf("val: %v\n\n", v)
        }


}

