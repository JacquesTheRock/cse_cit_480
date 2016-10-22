package util

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func PrintError(a ...interface{}) {
	if ConvertToLevelID("ERROR") < Config.GetLogLevel() {
		return
	}
	fmtString := strings.Replace(Config.ErrorFmt, "${time}", (time.Now()).Format(Config.TimeFmt), 1)
	fmtString = strings.Replace(fmtString, "${level}", "ERROR", 1)
	fmtString = strings.Replace(fmtString, "${msg}", "%s\n", 1)
	fmt.Fprintf(os.Stderr, fmtString, a...)
}
func PrintDebug(a ...interface{}) {
	if ConvertToLevelID("DEBUG") < Config.GetLogLevel() {
		return
	}
	fmtString := strings.Replace(Config.ErrorFmt, "${time}", (time.Now()).Format(Config.TimeFmt), 1)
	fmtString = strings.Replace(fmtString, "${level}", "DEBUG", 1)
	fmtString = strings.Replace(fmtString, "${msg}", "%s\n", 1)
	fmt.Fprintf(os.Stderr, fmtString, a...)
}
func PrintWarn(a ...interface{}) {
	if ConvertToLevelID("WARN") < Config.GetLogLevel() {
		return
	}
	fmtString := strings.Replace(Config.ErrorFmt, "${time}", (time.Now()).Format(Config.TimeFmt), 1)
	fmtString = strings.Replace(fmtString, "${level}", "WARN", 1)
	fmtString = strings.Replace(fmtString, "${msg}", "%s\n", 1)
	fmt.Fprintf(os.Stderr, fmtString, a...)
}
func PrintInfo(a ...interface{}) {
	if ConvertToLevelID("INFO") < Config.GetLogLevel() {
		return
	}
	fmtString := strings.Replace(Config.ErrorFmt, "${time}", (time.Now()).Format(Config.TimeFmt), 1)
	fmtString = strings.Replace(fmtString, "${level}", "INFO", 1)
	fmtString = strings.Replace(fmtString, "${msg}", "%s\n", 1)
	fmt.Fprintf(os.Stdout, fmtString, a...)
}
