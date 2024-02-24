package constants

var OTP_STATUS = struct {
	SUCCESS string
	FAILURE string
	PENDING string
}{
	SUCCESS: "SUCCESS",
	FAILURE: "FAILURE",
	PENDING: "PENDING",
}

var OTP_MEDIUM = map[string]bool{
	"email":    true,
	"sms":      true,
	"whatsapp": true,
}

var OTP_PROVIDER = struct {
	MSG_91    string
	SEND_GRID string
	WHATS_APP string
}{
	MSG_91:    "msg91",
	SEND_GRID: "send_grid",
	WHATS_APP: "whats_app",
}
