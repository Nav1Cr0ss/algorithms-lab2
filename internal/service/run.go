package service

import (
	"fmt"
	"github.com/Nav1Cr0ss/algorithms-lab2/internal/app"
	"github.com/Nav1Cr0ss/algorithms-lab2/internal/domain"
	"github.com/Nav1Cr0ss/algorithms-lab2/pkg/utilz"
	"runtime"
	"time"
)

func getSearchProviders(algorithm string) []domain.SearchAlgorithm {
	var providers []domain.SearchAlgorithm

	switch algorithm {
	case "BFS":
		providers = append(providers, app.NewBFS())
	case "LDFS":
		providers = append(providers, app.NewLDFS())
	case "ASTAR":
		providers = append(providers, app.NewAStar())
	default:
		providers = append(providers, app.NewBFS())
		providers = append(providers, app.NewLDFS())
		providers = append(providers, app.NewAStar())
	}

	return providers
}

func Run(matrix [][]int, start, end domain.Point, algorithm string, debug bool) {
	maze := domain.NewMaze(matrix)
	utilz.PrintMatrix("Maze", maze)

	for _, al := range getSearchProviders(algorithm) {
		utilz.PrintTitleString(al.GetTitle())

		startMemStats := new(runtime.MemStats)
		runtime.ReadMemStats(startMemStats)

		startTime := time.Now()
		result := maze.SearchExit(start, end, al, debug)
		timeDuration := utilz.MeasureTime(startTime)

		endMemStats := new(runtime.MemStats)
		runtime.ReadMemStats(endMemStats)

		utilz.PrintParagraphString("Path", result)

		utilz.PrintParagraphString("Time took", timeDuration)
		utilz.PrintParagraphString("Memory used", fmt.Sprintf("%.2f%s", float64(endMemStats.TotalAlloc-startMemStats.TotalAlloc)/1024, "kb"))

		fmt.Println()
	}
}
