package output

import (
	"fmt"
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

func DebugError(debugMode bool, err error) {
	tm.Println(tm.Color(fmt.Sprintf("[DEBUG] %s", err.Error()), tm.RED))
	tm.Flush()
}

func DebugString(debugMode bool, s string, color int) {
	tm.Println(tm.Color(fmt.Sprintf("[DEBUG] %s", s), color))
	tm.Flush()
}

func Error(err error) {
	tm.Println(tm.Color(err.Error(), tm.RED))
	tm.Flush()
}

func ErrorString(err string) {
	tm.Println(tm.Color(err, tm.RED))
	tm.Flush()
}

func ErrorAndExit(err error, exitCode int) {
	tm.Println(tm.Color(err.Error(), tm.RED))
	tm.Flush()
	os.Exit(exitCode)
}

func ErrorStringAndExit(err string, exitCode int) {
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
