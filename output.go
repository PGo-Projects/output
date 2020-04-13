package output

import (
	"fmt"

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

func DefaultPrint(message string) {
	fmt.Print(message)
}

func DefaultPrintf(message string, args ...interface{}) {
	fmt.Printf(message, args...)
}

func DefaultPrintln(message string) {
	fmt.Println(message)
}

func Print(message string, color int) {
	tm.Print(tm.Color(message, color))
	tm.Flush()
}

func Printf(color int, message string, args ...interface{}) {
	tm.Print(tm.Color(fmt.Sprintf(message, args...), color))
	tm.Flush()
}

func Printfln(color int, message string, args ...interface{}) {
	tm.Println(tm.Color(fmt.Sprintf(message, args...), color))
	tm.Flush()
}

func Println(message string, color int) {
	tm.Println(tm.Color(message, color))
	tm.Flush()
}
