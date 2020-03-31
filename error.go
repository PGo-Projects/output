package output

func Error(err error) {
	if err == nil {
		return
	}

	Print(err.Error(), RED)
}

func Errorln(err error) {
	if err == nil {
		return
	}
	Println(err.Error(), RED)
}

func ErrorString(err string) {
	Print(err, RED)
}

func ErrorStringln(err string) {
	Println(err, RED)
}

func ErrorAndPanic(err error) {
	if err == nil {
		return
	}
	Print(err.Error(), RED)
	panic(err)
}

func ErrorAndPanicln(err error) {
	if err == nil {
		return
	}
	Println(err.Error(), RED)
	panic(err)
}

func ErrorStringAndPanic(err string) {
	Print(err, RED)
	panic(err)
}

func ErrorStringAndPanicln(err string) {
	Println(err, RED)
	panic(err)
}

func Success(message string) {
	Print(message, GREEN)
}

func Successln(message string) {
	Println(message, GREEN)
}
