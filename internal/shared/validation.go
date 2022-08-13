package shared

var (
	// Account Role
	IsAllowedAccountRole = map[string]bool{
		SUPER_USER: true,
		CUSTOMER:   true,
		SELLER:     true,
	}

	// Store permission
	IsAllowedStoreAccountRole = map[string]bool{
		SUPER_USER: true,
		SELLER:     true,
	}
)
