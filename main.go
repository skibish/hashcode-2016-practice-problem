package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

// inspired: http://www.dotnetperls.com/remove-duplicates-slice
func removeDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

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
		if squareSize%2 != 0 { // iteraeting only on odds
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

	commands := make(map[string][]string)
	cmdPaintSquare := `"PAINT_SQUARE %v %v %v"`
	cmdEraseCell := `"ERASE_CELL %v %v"`
	cmdPaintLine := `"PAINT_LINE %v %v %v %V"`

	// iterating windows, from max ~> min
	for _, sqSize := range sortedKeys {
		// shift with help of which we will find center of square.
		// S in PAINT_SQUARE command
		shift, _ := math.Modf(float64(sqSize) / 2)

		for i := len(windows[sqSize]) - 1; i > 0; i-- {
			// find start coordinates to draw
			centerX := windows[sqSize][i].x + int(shift)
			centerY := windows[sqSize][i].y + int(shift)

			// now we need to now if we need to paint
			var toPaint bool
			var sameSize int
			// loop over Y axis and look if there is ONE empty column
			for i := windows[sqSize][i].x; i < windows[sqSize][i].x+sqSize; i++ {
				for j := windows[sqSize][i].y; j < windows[sqSize][i].y+sqSize; j++ {
					if dataArr[j][i] == true {
						sameSize++
					}
				}

				if sameSize == sqSize {
					toPaint = true
					break
				}

				sameSize = 0
			}

			// if no column found, search for ROW
			if !toPaint {
				for i := windows[sqSize][i].y; i < windows[sqSize][i].y+sqSize; i++ {
					for j := windows[sqSize][i].x; j < windows[sqSize][i].x+sqSize; j++ {
						if dataArr[i][j] == true {
							sameSize++
						}
					}

					if sameSize == sqSize {
						toPaint = true
						break
					}

					sameSize = 0
				}
			}

			// Only if ROW or COLUMN is found - paint

			if toPaint {
				// Fill mask (area painted with squares)
				for maskX := 0; maskX < sqSize; maskX++ {
					for maskY := 0; maskY < sqSize; maskY++ {
						y := windows[sqSize][i].y + maskY
						x := windows[sqSize][i].x + maskX
						// on mask we paint added squares

						// if cell is empty in original, we need to erase it (remember command)
						if dataArr[y][x] == false {
							commands["erase_cell"] = append(commands["erase_cell"], fmt.Sprintf(cmdEraseCell, x, y))
						}
						// pushing square command to array
						commands["paint_square"] = append(commands["paint_square"], fmt.Sprintf(cmdPaintSquare, centerX, centerY, shift))

						maskMatrix[y][x] = true
						// on original we erase them
						dataArr[y][x] = false
					}
				}

			}

		}
	}

	// find optimal lines TODO: not finished, need to loop to fill all lines

	var foundColored bool
	foundColored = true
	for {
		// if no colored found, we finished :)
		if foundColored {
			foundColored = false
		} else {
			break
		}

		for x := 0; x < len(dataArr[0]); x++ {
			var startX int
			var startY int
			var endX int
			var endY int
			for y := 0; y < len(dataArr); y++ {
				if dataArr[y][x] == true {
					foundColored = true // mark, that we found colored

					startY = y
					startX = x

					for i := x; i < len(dataArr[0]); i++ {
						if dataArr[y][i] == false {
							break
						}

						endX = i
					}

					for j := y; j < len(dataArr); j++ {
						if dataArr[j][x] == false {
							break
						}

						endY = j
					}

				}

				sizeByX := endX - startX
				sizeByY := endY - startY

				if sizeByX >= sizeByY {
					commands["paint_line"] = append(commands["paint_line"], fmt.Sprintf(cmdPaintLine, startX, startY, endX, startY))

					for xx := 0; xx <= endX; xx++ {
						dataArr[startY][xx] = false
					}

				} else {
					commands["paint_line"] = append(commands["paint_line"], fmt.Sprintf(cmdPaintLine, startX, startY, startX, endY))

					for yy := 0; yy <= endY; yy++ {
						dataArr[yy][startX] = false
					}
				}
			}
		}

	}

	// clean duplicate commands
	for k, v := range commands {
		commands[k] = removeDuplicates(v)
	}

	// var total int
	// for _, k := range commands {
	// 	total = total + len(k)
	// }
	// fmt.Println(total)
}
