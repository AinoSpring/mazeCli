package maze

import (
	"log"
	"os"
)

type Display struct {
	maze [][]byte
	size Vector
}

func DisplayMaze(maze Maze) (display Display) {
	display = NewDisplay(maze.Size)
	for x := 0; x < int(maze.Size[0]); x++ {
		for y := 0; y < int(maze.Size[0]); y++ {
			toSet := -1
			switch maze.Get(Vector{float64(x), float64(y)}) {
			case 0:
				toSet = '.'
				break
			case 1:
				toSet = '#'
				break
			}
			display.SetValue(Vector{float64(x), float64(y)}, byte(toSet))
		}
	}
	return
}

func NewDisplay(size Vector) (display Display) {
	display.size = size
	display.maze = make([][]byte, int(size[0]))
	for i := 0; i < int(size[0]); i++ {
		display.maze[i] = make([]byte, int(size[1]))
	}
	for y := 0; y < int(display.size[0]); y++ {
		for x := 0; x < int(display.size[1]); x++ {
			display.maze[y][x] = '#'
		}
	}
	return
}

func (display *Display) Save(path string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	for y := 0; y < int(display.size[0]); y++ {
		f.Write(display.maze[y])
		f.WriteString("\n")
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func (display *Display) SetValue(position Vector, value byte) {
	display.maze[int(position[0])][int(position[1])] = value
}

func (display *Display) Matrix(matrix Matrix, value byte) {
	for _, vector := range matrix {
		display.SetValue(vector, value)
	}
}
