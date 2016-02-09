package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

func getLargestSideSize(s, rows, cols int) int {
	c := 2 * (s - 1)
	side, _ := math.Modf(math.Sqrt(float64(c)))
	if int(side) > rows {
		side = float64(rows)
	}
	if int(side) > cols {
		side = float64(cols)
	}
	sideInt := int(side)
	for sideInt%2 == 0 {
		sideInt--
	}
	return sideInt
}

func generateMatrix(scanner *bufio.Scanner) (rows, cols, count int, dataArr [][]bool) {
	var lineNum int
	for scanner.Scan() {
		if lineNum > 0 {
			var lineSl []bool
			for _, v := range scanner.Text() {
				cols++
				switch string(v) {
				case ".":
					lineSl = append(lineSl, false)
				case "#":
					count++
					lineSl = append(lineSl, true)
				}
			}
			dataArr = append(dataArr, lineSl)
			rows++
		}
		lineNum++
	}
	return
}

type window struct {
	x        int
	y        int
	coloredN int
}

type sortWindows []window

func (s sortWindows) Less(i, j int) bool {
	return s[i].coloredN < s[j].coloredN
}

func (s sortWindows) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortWindows) Len() int {
	return len(s)
}

func main() {
	// Reading a file and generating the corresponding boolean matrix.
	file, err := os.Open("input/logo.in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Err() != nil {
		log.Fatal(err)
	}

	rows, cols, s, dataArr := generateMatrix(scanner)

	// Determining size of the largest square window.
	ls := getLargestSideSize(s, rows, cols)

	// Generating list of windows.
	windows := make(map[int][]window)
	minSquareSize := 3

	for squareSize := ls; squareSize >= minSquareSize; squareSize-- {
		for x := 0; x <= len(dataArr[0])-squareSize; x++ {
			for y := 0; y <= len(dataArr)-squareSize; y++ {
				windowData := window{
					x: x,
					y: y,
				}
				for i := x; i < x+squareSize; i++ {
					for j := y; j < y+squareSize; j++ {
						if dataArr[j][i] {
							windowData.coloredN++
						}
					}
				}
				if windowData.coloredN >= squareSize*squareSize-squareSize {
					windows[squareSize] = append(windows[squareSize], windowData)
				}
			}
		}
	}

	// Initializing mask matrix.
	maskMatrix := make([][]bool, len(dataArr))
	for i := range maskMatrix {
		maskMatrix[i] = make([]bool, len(dataArr[i]))
	}

	// Sorting windows data.
	var sortedKeys []int
	for k := range windows {
		sortedKeys = append(sortedKeys, k)
		sort.Sort(sortWindows(windows[k]))
	}

	// sort in order max ~> min
	sort.Sort(sort.Reverse(sort.IntSlice(sortedKeys)))

	// initial draft for iterating windows
	// for _, sqSize := range sortedKeys {
	// 	for i := len(windows[sqSize]) - 1; i > 0; i-- {
	// 		x := windows[sqSize][i].x
	// 		y := windows[sqSize][i].y
	//
	// 	}
	// }

	for _, k := range sortedKeys {
		fmt.Println(k)
	}
}
