package main

import (
	"fmt"
	"github.com/Nav1Cr0ss/algorithms-lab2/internal/app"
	"github.com/Nav1Cr0ss/algorithms-lab2/internal/domain"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestFindExit(t *testing.T) {

	tests := []struct {
		name         string
		algorithm    domain.SearchAlgorithm
		start        domain.Point
		end          domain.Point
		matrix       [][]int
		expectedPath string
	}{
		{
			name:      "A*: Success Search Exit in 20x20 maze",
			algorithm: app.NewAStar(),
			start:     domain.Point{X: 0, Y: 0},
			end:       domain.Point{X: 19, Y: 19},
			matrix: [][]int{
				{0, 0, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0},
				{1, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
				{1, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1},
				{1, 1, 0, 0, 0, 1, 0, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
				{1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1},
				{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1, 1},
				{1, 1, 1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 0, 1, 1, 1, 0},
				{1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0, 0, 1, 0, 1, 1, 1, 1, 1, 0},
				{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0},
				{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 0, 1},
				{1, 0, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1},
				{0, 0, 1, 0, 0, 0, 1, 0, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 0, 1},
				{1, 1, 1, 1, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
				{1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1},
				{1, 1, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0},
			},
			expectedPath: "[{0 0} {0 1} {1 1} {1 2} {2 2} {3 2} {3 3} {3 4} {2 4} {1 4} {1 5} {1 6} {2 6} {3 6} {4 6} {5 6} {5 5} {5 4} {5 3} {5 2} {6 2} {7 2} {7 3} {8 3} {9 3} {9 4} {9 5} {8 5} {7 5} {7 6} {7 7} {7 8} {7 9} {8 9} {9 9} {9 10} {9 11} {10 11} {11 11} {12 11} {12 10} {12 9} {12 8} {12 7} {13 7} {14 7} {14 8} {14 9} {15 9} {16 9} {16 10} {16 11} {16 12} {16 13} {16 14} {16 15} {16 16} {17 16} {18 16} {18 17} {18 18} {18 19} {19 19}]",
		},
		{
			name:      "A*: Success Search Exit in 5x5 maze without exit",
			algorithm: app.NewAStar(),
			start:     domain.Point{X: 0, Y: 0},
			end:       domain.Point{X: 4, Y: 4},
			matrix: [][]int{
				{0, 0, 1, 1, 1},
				{1, 0, 0, 1, 0},
				{1, 1, 0, 1, 0},
				{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 1},
			},
			expectedPath: "[]",
		},
		{
			name:      "BFS: Success Search Exit in 5x5 maze without exit",
			algorithm: app.NewBFS(),
			start:     domain.Point{X: 0, Y: 0},
			end:       domain.Point{X: 4, Y: 4},
			matrix: [][]int{
				{0, 0, 1, 1, 1},
				{1, 0, 0, 1, 0},
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 0},
			},
			expectedPath: "[{0 0} {0 1} {1 1} {1 2} {2 2} {3 2} {3 3} {3 4} {4 4}]",
		},
		{
			name:      "LDFS: Success Search Exit in 5x5 maze without exit",
			algorithm: app.NewLDFS(),
			start:     domain.Point{X: 0, Y: 0},
			end:       domain.Point{X: 4, Y: 4},
			matrix: [][]int{
				{0, 0, 1, 1, 1},
				{1, 0, 0, 1, 0},
				{1, 1, 0, 1, 0},
				{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 0},
			},
			expectedPath: "[{0 0} {0 1} {1 1} {1 2} {2 2} {3 2} {3 3} {3 4} {4 4}]",
		},
		{
			name:      "A*: Success Search Exit in 5x5 maze without exit",
			algorithm: app.NewAStar(),
			start:     domain.Point{X: 0, Y: 0},
			end:       domain.Point{X: 4, Y: 4},
			matrix: [][]int{
				{0, 0, 1, 1, 1},
				{1, 0, 0, 1, 0},
				{1, 1, 0, 1, 0},
				{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 0},
			},
			expectedPath: "[{0 0} {0 1} {1 1} {1 2} {2 2} {3 2} {3 3} {3 4} {4 4}]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maze := domain.NewMaze(tt.matrix)
			result := maze.SearchExit(tt.start, tt.end, tt.algorithm, false)

			assert.Equal(t, tt.expectedPath, fmt.Sprintf("%v", result))

		})
	}
}
