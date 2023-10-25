package app

import "github.com/Nav1Cr0ss/algorithms-lab2/internal/domain"

type BFS struct {
	title string
}

func NewBFS() *BFS {
	return &BFS{title: "BFS"}
}

func (b *BFS) GetTitle() string {
	return b.title
}

func (b *BFS) SearchExit(maze domain.Maze, start, end domain.Point) []domain.Point {

	queue := []domain.Point{start}

	visited := make(map[domain.Point]bool)

	parent := make(map[domain.Point]domain.Point)

	directions := []domain.Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			path := []domain.Point{end}
			for path[len(path)-1] != start {
				path = append(path, parent[path[len(path)-1]])
			}
			return path
		}

		for _, dir := range directions {
			next := domain.Point{current.X + dir.X, current.Y + dir.Y}

			if maze.IsValid(next) && !visited[next] {
				queue = append(queue, next)
				visited[next] = true
				parent[next] = current
			}
		}
	}

	return nil
}

func (b *BFS) SearchExitDebug(maze domain.Maze, start, end domain.Point, stats *domain.SearchStats) []domain.Point {
	if stats == nil {
		stats = &domain.SearchStats{}
	}

	queue := []domain.Point{start}
	visited := make(map[domain.Point]bool)
	parent := make(map[domain.Point]domain.Point)
	directions := []domain.Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	stats.Iterations = 0
	stats.VisitedStates = 0
	stats.StatesInMemory = 0
	stats.MaxQueueSize = 0
	stats.MaxDepthReached = 0
	stats.FoundGoal = false

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
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

			if maze.IsValid(next) && !visited[next] {
				queue = append(queue, next)
				visited[next] = true
				parent[next] = current
				stats.VisitedStates++
				if len(queue) > stats.MaxQueueSize {
					stats.MaxQueueSize = len(queue)
				}
			}
		}

		if len(queue) > stats.StatesInMemory {
			stats.StatesInMemory = len(queue)
		}

		if len(queue) > stats.MaxDepthReached {
			stats.MaxDepthReached = len(queue)
		}
	}

	return nil
}
