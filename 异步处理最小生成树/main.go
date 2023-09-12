package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 定义边的数据结构
type Edge struct {
	Start, End int
	Weight     int
}

// 优先队列的数据结构
type PriorityQueue struct {
	edge       []Edge      // 存储边的切片
	pointIndex map[int]int // 用于跟踪顶点在优先队列中的位置
}

func (pq PriorityQueue) Len() int { return len(pq.edge) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq.edge[i].Weight < pq.edge[j].Weight
}

func (pq PriorityQueue) Swap(i, j int) {
	// 交换边的位置并更新pointIndex中的索引
	pq.edge[i], pq.edge[j] = pq.edge[j], pq.edge[i]
	pq.pointIndex[pq.edge[i].End] = i
	pq.pointIndex[pq.edge[j].End] = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	// 向优先队列中添加边并更新pointIndex中的索引
	pq.edge = append(pq.edge, x.(Edge))
	pq.pointIndex[x.(Edge).End] = pq.Len() - 1
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *&pq.edge
	n := len(old)
	x := old[n-1]
	pq.edge = old[0 : n-1]
	return x
}

// 定义一个结构来包含优先队列和锁
type PriorityQueueWithLock struct {
	PriorityQueue // 嵌套PriorityQueue作为内部字段
	mutex         sync.Mutex
}

// 添加边到优先队列，加锁以确保线程安全
func (pq *PriorityQueueWithLock) AddEdge(edge Edge) {
	pq.mutex.Lock()
	heap.Push(&pq.PriorityQueue, edge)
	pq.mutex.Unlock()
}

// // 读取pointIndex中的索引，加锁以确保线程安全
// func (pq *PriorityQueueWithLock) ReadMap(key int) (index int, ok bool) {
// 	pq.mutex.Lock()
// 	index, ok = pq.pointIndex[key]
// 	pq.mutex.Unlock()
// 	return
// }

// // 写入pointIndex中的索引，加锁以确保线程安全
// func (pq *PriorityQueueWithLock) WriteMap(key, value int) {
// 	pq.mutex.Lock()
// 	pq.pointIndex[key] = value
// 	pq.mutex.Unlock()
// }

// 从优先队列中弹出最小边，加锁以确保线程安全
func (pq *PriorityQueueWithLock) PopMinEdge() Edge {
	pq.mutex.Lock()
	minEdge := heap.Pop(&pq.PriorityQueue).(Edge)
	pq.mutex.Unlock()
	return minEdge
}

// 使用Goroutine实现Prim算法
func Prim(graph [][]int, wg *sync.WaitGroup) int {
	defer wg.Done()

	numVertices := len(graph)
	mst := make([]bool, numVertices)
	mstEdges := make([]Edge, 0)
	pq := &PriorityQueueWithLock{PriorityQueue{pointIndex: make(map[int]int)}, sync.Mutex{}}

	// 在初始化时将第一个顶点添加到最小生成树中
	mst[0] = true

	// 将与第一个顶点相邻的边添加到优先队列中
	for j := 0; j < numVertices; j++ {
		if graph[0][j] > 0 {
			pq.AddEdge(Edge{0, j, graph[0][j]})
		}
	}
	routinueNums := make(chan struct{}, 50)
	for i := 0; i < numVertices-1; i++ {
		// 从优先队列中弹出最小边，确保边的终点不在最小生成树中
		var minEdge Edge
		for {
			minEdge = pq.PopMinEdge()
			if !mst[minEdge.End] {
				break
			}
		}

		mst[minEdge.End] = true
		mstEdges = append(mstEdges, minEdge)

		// 同时启动多个Goroutine来探索连接到最小生成树的未访问顶点并添加边
		var innerWg sync.WaitGroup
		for j := 0; j < numVertices; j++ {
			edgeValue := graph[minEdge.End][j]
			if !mst[j] && edgeValue > 0 {
				innerWg.Add(1)
				routinueNums <- struct{}{}
				go func(j int) {
					defer innerWg.Done()
					if waitChangeIndex, ok := pq.pointIndex[j]; ok {
						if pq.edge[waitChangeIndex].Weight > edgeValue {
							pq.edge[waitChangeIndex].Weight = edgeValue
							pq.edge[waitChangeIndex].Start = minEdge.End
							heap.Fix(pq, waitChangeIndex)
						}
					} else {
						pq.AddEdge(Edge{minEdge.End, j, graph[minEdge.End][j]})
					}
					<-routinueNums
				}(j)
			}
		}
		innerWg.Wait()
	}

	// 计算最小生成树的权重
	totalWeight := 0
	for _, edge := range mstEdges {
		totalWeight += edge.Weight
	}
	return totalWeight
}

// 生成一个随机图
func generateRandomGraph(size int) [][]int {
	rand.Seed(time.Now().UnixNano())

	// 初始化图矩阵
	graph := make([][]int, size)
	for i := range graph {
		graph[i] = make([]int, size)
	}

	// 随机生成边的权重（大于0且小于10）
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			weight := rand.Intn(9) + 1 // 随机生成1到9的权重
			graph[i][j] = weight
			graph[j][i] = weight // 无向图，需要对称
		}
	}

	return graph
}

func main() {
	// 示例图的邻接矩阵
	graph := generateRandomGraph(10000)
	var wg sync.WaitGroup
	wg.Add(1)
	start := time.Now()
	go func() {
		defer wg.Done()
		wg.Add(1)
		weight := Prim(graph, &wg)
		end := time.Now()
		fmt.Printf("最小生成树的权重：%d\n时间消耗为 %v", weight, end.Sub(start))
	}()
	wg.Wait()
}
