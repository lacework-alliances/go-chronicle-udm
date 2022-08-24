package udm

var AuthTypes = []string{SSO, MACHINE, VPN, TACACS, PHYSICAL}
var MechanismTypes = []string{BADGE_READER, BATCH, CACHED_INTERACTIVE, HARDWARE_KEY, LOCAL, MECHANISM_OTHER, NETWORK,
	NETWORK_CLEAR_TEXT, NEW_CREDENTIALS, OTP, REMOTE, REMOTE_INTERACTIVE, SERVICE, UNLOCK, USERNAME_PASSWORD}

// UserLoginLogout validates the USER_LOGIN and USER_LOGOUT event types
// Required Fields Link: https://cloud.google.com/chronicle/docs/unified-data-model/udm-usage#user_login
// Must have principal.hostname and target.hostname
// Must have extension.auth.type and extension.auth.mechanism
func UserLoginLogout(udm *UDM) (bool, []string) {
	var message []string
	var failed bool
	// validate principal.hostname and target.hostname exist
	if udm.Principal == nil || udm.Target == nil {
		message = append(message, "Principal and/or Target are missing")
		failed = true
	} else if udm.Principal.Hostname == "" || udm.Target.Hostname == "" {
		message = append(message, "Principal and/or Target are missing the required hostname field")
		failed = true
	}
	// validate extension.auth.type and extension.auth.mechanism
	if udm.Extensions == nil {
		message = append(message, "Authentication Extension is missing")
		failed = true
	} else if !stringInSlice(udm.Extensions.Auth.AuthType, AuthTypes) {
		udm.Extensions.Auth.AuthType = AUTHTYPE_UNSPECIFIED
	} else if len(udm.Extensions.Auth.Mechanism) == 0 {
		udm.Extensions.Auth.Mechanism = []string{MECHANISM_UNSPECIFIED}
	}
	return failed, message
}

// UserResource returns the minimum required fields for the following event types
// USER_RESOURCE_ACCESS, USER_RESOURCE_CREATION, USER_RESOURCE_DELETION, USER_RESOURCE_UPDATE_CONTENT
// USER_RESOURCE_UPDATE_PERMISSIONS, USER_UNCATEGORIZED
// returns two Noun structs, principal and target
func UserResource(principalUser string, targetResource string) (Noun, Noun) {
	var principal, target Noun
	principal.User = &User{UserID: principalUser}
	target.Resource = &Resource{Name: targetResource}
	return principal, target
}
