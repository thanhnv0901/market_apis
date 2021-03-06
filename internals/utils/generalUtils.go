package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"market_apis/internals/errorstrack"
	"path"
	"runtime"
)

// GetDirCurrent ..
func GetDirCurrent() (string, int, bool) {
	// Get current directory path
	_, filename, line, isOk := runtime.Caller(1)
	pathDir := path.Dir(filename)

	return pathDir, line, isOk
}

// ErrorTrackingDefer ..
func ErrorTrackingDefer() {
	if err := recover(); err != nil {
		msg := fmt.Sprintf("%s", err)
		errTrack := errorstrack.ErrorsTrack{Message: msg}
		log.Println(errTrack.PrintTrackingJSON())
	}
}

// ToJSON ..
func ToJSON(obj interface{}) string {
	data, _ := json.Marshal(obj)
	return string(data)
}
