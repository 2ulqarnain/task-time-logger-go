package logger

import (
	"log"
	"os"
)

const (
	ColorGray  = "\033[90m"
	ColorReset = "\033[0m"

	// Dark Colors (Regular)
	ColorDarkRed    = "\033[31m"
	ColorDarkBlue   = "\033[34m"
	ColorDarkGreen  = "\033[32m"
	ColorDarkYellow = "\033[33m"
	ColorDarkPurple = "\033[35m"
	ColorDarkCyan   = "\033[36m"
	ColorDarkWhite  = "\033[37m"

	// Light Colors (Bright)
	ColorLightRed    = "\033[91m"
	ColorLightBlue   = "\033[94m"
	ColorLightGreen  = "\033[92m"
	ColorLightYellow = "\033[93m"
	ColorLightPurple = "\033[95m"
	ColorLightCyan   = "\033[96m"
	ColorLightWhite  = "\033[97m"
)

var (
	AppLogger = log.New(os.Stdout, ColorGray, log.LstdFlags|log.Lshortfile)
)
