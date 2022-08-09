package core

const (
	httpFailed  = -1
	httpSuccess = 0
)

const (
	// ProviderTypeCPL https://github.com/clarketm/proxy-list
	ProviderTypeCPL = "CPL"
	// ProviderTypeTSXPL https://github.com/TheSpeedX/PROXY-List
	ProviderTypeTSXPL = "TSX"
	// ProviderTypeSTRPL https://github.com/ShiftyTR/Proxy-List
	ProviderTypeSTRPL = "STR"
	// ProviderTypeIHuan https://ip.ihuan.me/ti.html
	ProviderTypeIHuan = "IHUAN"
	// ProviderTypeMix all type
	ProviderTypeMix = "MIX"
)

var (
	AllProviderType = []string{
		ProviderTypeCPL,
		ProviderTypeTSXPL,
		ProviderTypeSTRPL,
		ProviderTypeIHuan,
	}
	UserAvailableProviderType = []string{
		ProviderTypeCPL,
		ProviderTypeTSXPL,
		ProviderTypeSTRPL,
		ProviderTypeIHuan,
		ProviderTypeMix,
	}
)
