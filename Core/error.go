package core

import (
	"fmt"
	"os"
)

func ReportError(errorType string, message string) {
	fmt.Printf("%sError: %s\n", errorType, message)
	os.Exit(1)
}
