package dijkstra

import (
	"container/heap"
	"sync"
)

type Node struct {
	val       int
	neighbors []*Node
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].val < pq[j].val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	node := x.(*Node)
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[:n-1]
	return node
}

func DijkstraParallel(graph [][]float64, startNode *Node, endNode *Node) (map[int]float64, map[int]*Node) {
	if startNode == nil || endNode == nil {
		return nil, nil
	}

	queue := make(PriorityQueue, 0)
	heap.Init(&queue)
	heap.Push(&queue, startNode)
	visited := make(map[*Node]bool)
	distances := make(map[int]float64)
	previous := make(map[int]*Node)
	var mutex sync.Mutex // Add a mutex for synchronization

	for i := range graph {
		distances[i] = float64(1<<63 - 1) // Infinity
	}
	distances[startNode.val] = 0

	var worker func()
	worker = func() {
		for {
			mutex.Lock() // Lock before accessing maps
			if queue.Len() == 0 {
				mutex.Unlock() // Unlock if the queue is empty
				return
			}
			node := heap.Pop(&queue).(*Node)
			if visited[node] {
				mutex.Unlock() // Unlock if the node has already been visited
				continue
			}
			visited[node] = true
			mutex.Unlock() // Unlock before iterating over neighbors
			for _, neighbor := range node.neighbors {
				distance := distances[node.val] + graph[node.val][neighbor.val]
				mutex.Lock() // Lock before modifying maps
				if distance < distances[neighbor.val] {
					distances[neighbor.val] = distance
					previous[neighbor.val] = node
					heap.Push(&queue, neighbor)
				}
				mutex.Unlock() // Unlock after modifying maps
			}
		}
	}

	numThreads := 4 // Set number of goroutines
	done := make(chan bool)

	for i := 0; i < numThreads; i++ {
		go func() {
			worker()
			done <- true
		}()
	}

	// Wait for all goroutines to finish
	for i := 0; i < numThreads; i++ {
		<-done
	}

	return distances, previous
}

func GetPath(previous map[int]*Node, endNode *Node, locations []string) ([]string, []int) {
	path := make([]string, 0)
	i := make([]int, 0)
	currentNode := endNode
	for currentNode != nil {
		path = append(path, locations[currentNode.val])
		i = append(i, currentNode.val)
		currentNode = previous[currentNode.val]
	}

	// Reverse the path to get it from start to end
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path, i
}

func Initialize(start int, end int) ([]string, float64, []int) {
	// Define the graph
	graph := [][]float64{
		{0, 3, 5, 6, 2.5, 4, 2, 4, 7, 7},
		{3.3, 0, 4.9, 6, 2.3, 3.8, 1.9, 4.1, 8, 6.9},
		{5, 1.5, 0, 1.5, 5, 2, 6, 2.5, 6, 2},
		{6, 2.8, 1.5, 0, 6.5, 3, 7.5, 2.6, 7.6, 2.3},
		{2.5, 2.3, 5, 6.5, 0, 4, 1.5, 5, 6, 7},
		{4, 3.8, 2, 3, 4, 0, 5, 4.3, 4.9, 3.2},
		{1.9, 4.8, 6.3, 7.5, 1.3, 5, 0, 6, 7.6, 8},
		{4, 4.1, 2.5, 2.6, 5, 4.3, 6, 0, 8, 4},
		{7, 8, 6, 7.6, 6, 4.9, 7.6, 8, 0, 7.3},
		{7, 6.9, 2, 2.3, 7, 3.2, 8, 4, 7.3, 0},
	}
	locations := []string{
		"Казахский Гос Цирк", "Казахский Нац Театр Оперы", "Вознес Каф Собор",
		"Центральная мечеть", "Парк Достык", "Гостиница \"Казахстан\"",
		"Главный Ботанический сад", "MegaPark", "Коктобе телебашня", "Алматинский зоопарк",
	}

	// Create the nodes
	nodes := make([]*Node, len(graph))
	for i := range nodes {
		nodes[i] = &Node{val: i}
	}

	// Set the neighbors for each node
	for i, node := range nodes {
		for j := range graph[i] {
			if graph[i][j] > 0 {
				node.neighbors = append(node.neighbors, nodes[j])
			}
		}
	}

	// Run Dijkstra's algorithm
	startNode := nodes[start]
	endNode := nodes[end]
	distances, previous := DijkstraParallel(graph, startNode, endNode)

	// Get the shortest path
	shortestPath, myindex := GetPath(previous, endNode, locations)

	return shortestPath, distances[endNode.val], myindex
}
