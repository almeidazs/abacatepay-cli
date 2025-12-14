package logger

import (
	"fmt"
	"os"
	"time"
	"unicode/utf8"

	"github.com/fatih/color"
)

const maxContentSize = 75

var SuccessPrefix = color.New(color.FgHiGreen).Sprint("AbacatePay")

func Success(msg string, args ...any) {
	fmt.Printf(
		"%s (%s): %s\n",
		SuccessPrefix,
		timestamp(),
		truncate(fmt.Sprintf(msg, args...)),
	)
}

var ErrorPrefix = color.New(color.FgHiRed).Sprint("AbacatePay")

func Error(err error, args ...any) {
	if err == nil {
		return
	}

	fmt.Fprintf(
		os.Stderr,
		"%s (%s): %v\n",
		ErrorPrefix,
		timestamp(),
		err,
	)
}
func truncate(str string) string {
	if utf8.RuneCountInString(str) <= maxContentSize {
		return str
	}

	runes := []rune(str)

	return string(runes[:maxContentSize]) + "..."
}

func timestamp() string {
	return time.Now().Format("15:04:05")
}
