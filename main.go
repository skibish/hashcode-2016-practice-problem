package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
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

	for squareSize := ls; squareSize > 0; squareSize-- {
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
				windows[squareSize] = append(windows[squareSize], windowData)
			}
		}
	}

	for windowSize, windowData := range windows {
		fmt.Println(windowSize, windowData)
	}
}
