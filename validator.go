package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	w := 80
	h := 15

	file, err := os.Open("result.out")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Err() != nil {
		log.Fatal(err)
	}

	matrix := make([][]string, h)
	for i := range matrix {
		matrix[i] = make([]string, w)
	}

	for scanner.Scan() {
		command := strings.Split(scanner.Text(), " ")
		switch command[0] {
		case "PAINT_SQUARE":
			if len(command) != 4 {
				log.Fatal("invalid input for square")
			}

			inputY, err := strconv.Atoi(command[1])
			if err != nil {
				log.Fatal(err)
			}
			inputX, err := strconv.Atoi(command[2])
			if err != nil {
				log.Fatal(err)
			}
			inputS, err := strconv.Atoi(command[3])
			if err != nil {
				log.Fatal(err)
			}

			startX := inputX - inputS
			endX := inputX + inputS
			startY := inputY - inputS
			endY := inputY + inputS

			for i := startX; i <= endX; i++ {
				for j := startY; j <= endY; j++ {
					matrix[j][i] = "#"
				}
			}

		case "PAINT_LINE":
			if len(command) != 5 {
				log.Fatal("invalid input for line")
			}

			inputY0, err := strconv.Atoi(command[1])
			if err != nil {
				log.Fatal(err)
			}
			inputX0, err := strconv.Atoi(command[2])
			if err != nil {
				log.Fatal(err)
			}
			inputY1, err := strconv.Atoi(command[3])
			if err != nil {
				log.Fatal(err)
			}
			inputX1, err := strconv.Atoi(command[4])
			if err != nil {
				log.Fatal(err)
			}

			if inputX0 == inputX1 {
				for j := inputY0; j <= inputY1; j++ {
					matrix[j][inputX0] = "#"
				}
			} else if inputY0 == inputY1 {
				for i := inputY0; i <= inputY1; i++ {
					matrix[inputY0][i] = "#"
				}
			} else {
				log.Fatal("invalid input for line (it's not straight)")
			}

		case "ERASE_CELL":
			if len(command) != 3 {
				log.Fatal("invalid input for eracing")
			}

			inputY, err := strconv.Atoi(command[1])
			if err != nil {
				log.Fatal(err)
			}
			inputX, err := strconv.Atoi(command[2])
			if err != nil {
				log.Fatal(err)
			}

			matrix[inputY][inputX] = "."
		}
	}

	for i := 0; i < len(matrix[0]); i++ {
		for j := 0; j < len(matrix); j++ {
			if matrix[j][i] != "#" {
				matrix[j][i] = "."
			}
		}
	}

	for _, vy := range matrix {
		for _, vx := range vy {
			fmt.Print(vx)
		}
		fmt.Print("\n")
	}
}
