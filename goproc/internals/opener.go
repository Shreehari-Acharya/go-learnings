package internals

import (
	"errors"
	"io"
	"os"
)


var fileMap = map[string]string{
	"meminfo": "/proc/meminfo",
	"cpuinfo": "/proc/cpuinfo",
	"loadavg": "/proc/loadavg",
}

// GetSource opens the specified source file and returns a ReadCloser
// If the source is invalid, it returns an error
// Supported sources are "meminfo" , "cpuinfo", "loadavg"
func GetSource(filename string) (io.ReadCloser, error) {
	
	path, exists := fileMap[filename]

	if !exists {
		return nil, errors.New("invalid source file")
	} 

	return os.Open(path)
}