# goproc

`goproc` is a lightweight, high-performance CLI tool written in Go for monitoring system health on Linux. It parses kernel data from `/proc` and outputs detailed system statistics in a structured JSON format, making it perfect for automation, logging, or integration with other tools.

## 🚀 Features

* **Fast & Efficient**: Directly parses `/proc/cpuinfo`, `/proc/meminfo`, and `/proc/loadavg`.
* **JSON Native**: Designed for modern workflows—all output is valid, indented JSON.
* **Selective Querying**: Use flags to fetch only the specific hardware data you need.
* **Zero Dependencies**: Once compiled, it's a single static binary.

## 📦 Installation

Since `goproc` uses Go modules, you can install it directly from GitHub:

```bash
go install github.com/Shreehari-Acharya/go-learnings/goproc@latest

```

*Note: Ensure your `$GOPATH/bin` (typically `~/go/bin`) is in your system's `PATH`.*

## 🛠 Usage

By default, running `goproc` will return a complete snapshot of your CPU, Memory, and Load Average.

```bash
goproc

```

### Filtering Output

Use the `--source` (or `-s`) flag to get specific data. This is faster as it skips parsing the files you don't need.

| Command | Description |
| --- | --- |
| `goproc -s cpuinfo` | Returns only CPU vendor, clock speed, cores, etc. |
| `goproc -s meminfo` | Returns RAM and Swap statistics. |
| `goproc -s loadavg` | Returns 1, 5, and 15-minute load averages and process counts. |
| `goproc -s all` | (Default) Returns all available system data. |

### Integration Examples

Since the output is JSON, you can easily pipe it into `jq` for filtering:

```bash
# Get only the CPU model name
goproc -s cpuinfo | jq '.cpu.model_name'

# Check if available memory is below 2GB
goproc -s meminfo | jq '.memory.available_kb < 2000000'

```

## 📊 Output Schema

The tool generates a JSON object with the following structure:

```json
{
  "cpu": {
    "vendor_id": "AuthenticAMD",
    "cpu_count": 8,
    "cpu_family": "23",
    "clock_speed_mhz": 1345.529,
    "model_name": "AMD Ryzen 5 3500U...",
    "core_count": 4,
    "cache_size_kb": 512
  },
  "memory": {
    "total_kb": 30696596,
    "free_kb": 22071352,
    "available_kb": 26670228,
    "buffers_kb": 7152,
    "swap_total_kb": 8388604,
    "swap_free_kb": 8388604
  },
  "load_avg": {
    "load_1m": 0.45,
    "load_5m": 0.52,
    "load_15m": 0.48,
    "running_processes": 2,
    "total_processes": 845
  }
}

```

