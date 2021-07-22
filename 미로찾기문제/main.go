package main

import (
	"container/list"
	"fmt"
	"os"
)

const MAX_STACK_SIZE int = 100

type offsets struct {
	vert  int
	horiz int
}

type Stack struct {
	v *list.List
}

type element struct {
	dir int
	row int
	col int
}

const (
	N int = iota
	NE
	E
	SE
	S
	SW
	W
	NW
)

var move = [8]offsets{
	{-1, 0},  //N
	{-1, 1},  //NE
	{0, 1},   //E
	{1, 1},   //SE
	{1, 0},   //S
	{1, -1},  //SW
	{0, -1},  //W
	{-1, -1}, //NW
}

var maze [][]int
var mark [][]int
var stack [MAX_STACK_SIZE]element
var top = -1
var globalRow, globalCol int
var enRow, enCol, exRow, exCol int
var c rune

func nothing(v ...interface{}) {}

func Push(val element) {
	top++
	stack[top] = val
}

func Pop() element {
	if top != -1 {
		temp := stack[top]
		top--
		return temp
	} else {
		return stack[top]
	}
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func LoadMaze() {
	var filepath string = "C:\\Users\\jsviv\\OneDrive\\maze1.txt"
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error 001 : failed to open file")
		panic(err)
	}
	defer f.Close()
	fmt.Fscanf(f, "%d %d", &globalRow, &globalCol)

	for i := 0; i < globalRow; i++ {
		for j := 0; j < globalCol; j++ {
			fmt.Fscanf(f, "%d", maze[i][j])
		}
	}
	globalRow += 2
	globalCol += 2
	fmt.Println(globalRow, globalCol)
	maze = make([][]int, globalRow)
	mark = make([][]int, globalRow)
	for i := 0; i < globalRow; i++ {
		maze[i] = make([]int, globalCol)
		mark[i] = make([]int, globalCol)
	}

	var temp int
	for i := 0; i < globalRow; i++ {
		for j := 0; j < globalCol; j++ {
			if (i == 0) || (i == (globalRow - 1)) || (j == 0) || (j == (globalCol - 1)) {
				maze[i][j] = 1
			} else {

				fmt.Fscanf(f, "%d", &temp)

				fmt.Print(" ", temp, " ")
				maze[i][j] = temp
			}
		}
		fmt.Fscanf(f, "\n")

	}

	fmt.Fscanf(f, "%d%d %d %d", &enRow, &enCol, &exRow, &exCol)
	fmt.Print(enRow, " ", enCol, " ", exRow, " ", exCol, "\n")
	enRow++
	enCol++
	exRow++
	exCol++

}

func PrintMaze() {
	for i := 0; i < globalRow; i++ {
		for j := 0; j < globalCol; j++ {
			fmt.Printf("%3d", maze[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func PrintResult() {
	for top > -1 {
		maze[stack[top].row][stack[top].col] = 2
		top--
	}
	maze[enRow][enCol] = 3
	maze[exRow][exCol] = 4

	for i := 0; i < globalRow; i++ {
		for j := 0; j < globalCol; j++ {
			if maze[i][j] == 2 {
				fmt.Printf("%3c", 'X')
			} else if maze[i][j] == 3 {
				fmt.Printf("%3c", 'S')
			} else if maze[i][j] == 4 {
				fmt.Printf("%3c", 'F')
			} else {
				fmt.Printf("%3d", maze[i][j])
			}
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}

func Path() {
	fmt.Printf("Maze with 1 boundaries\n\n")
	PrintMaze()
	var row, col, nextRow, nextCol int
	var dir int
	var found bool = false
	var position element

	mark[enRow][enCol] = 1
	top = 0
	stack[0].row = enRow
	stack[0].col = enCol
	stack[0].dir = N

	for top > -1 && !found {
		position = Pop()
		row = position.row
		col = position.col
		dir = position.dir

		for dir < 8 && !found {
			nextRow = row + move[dir].vert
			nextCol = col + move[dir].horiz

			if nextRow == exRow && nextCol == exCol {
				found = true
			} else if maze[nextRow][nextCol] == 0 && mark[nextRow][nextCol] == 0 {
				mark[nextRow][nextCol] = 1
				position.row = row
				position.col = col
				position.dir = dir
				Push(position)
				row = nextRow
				col = nextCol
				dir = N
			} else {
				dir++
			}

		}
	}
	position.row = row
	position.col = col
	Push(position)

	if found {
		fmt.Printf("The path is : \n")
		PrintResult()
	} else {
		PrintMaze()
		fmt.Println("The maze does not have any path!")
	}

}

func main() {
	LoadMaze()
	Path()

}
