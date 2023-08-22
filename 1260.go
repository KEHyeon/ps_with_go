package main

import (
	"container/list"
	"fmt"
	"sort"
)

var (
	n, m, startNode int
	graph           map[int][]int
	visited         map[int]bool
)

func DFS(start int) {
	visited[start] = true
	fmt.Printf("%d ", start)
	for _, neighbor := range graph[start] {
		if !visited[neighbor] {
			DFS(neighbor)
		}
	}
}
func BFS(start int) {
	queue := list.New()
	queue.PushBack(start)
	for queue.Len() > 0 {
		cur := queue.Front().Value.(int)
		queue.Remove(queue.Front())
		for _, nxt := range graph[cur] {
			if !visited[nxt] {
				visited[nxt] = true
				fmt.Printf("%d ", nxt)
				queue.PushBack(nxt)
			}
		}
	}
}
func main() {
	fmt.Scan(&n, &m, &startNode)
	graph = make(map[int][]int)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	for i := 1; i <= n; i++ {
		sort.Ints(graph[i])
	}
	visited = make(map[int]bool)
	DFS(startNode)
	visited = make(map[int]bool)
	fmt.Println("")
	visited[startNode] = true
	fmt.Printf("%d ", startNode)
	BFS(startNode)
}
