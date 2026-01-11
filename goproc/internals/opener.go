package internals

import (
	"errors"
	"fmt"
	"io"
	"os"
)


var fileMap = map[string]string{
	"meminfo": "/proc/meminfo",
	"cpuinfo": "/proc/cpuinfo",
}

func GetSource(filename string) (io.ReadCloser, error) {
	if(filename != "meminfo" && filename != "cpuinfo" && filename != "loadavg") {
		return nil,  fmt.Errorf("invalid source file")
	}

	path, exists := fileMap[filename]
	if !exists {
		return os.Open(path)
	} else {
		return nil, errors.New("invalid source file")
	}
}