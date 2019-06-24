package output

import (
	"os"

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

func ErrorAndExit(err error, exitCode bool) {
	tm.Println(tm.Color(err.Error(), tm.RED))
	tm.Flush()
	os.Exit(exitCode)
}

func ErrorStringAndExit(err error, exitCode bool) {
	tm.Println(tm.Color(err, tm.RED))
	tm.Flush()
	os.Exit(exitCode)
}

func Success(message string) {
	tm.Println(tm.Color(message, tm.GREEN))
	tm.Flush()
}

func Println(message string, color int) {
	tm.Println(tm.Color(message, color))
	tm.Flush()
}
