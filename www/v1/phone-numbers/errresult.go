package verboten

import (
	"errors"

	"github.com/reiver/go-e164"
)

type ErrResult struct {
	PhoneNumber string `json:"phoneNumber"`
	Error struct {
		CountryCode      string `json:"countryCode,omitempty"`
		AreaCode         string `json:"areaCode,omitempty"`
		LocalPhoneNumber string `json:"localPhoneNumber,omitempty"`
		Message          string `json:"message,omitempty"`
	} `json:"error"`
}

func (receiver *ErrResult) SetError(err error) {
	if nil == receiver {
		return
	}
	if nil == err {
		return
	}

	switch {
	case errors.Is(err, e164.ErrBadCountryCode):
		receiver.Error.CountryCode = err.Error()
	case errors.Is(err, e164.ErrBadNationalDestinationCode):
		receiver.Error.AreaCode = err.Error()
	case errors.Is(err, e164.ErrBadSubscriberNumber):
		receiver.Error.LocalPhoneNumber = err.Error()
	default:
		receiver.Error.Message = err.Error()
	}
}
