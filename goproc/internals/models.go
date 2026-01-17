package internals


type CPUInfo struct {
    VendorId   string  `json:"vendor_id,omitempty"`
    CPUCount   int     `json:"cpu_count,omitempty"`   
    CPUFamily  string  `json:"cpu_family,omitempty"`
    ClockSpeed float64 `json:"clock_speed_mhz,omitempty"`
    ModelName  string  `json:"model_name,omitempty"`
    CoreCount  int     `json:"core_count,omitempty"`   
    CacheSize  int     `json:"cache_size_kb,omitempty"`
}

type MemInfo struct {
    Total     int64 `json:"total_kb,omitempty"`
    Free      int64 `json:"free_kb,omitempty"`
    Available int64 `json:"available_kb,omitempty"`
    Buffers   int64 `json:"buffers_kb,omitempty"`
    SwapTotal int64 `json:"swap_total_kb,omitempty"`
    SwapFree  int64 `json:"swap_free_kb,omitempty"`
}

type LoadAvg struct {
    OneMin         float64 `json:"load_1m,omitempty"`
    FiveMin        float64 `json:"load_5m,omitempty"`
    FifteenMin     float64 `json:"load_15m,omitempty"`
    ProcessCount   int     `json:"running_processes,omitempty"`
    TotalProcesses int     `json:"total_processes,omitempty"`
}

type SystemInfo struct {
    CPU     CPUInfo `json:"cpu"`
    Memory  MemInfo `json:"memory"`
    LoadAvg LoadAvg `json:"load_avg"`
}