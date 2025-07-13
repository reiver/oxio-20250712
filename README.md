# oxio-20250712

This has a single HTTP API end-point:

`/v1/phone-numbers`

It takes request parameters and returns information about a phone number.

## Request parameter

### `phoneNumber`

* in E.164 format: [+][country code][area code][local phone number]
* `+` is optional
* phone number is sequence of digits
* phones can spaces between country, area code and local phone number
  * any other space is invalid
  * any other character is invalid
* For example
  * +12125690123 (valid)
  * +52 631 3118150 (valid)
  * 34 915 872200 (valid)
  * 351 21 094 2000 (invalid)
* if phoneNumber is missing country code then user must provide countryCode parameter in ISO 3166-1 alpha-2 format.

### `countryCode`

* in ISO 3166-1 alpha-2 format
* For example:
  * US (valid)
  * MX (valid)
  * ESP (invalid)

## Response

### success

`/v1/phone-numbers?phoneNumber=%2B12125690123`

Note: `+` is URL encodes as `%2B`

```json
{
	"phoneNumber": "+12125690123",
	"countryCode": "US",
	"areaCode": "212",
	"localPhoneNumber": "5690123"
}
```

### error

`/v1/phone-numbers?phoneNumber=631%20311%208150`

Note: `[space]` is URL-encoded as `%20`

```json
{
	phoneNumber": "25690123",
	"error": {
		"countryCode": "required value is missing"
	}
}
```

