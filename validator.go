package udm

var AuthTypes = []string{SSO, MACHINE, VPN, TACACS, PHYSICAL}
var MechanismTypes = []string{BADGE_READER, BATCH, CACHED_INTERACTIVE, HARDWARE_KEY, LOCAL, MECHANISM_OTHER, NETWORK,
	NETWORK_CLEAR_TEXT, NEW_CREDENTIALS, OTP, REMOTE, REMOTE_INTERACTIVE, SERVICE, UNLOCK, USERNAME_PASSWORD}

// UserLoginLogout validates the USER_LOGIN and USER_LOGOUT event types
// Required Fields Link: https://cloud.google.com/chronicle/docs/unified-data-model/udm-usage#user_login
// Must have principal.hostname and target.hostname
// Must have extension.auth.type and extension.auth.mechanism
func UserLoginLogout(udm *UDM) bool {
	var message []string
	var failed bool
	// validate principal.hostname and target.hostname exist
	if udm.Principal.Hostname == "" || udm.Target.Hostname == "" {
		message = append(message, "Principal and/or Target are missing the required hostname field")
		failed = true
	}
	// validate extension.auth.type and extension.auth.mechanism
	if !stringInSlice(udm.Extensions.Auth.AuthType, AuthTypes) {
		udm.Extensions.Auth.AuthType = AUTHTYPE_UNSPECIFIED
	}
	if len(udm.Extensions.Auth.Mechanism) == 0 {
		udm.Extensions.Auth.Mechanism = []string{MECHANISM_UNSPECIFIED}
	}
	return failed
}
