package internals

import "errors"

func GetSystemInfo(include string) (*SystemInfo, error) {

	systemInfo := &SystemInfo{}

	switch include {
	case "cpuinfo":
		cpuInfo, err := CPUParser()
		if err != nil {
			return nil, err
		}
		systemInfo.CPU = *cpuInfo
	case "meminfo":
		memInfo, err := MemParser()
		if err != nil {
			return nil, err
		}
		systemInfo.Memory = *memInfo
	case "loadavg":
		loadAvg, err := LoadParser()
		if err != nil {
			return nil, err
		}
		systemInfo.LoadAvg = *loadAvg
	case "all":
		cpuInfo, err := CPUParser()
		if err != nil {
			return nil, err
		}
		systemInfo.CPU = *cpuInfo

		memInfo, err := MemParser()
		if err != nil {
			return nil, err
		}
		systemInfo.Memory = *memInfo

		loadAvg, err := LoadParser()
		if err != nil {
			return nil, err
		}
		systemInfo.LoadAvg = *loadAvg
	default:
		return nil, errors.New("invalid source option")
	}

	return systemInfo, nil

}

