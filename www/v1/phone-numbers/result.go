package verboten

type Result struct {
	PhoneNumber      string `json:"phoneNumber"`
	CountryCode      string `json:"countryCode"`
	AreaCode         string `json:"areaCode"`
	LocalPhoneNumber string `json:"localPhoneNumber"`
}
