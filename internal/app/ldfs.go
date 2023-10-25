package app

import "github.com/Nav1Cr0ss/algorithms-lab2/internal/domain"

type LDFS struct {
	title string
}

func NewLDFS() *LDFS {
	return &LDFS{title: "LDFS"}
}

func (b *LDFS) GetTitle() string {
	return b.title
}

func (l *LDFS) SearchExit(maze domain.Maze, start, end domain.Point) []domain.Point {
	stack := []domain.Point{start}

	visited := make(map[domain.Point]bool)

	parent := make(map[domain.Point]domain.Point)

	directions := []domain.Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

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
				stack = append(stack, next)
				visited[next] = true
				parent[next] = current
			}
		}
	}

	return nil
}

func (l *LDFS) SearchExitDebug(maze domain.Maze, start, end domain.Point, stats *domain.SearchStats) []domain.Point {
	if stats == nil {
		stats = &domain.SearchStats{}
	}

	stack := []domain.Point{start}
	visited := make(map[domain.Point]bool)
	parent := make(map[domain.Point]domain.Point)
	directions := []domain.Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	stats.Iterations = 0
	stats.VisitedStates = 0
	stats.StatesInMemory = 0
	stats.MaxQueueSize = 0
	stats.MaxDepthReached = 0
	stats.FoundGoal = false

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
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
				stack = append(stack, next)
				visited[next] = true
				parent[next] = current
				stats.VisitedStates++
				if len(stack) > stats.MaxQueueSize {
					stats.MaxQueueSize = len(stack)
				}
			}
		}

		if len(stack) > stats.StatesInMemory {
			stats.StatesInMemory = len(stack)
		}

		if len(stack) > stats.MaxDepthReached {
			stats.MaxDepthReached = len(stack)
		}
	}

	return nil
}
