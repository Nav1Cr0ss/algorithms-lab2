package app

import (
	"container/heap"
	"github.com/Nav1Cr0ss/algorithms-lab2/internal/domain"
	"github.com/Nav1Cr0ss/algorithms-lab2/pkg/data_types"
	"math"
)

type AStar struct {
	title string
}

func NewAStar() *AStar {

	return &AStar{title: "A*"}
}

func (a *AStar) GetTitle() string {
	return a.title
}

func (a *AStar) SearchExit(maze domain.Maze, start, end domain.Point) []domain.Point {
	pq := make(data_types.PriorityQueue, 0)
	heap.Init(&pq)

	gScore := make(map[domain.Point]float64)

	parent := make(map[domain.Point]domain.Point)

	heap.Push(&pq, &data_types.Item{Value: start, Priority: 0})
	gScore[start] = 0

	directions := []domain.Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for pq.Len() > 0 {
		currentItem := heap.Pop(&pq).(*data_types.Item)

		current, ok := currentItem.Value.(domain.Point)
		if !ok {
			return nil
		}

		if current == end {
			path := []domain.Point{end}
			for path[len(path)-1] != start {
				path = append(path, parent[path[len(path)-1]])
			}
			return path
		}

		for _, dir := range directions {
			next := domain.Point{current.X + dir.X, current.Y + dir.Y}

			if maze.IsValid(next) {
				tentativeGScore := gScore[current] + 1 // 1 - cost of moving to a neighbor

				if _, exists := gScore[next]; !exists || tentativeGScore < gScore[next] {
					parent[next] = current
					gScore[next] = tentativeGScore
					fScore := tentativeGScore + a.heuristic(next, end)
					heap.Push(&pq, &data_types.Item{Value: next, Priority: fScore})
				}
			}
		}
	}

	return nil
}

func (a *AStar) SearchExitDebug(maze domain.Maze, start, end domain.Point, stats *domain.SearchStats) []domain.Point {
	if stats == nil {
		stats = &domain.SearchStats{}
	}

	pq := make(data_types.PriorityQueue, 0)
	heap.Init(&pq)

	gScore := make(map[domain.Point]float64)
	parent := make(map[domain.Point]domain.Point)
	directions := []domain.Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	stats.Iterations = 0
	stats.VisitedStates = 0
	stats.StatesInMemory = 0
	stats.MaxQueueSize = 0
	stats.MaxDepthReached = 0
	stats.FoundGoal = false

	heap.Push(&pq, &data_types.Item{Value: start, Priority: 0})
	gScore[start] = 0

	for pq.Len() > 0 {
		currentItem := heap.Pop(&pq).(*data_types.Item)
		current, ok := currentItem.Value.(domain.Point)
		if !ok {
			return nil
		}
		stats.Iterations++

		if current == end {
			path := []domain.Point{end}
			for path[len(path)-1] != start {
				path = append(path, parent[path[len(path)-1]])
			}
			stats.FoundGoal = true
			return path
		}

		for _, dir := range directions {
			next := domain.Point{current.X + dir.X, current.Y + dir.Y}

			if maze.IsValid(next) {
				tentativeGScore := gScore[current] + 1

				if _, exists := gScore[next]; !exists || tentativeGScore < gScore[next] {
					parent[next] = current
					gScore[next] = tentativeGScore
					fScore := tentativeGScore + a.heuristic(next, end)
					heap.Push(&pq, &data_types.Item{Value: next, Priority: fScore})
				}
			}
		}

		if pq.Len() > stats.MaxQueueSize {
			stats.MaxQueueSize = pq.Len()
		}

		if len(pq) > stats.StatesInMemory {
			stats.StatesInMemory = len(pq)
		}

		if pq.Len() > stats.MaxDepthReached {
			stats.MaxDepthReached = pq.Len()
		}
	}

	return nil
}

func (a *AStar) heuristic(point, end domain.Point) float64 {
	dx := float64(end.X - point.X)
	dy := float64(end.Y - point.Y)
	return math.Sqrt(dx*dx + dy*dy)
}
