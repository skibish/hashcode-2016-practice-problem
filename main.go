package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

// Downloads/hashcode/test_round

func check(e error) {
	if e != nil {
		panic(e)
	}
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
			rows++
		}
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
			dataArr = append(dataArr, lineSl)
		}
		lineNum++
	}
	return
}

type coloredIndexes struct {
	x       int
	y       int
	colored int
}

func main() {
	file, err := os.Open("/Users/sk/Downloads/hashcode/test_round/logo.in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// s, dataArr := generateMatrix(scanner)
	rows, cols, s, dataArr := generateMatrix(scanner)

	if scanner.Err() != nil {
		log.Fatal(err)
	}

	ls := getLargestSideSize(s, rows, cols)
	fmt.Println("===>>", ls)

	maskMatrix := make([][]int, len(dataArr))
	for i := 0; i < len(maskMatrix); i++ {
		maskMatrix[i] = make([]int, len(dataArr[i]))
	}

	windows := make(map[int]coloredIndexes)
	currentSize := ls
	var coloredCount, x, y int
	for {
		var colored int
		for i := coloredCount; i < ls; i++ {
			for j := coloredCount; j < ls; j++ {
				if dataArr[i][j] {
					colored++
				}
			}
		}
		windows[currentSize] = coloredIndexes{
			x:       x,
			y:       y,
			colored: colored,
		}
	}
git remote add origin git@github.com:skibish/round1.git
	/*for _, v := range dataArr {
		fmt.Println(v)
	}*/

}
