package verboten

type ErrResult struct {
	PhoneNumber string `json:"phoneNumber"`
	Error       string `json:"error"`
}
