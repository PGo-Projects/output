package output

import tm "github.com/buger/goterm"

func Error(err error) {
	if err == nil {
		return
	}
	tm.Print(tm.Color(err.Error(), tm.RED))
	tm.Flush()
}

func Errorln(err error) {
	if err == nil {
		return
	}
	tm.Println(tm.Color(err.Error(), tm.RED))
	tm.Flush()
}

func ErrorString(err string) {
	tm.Print(tm.Color(err, tm.RED))
	tm.Flush()
}

func ErrorStringln(err string) {
	tm.Println(tm.Color(err, tm.RED))
	tm.Flush()
}

func ErrorAndPanic(err error) {
	if err == nil {
		return
	}
	tm.Print(tm.Color(err.Error(), tm.RED))
	tm.Flush()
	panic(err)
}

func ErrorAndPanicln(err error) {
	if err == nil {
		return
	}
	tm.Println(tm.Color(err.Error(), tm.RED))
	tm.Flush()
	panic(err)
}

func ErrorStringAndPanic(err string) {
	tm.Print(tm.Color(err, tm.RED))
	tm.Flush()
	panic(err)
}

func ErrorStringAndPanicln(err string) {
	tm.Println(tm.Color(err, tm.RED))
	tm.Flush()
	panic(err)
}

func Success(message string) {
	tm.Print(tm.Color(message, tm.GREEN))
	tm.Flush()
}

func Successln(message string) {
	tm.Println(tm.Color(message, tm.GREEN))
	tm.Flush()
}
