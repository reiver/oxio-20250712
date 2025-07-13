package verboten

import (
	"encoding/json"
	"net/http"

	"github.com/reiver/go-e164"
	"github.com/reiver/go-erorr"
	"github.com/reiver/go-http400"
	"github.com/reiver/go-http500"

	"github.com/reiver/oxio-20250712/srv/http"
)

const path string = "/v1/phone-numbers"

const queryPhoneNumber string = "phoneNumber"
const queryCountryCode string = "countryCode"

func init() {
	var handler http.Handler = http.HandlerFunc(serveHTTP)

	err := httpsrv.Mux.HandlePath(handler, path)
	if nil != err {
		e := erorr.Errorf("problem registering http-handler with path-mux for path %q: %w", path, err)
		panic(e)
	}
}

func serveHTTP(responsewriter http.ResponseWriter, request *http.Request) {

	if nil == responsewriter {
		return
	}
	if nil == request {
		http500.InternalServerError(responsewriter, request)
		return
	}
	if nil == request.URL {
		http500.InternalServerError(responsewriter, request)
		return
	}

	query := request.URL.Query()
	if len(query) < 1 {
		http400.BadRequest(responsewriter, request)
		return
	}

	queryPhoneNumberValue := query.Get(queryPhoneNumber)
	queryCountryCodeValue := query.Get(queryCountryCode)

	var countryCode string
	var nationalDestinationCode string
	var subscriberNumber string
	switch {
	case "" == queryCountryCodeValue:
		var err error

		countryCode, nationalDestinationCode, subscriberNumber, err = e164.ParseTolerantly(queryPhoneNumberValue)
		if nil != err {
			var errResult = ErrResult{
				PhoneNumber: queryPhoneNumberValue,
			}
			errResult.SetError(err)

			responsewriter.Header().Add("Content-Type", "application/json")
			json.NewEncoder(responsewriter).Encode(errResult)
			return
		}
	default:
		var err error

		countryCode, nationalDestinationCode, subscriberNumber, err = e164.ParseWithIsoCountryCodeTolerantly(queryCountryCodeValue, queryPhoneNumberValue)
		if nil != err {
			var errResult = ErrResult{
				PhoneNumber: queryPhoneNumberValue,
			}
			errResult.SetError(err)

			responsewriter.Header().Add("Content-Type", "application/json")
			json.NewEncoder(responsewriter).Encode(errResult)
			return
		}
	}

	var result = Result{
		PhoneNumber:      queryPhoneNumberValue,
		CountryCode:      e164.ISOCountryCode(countryCode, nationalDestinationCode),
		AreaCode:         nationalDestinationCode,
		LocalPhoneNumber: subscriberNumber,
	}

	responsewriter.Header().Add("Content-Type", "application/json")
	json.NewEncoder(responsewriter).Encode(result)
	return
}
