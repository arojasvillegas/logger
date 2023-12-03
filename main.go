package logger

import (
	"fmt"
)

var Version string = "1.0"

func Log(mess string) {
	fmt.Println("[LOG] " + mess)
}
func Happy() {
	fmt.Println("8" + "=" + "D")
}
