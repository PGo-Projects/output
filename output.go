package output

import (
	tm "github.com/buger/goterm"
)

const (
	BLACK = iota
	RED
	GREEN
	YELLOW
	BLUE
	MAGENTA
	CYAN
	WHITE
)

func Print(message string, color int) {
	tm.Print(tm.Color(message, color))
	tm.Flush()
}

func Println(message string, color int) {
	tm.Println(tm.Color(message, color))
	tm.Flush()
}
