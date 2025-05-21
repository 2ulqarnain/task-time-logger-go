package out

import (
	"fmt"
	"task-time-logger-go/utils/enums/TerminalColors"
)

func Errorln(error string) {
	fmt.Println(TerminalColors.Red + "Error: " + TerminalColors.Reset + error)
}

func Successln(message string) {
	fmt.Println(TerminalColors.Green + message + TerminalColors.Reset)
}

// Warning

func Warningln(message string) {
	fmt.Println(TerminalColors.Yellow + "Warning: " + TerminalColors.Reset + message)
}

func Warningf(message string, value string) {
	fmt.Printf(TerminalColors.Yellow+"Warning: "+TerminalColors.Reset+message, value)
}
