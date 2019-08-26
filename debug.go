package output

import (
	"fmt"

	tm "github.com/buger/goterm"
)

func DebugError(debugMode bool, err error) {
	if err == nil {
		return
	}
	tm.Print(tm.Color(fmt.Sprintf("[DEBUG] %s", err.Error()), tm.RED))
	tm.Flush()
}

func DebugErrorln(debugMode bool, err error) {
	if err == nil {
		return
	}
	tm.Println(tm.Color(fmt.Sprintf("[DEBUG] %s", err.Error()), tm.RED))
	tm.Flush()
}

func DebugString(debugMode bool, s string, color int) {
	tm.Print(tm.Color(fmt.Sprintf("[DEBUG] %s", s), color))
	tm.Flush()
}

func DebugStringln(debugMode bool, s string, color int) {
	tm.Println(tm.Color(fmt.Sprintf("[DEBUG] %s", s), color))
	tm.Flush()
}
