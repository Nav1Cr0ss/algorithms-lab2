package domain

import "fmt"

type Maze [][]int

func NewMaze(data [][]int) Maze {
	maze := make(Maze, len(data))
	for i := range maze {
		maze[i] = make([]int, len(data[i]))
		copy(maze[i], data[i])
	}
	return maze
}

type SearchAlgorithm interface {
	SearchExit(Maze, Point, Point) []Point
	SearchExitDebug(Maze, Point, Point, *SearchStats) []Point
	GetTitle() string
}

func (m Maze) SearchExit(start, end Point, sa SearchAlgorithm, debug bool) []Point {

	var reveredPath []Point

	switch debug {
	case true:
		var stats SearchStats
		reveredPath = sa.SearchExitDebug(m, start, end, &stats)

		printDebugInfo(stats)

	default:
		reveredPath = sa.SearchExit(m, start, end)
	}

	m.reversePath(reveredPath)
	return reveredPath
}

func printDebugInfo(stats SearchStats) {
	fmt.Println("Search Statistics:")
	fmt.Printf("Iterations: %d\n", stats.Iterations)
	fmt.Printf("Visited States: %d\n", stats.VisitedStates)
	fmt.Printf("States In Memory: %d\n", stats.StatesInMemory)
	fmt.Printf("Max Queue Size: %d\n", stats.MaxQueueSize)
	fmt.Printf("Max Depth Reached: %d\n", stats.MaxDepthReached)
	fmt.Printf("Found Goal: %t\n", stats.FoundGoal)
}

func (m Maze) IsValid(point Point) bool {
	rows := len(m)
	cols := len(m[0])
	return point.X >= 0 && point.X < rows && point.Y >= 0 && point.Y < cols && m[point.X][point.Y] == 0
}

func (m Maze) reversePath(path []Point) {
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
}
