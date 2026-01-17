package internals


type CPUInfo struct {
    VendorId   string  `json:"vendor_id"`
    CPUCount   int     `json:"cpu_count"`   
    CPUFamily  string  `json:"cpu_family"`
    ClockSpeed float64 `json:"clock_speed_mhz"`
    ModelName  string  `json:"model_name"`
    CoreCount  int     `json:"core_count"`   
    CacheSize  int     `json:"cache_size_kb"`
}

type MemInfo struct {
    Total     int64 `json:"total_kb"`
    Free      int64 `json:"free_kb"`
    Available int64 `json:"available_kb"`
    Buffers   int64 `json:"buffers_kb"`
    SwapTotal int64 `json:"swap_total_kb"`
    SwapFree  int64 `json:"swap_free_kb"`
}

type LoadAvg struct {
    OneMin         float64 `json:"load_1m"`
    FiveMin        float64 `json:"load_5m"`
    FifteenMin     float64 `json:"load_15m"`
    ProcessCount   int     `json:"running_processes"`
    TotalProcesses int     `json:"total_processes"`
}

type SystemInfo struct {
    CPU     CPUInfo `json:"cpu"`
    Memory  MemInfo `json:"memory"`
    LoadAvg LoadAvg `json:"load_avg"`
}