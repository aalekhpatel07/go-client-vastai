package data_types

import "encoding/json"

type InstanceType int

const (
	OnDemand InstanceType = iota
	Reserved
	Bid
)

func (t InstanceType) MarshalJSON() ([]byte, error) {
	var s string
	switch t {
	default:
		s = ""
	case OnDemand:
		s = "on-demand"
	case Bid:
		s = "bid"
	case Reserved:
		s = "reserved"
	}
	return json.Marshal(s)
}

type Instance struct {
	AskContractId      int     `json:"ask_contract_id"`
	BundleId           int     `json:"bundle_id"`
	BundledResults     int     `json:"bundled_results"`
	BwNvlink           float64 `json:"bw_nvlink"`
	ComputeCap         int     `json:"compute_cap"`
	CpuArch            string  `json:"cpu_arch"`
	CpuCores           int     `json:"cpu_cores"`
	CpuCoresEffective  float64 `json:"cpu_cores_effective"`
	CpuGhz             float64 `json:"cpu_ghz"`
	CpuName            string  `json:"cpu_name"`
	CpuRam             int     `json:"cpu_ram"`
	CreditDiscountMax  float64 `json:"credit_discount_max"`
	CudaMaxGood        float64 `json:"cuda_max_good"`
	DirectPortCount    int     `json:"direct_port_count"`
	DiscountRate       float64 `json:"discount_rate"`
	DiscountedDphTotal float64 `json:"discounted_dph_total"`
	DiscountedHourly   float64 `json:"discounted_hourly"`
	DiskBw             float64 `json:"disk_bw"`
	DiskName           string  `json:"disk_name"`
	DiskSpace          float64 `json:"disk_space"`
	Dlperf             float64 `json:"dlperf"`
	DlperfPerDphtotal  float64 `json:"dlperf_per_dphtotal"`
	DphBase            float64 `json:"dph_base"`
	DphTotal           float64 `json:"dph_total"`
	DriverVers         int     `json:"driver_vers"`
	DriverVersion      string  `json:"driver_version"`
	Duration           float64 `json:"duration"`
	EndDate            float64 `json:"end_date"`
	External           string  `json:"external"`
	FlopsPerDphtotal   float64 `json:"flops_per_dphtotal"`
	Geolocation        string  `json:"geolocation"`
	Geolocode          int     `json:"geolocode"`
	GpuArch            string  `json:"gpu_arch"`
	GpuDisplayActive   bool    `json:"gpu_display_active"`
	GpuFrac            float64 `json:"gpu_frac"`
	GpuIds             []int   `json:"gpu_ids"`
	GpuLanes           int     `json:"gpu_lanes"`
	GpuMemBw           float64 `json:"gpu_mem_bw"`
	GpuName            string  `json:"gpu_name"`
	GpuRam             int     `json:"gpu_ram"`
	GpuTotalRam        int     `json:"gpu_total_ram"`
	HasAvx             int     `json:"has_avx"`
	HostId             int     `json:"host_id"`
	HostingType        int     `json:"hosting_type"`
	Hostname           string  `json:"hostname"`
	Id                 int     `json:"id"`
	InetDown           float64 `json:"inet_down"`
	InetDownCost       float64 `json:"inet_down_cost"`
	InetUp             float64 `json:"inet_up"`
	InetUpCost         float64 `json:"inet_up_cost"`
	IsBid              bool    `json:"is_bid"`
	Logo               string  `json:"logo"`
	MachineId          int     `json:"machine_id"`
	MinBid             float64 `json:"min_bid"`
	MoboName           string  `json:"mobo_name"`
	NumGpus            int     `json:"num_gpus"`
	OsVersion          string  `json:"os_version"`
	PciGen             float64 `json:"pci_gen"`
	PcieBw             float64 `json:"pcie_bw"`
	PublicIpaddr       string  `json:"public_ipaddr"`
	Reliability        float64 `json:"reliability"`
	Reliability2       float64 `json:"reliability2"`
	ReliabilityMult    float64 `json:"reliability_mult"`
	Rentable           bool    `json:"rentable"`
	Rented             bool    `json:"rented"`
	Rn                 int     `json:"rn"`
	Score              float64 `json:"score"`
	StartDate          float64 `json:"start_date"`
	StaticIp           bool    `json:"static_ip"`
	StorageCost        float64 `json:"storage_cost"`
	StorageTotalCost   float64 `json:"storage_total_cost"`
	TotalFlops         float64 `json:"total_flops"`
	Vericode           int     `json:"vericode"`
	Verification       string  `json:"verification"`
	VramCostperhour    float64 `json:"vram_costperhour"`
	Webpage            string  `json:"webpage,omitempty"`
}
