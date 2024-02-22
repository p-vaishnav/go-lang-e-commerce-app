package otpsvc

type SendOTPReq struct {
	Mobile string `json:"mobile"`
	Email  string `json:"email"`
	Medium string `json:"medium"`
}
