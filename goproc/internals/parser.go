package internals

import (
	"bufio"
	"fmt"
	"strings"
)

var cpuInfo CPUInfo
var memInfo MemInfo
var loadAvg LoadAvg
var completedFields = 0

func CPUParser() *CPUInfo {

	completedFields = 0

	reader, err := GetSource("cpuinfo")

	if err != nil {
		fmt.Println("[ERROR] :", err)
	}

	defer reader.Close()

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		if completedFields < 8 {
			line := scanner.Text()
			parts := strings.SplitN(line, ":", 2)

			if len(parts) < 2 {
				continue
			}

			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			switch key {
			case "vendor_id":
				cpuInfo.VendorId = value
				completedFields++
			case "cpu family":
				cpuInfo.CPUFamily = value
				completedFields++
			case "model name":
				cpuInfo.ModelName = value
				completedFields++
			case "cpu MHz":
				fmt.Sscanf(value, "%f", &cpuInfo.ClockSpeed)
				completedFields++
			case "cpu cores":
				fmt.Sscanf(value, "%d", &cpuInfo.CoreCount)
				completedFields++
			case "cache size":
				fmt.Sscanf(value, "%d", &cpuInfo.CacheSize)
				completedFields++
			case "siblings":
				fmt.Sscanf(value, "%d", &cpuInfo.CPUCount)
				completedFields++
			}
		}
	}

	return &cpuInfo
}

func MemParser() *MemInfo {

	completedFields = 0
	reader, err := GetSource("meminfo")

	if err != nil {
		fmt.Println("[ERROR] :", err)
	}

	defer reader.Close()

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		if completedFields < 6 {
			line := scanner.Text()
			parts := strings.SplitN(line, ":", 2)

			if len(parts) < 2 {
				continue
			}

			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			switch key {
			case "MemTotal":
				fmt.Sscanf(value, "%d", &memInfo.Total)
				completedFields++
			case "MemFree":
				fmt.Sscanf(value, "%d", &memInfo.Free)
				completedFields++
			case "MemAvailable":
				fmt.Sscanf(value, "%d", &memInfo.Available)
				completedFields++
			case "Buffers":
				fmt.Sscanf(value, "%d", &memInfo.Buffers)
				completedFields++
			case "SwapTotal":
				fmt.Sscanf(value, "%d", &memInfo.SwapTotal)
				completedFields++
			case "SwapFree":
				fmt.Sscanf(value, "%d", &memInfo.SwapFree)
				completedFields++
			}
		}
	}

	return &memInfo
}

func LoadParser() *LoadAvg {

	reader, err := GetSource("loadavg")

	if err != nil {
		fmt.Println("[ERROR] :", err)
	}

	defer reader.Close()

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) < 5 {
			continue
		}

		fmt.Sscanf(parts[0], "%f", &loadAvg.OneMin)
		fmt.Sscanf(parts[1], "%f", &loadAvg.FiveMin)
		fmt.Sscanf(parts[2], "%f", &loadAvg.FifteenMin)

		procParts := strings.SplitN(parts[3], "/", 2)
		if len(procParts) == 2 {
			fmt.Sscanf(procParts[0], "%d", &loadAvg.ProcessCount)
			fmt.Sscanf(procParts[1], "%d", &loadAvg.TotalProcesses)
		}
	}

	return &loadAvg
}
