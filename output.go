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

func Error(err error) {
	tm.Println(tm.Color(err.Error(), tm.RED))
	tm.Flush()
}

func ErrorString(err string) {
	tm.Println(tm.Color(err, tm.RED))
	tm.Flush()
}

func Success(message string) {
	tm.Println(tm.Color(message, tm.GREEN))
	tm.Flush()
}

func Println(message string, color int) {
	tm.Println(tm.Color(message, color))
	tm.Flush()
}
