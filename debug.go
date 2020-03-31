package output

var DEBUG = false

func DebugError(err error) {
	if DEBUG {
		if err == nil {
			Print("[DEBUG] No error\n", GREEN)
			return
		}

		Printf(RED, "[DEBUG] %s", err.Error())
	}
}

func DebugErrorln(err error) {
	if DEBUG {
		if err == nil {
			Println("[DEBUG] No error", GREEN)
			return
		}

		Printf(RED, "[DEBUG] %s\n", err.Error())
	}
}

func DebugString(s string, color int) {
	if DEBUG {
		Printf(color, "[DEBUG] %s", s)
	}
}

func DebugStringln(s string, color int) {
	if DEBUG {
		Printf(color, "[DEBUG] %s\n", s)
	}
}
