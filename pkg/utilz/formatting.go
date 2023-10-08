package utilz

import (
	"fmt"
	"strings"
)

func PrintTitleString(title string) {
	line := strings.Repeat("=", 10)
	fmt.Println(fmt.Sprintf("%s %s %s", line, title, line))

}

func PrintParagraphString(title string, body any) {
	fmt.Printf("%s: %v\n", title, body)
}

func PrintMatrix(title string, matrix [][]int) {
	fmt.Printf("%s: ", title)
	for i, row := range matrix {

		space := len(title) + 2
		if i == 0 {
			space = 0
		}

		fmt.Printf("%s", strings.Repeat(" ", space))
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
	fmt.Println()
}
