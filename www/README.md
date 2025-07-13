# `www/`

This contains the `http.Handler`s.

The directory structure roughly matches the HTTP request-paths.

So, for example, the HTTP request-path:

`/v1/phone-numbers`

... is handled by the Go code:

`www/v1/phone-numbers/servehttp.go`
