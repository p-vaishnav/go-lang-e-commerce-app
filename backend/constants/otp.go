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
