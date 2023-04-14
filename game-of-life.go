package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	height = 40
)

type Field struct {
	s    [][]bool
	w, h int
}

func NewField() *Field {
	s := make([][]bool, height)
	for i := range s {
		s[i] = make([]bool, width)
	}
	return &Field{s: s, w: width, h: height}
}

func (f *Field) Set(x, y int, b bool) {
	f.s[y][x] = b
}

func (f *Field) Alive(x, y int) bool {
	x += f.w
	x %= f.w
	y += f.h
	y %= f.h
	return f.s[y][x]
}

func (f *Field) Next(x, y int) bool {
	alive := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (j != 0 || i != 0) && f.Alive(x+i, y+j) {
				alive++
			}
		}
	}
	return alive == 3 || alive == 2 && f.Alive(x, y)
}

func (f *Field) Show() {
	for i := 0; i < f.h; i++ {
		for j := 0; j < f.w; j++ {
			if f.Alive(j, i) {
				// Print blue character with black background
				fmt.Print("\033[34;40m") // Set text color to blue and background color to black
				fmt.Print("■")             // Print a block character as the colored cell
				fmt.Print("\033[0m")       // Reset text color and background color
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	field := NewField()
	for i := 0; i < (field.w * field.h / 4); i++ {
		field.Set(rand.Intn(field.w), rand.Intn(field.h), true)
	}

	for {
		fmt.Print("\033[2J") // ANSI escape code to clear the screen
		fmt.Print("\033[H")  // ANSI escape code to set the cursor position to the top-left corner
		field.Show()
		time.Sleep(time.Second / 30)

		field2 := NewField()
		for y := 0; y < field.h; y++ {
			for x := 0; x < field.w; x++ {
				field2.Set(x, y, field.Next(x, y))
			}
		}
		field, field2 = field2, field
	}
}

