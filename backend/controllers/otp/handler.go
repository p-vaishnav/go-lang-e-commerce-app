package otp

type otpHandler struct{}

func Handler() *otpHandler {
	return &otpHandler{}
}
