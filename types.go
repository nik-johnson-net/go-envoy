package envoy

// DevControl is a list of controls for devices
type DevControl struct {
	Gficlearset bool `json:"gficlearset,omitempty"`
}

// Device describes a device attached to the Envoy system
type Device struct {
	PartNum        string       `json:"part_num,omitempty"`
	Installed      int          `json:"installed,omitempty"`
	SerialNum      int          `json:"serial_num,omitempty"`
	DeviceStatus   []string     `json:"device_status,omitempty"`
	LastReportDate int          `json:"last_report_date,omitempty"`
	AdminState     int          `json:"admin_state,omitempty"`
	DevType        int          `json:"dev_type,omitempty"`
	CreatedDate    int          `json:"created_date,omitempty"`
	ImgLoadDate    int          `json:"img_load_date,omitempty"`
	ImgPnumRunning string       `json:"img_pnum_running,omitempty"`
	Ptpn           string       `json:"ptpn,omitempty"`
	Chaneid        int          `json:"chaneid,omitempty"`
	DeviceControl  []DevControl `json:"device_control,omitempty"`
	Producing      bool         `json:"producing,omitempty"`
	Communicating  bool         `json:"communicating,omitempty"`
	Provisioned    bool         `json:"provisioned,omitempty"`
	Operating      bool         `json:"operating,omitempty"`
}

// Inventory describes a list of Devices of a certain Type
type Inventory struct {
	Type    string   `json:"type,omitempty"`
	Devices []Device `json:"devices,omitempty"`
}

// ProductionData describes a power reading from a particular sensor. May be a consumption meter.
type ProductionData struct {
	Type             string  `json:"type,omitempty"`
	ActiveCount      int     `json:"activeCount,omitempty"`
	MeasurementType  string  `json:"measurementType,omitempty"`
	ReadingTime      int     `json:"readingTime,omitempty"`
	WNow             float64 `json:"wNow,omitempty"`
	WhLifetime       float64 `json:"whLifetime,omitempty"`
	VarhLeadLifetime float64 `json:"varhLeadLifetime,omitempty"`
	VarhLagLifetime  float64 `json:"varhLagLifetime,omitempty"`
	VahLifetime      float64 `json:"vahLifetime,omitempty"`
	RmsCurrent       float64 `json:"rmsCurrent,omitempty"`
	RmsVoltage       float64 `json:"rmsVoltage,omitempty"`
	ReactPwr         float64 `json:"reactPwr,omitempty"`
	ApprntPwr        float64 `json:"apprntPwr,omitempty"`
	PwrFactor        float64 `json:"pwrFactor,omitempty"`
	WhToday          float64 `json:"whToday,omitempty"`
	WhLastSevenDays  float64 `json:"whLastSevenDays,omitempty"`
	VahToday         float64 `json:"vahToday,omitempty"`
	VarhLeadToday    float64 `json:"varhLeadToday,omitempty"`
	VarhLagToday     float64 `json:"varhLagToday,omitempty"`
	State            string  `json:"state,omitempty"`
}

// Production is the collection of all power sensors in the system.
type Production struct {
	Production  []ProductionData `json:"production,omitempty"`
	Consumption []ProductionData `json:"consumption,omitempty"`
	Storage     []ProductionData `json:"storage,omitempty"`
}
