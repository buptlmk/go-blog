package utils

const (
	StatusOK       = 200 // RFC 7231, 6.3.1
	StatusCreated  = 201 // RFC 7231, 6.3.2
	StatusAccepted = 202 // RFC 7231, 6.3.3

	StatusMultipleChoices  = 300 // RFC 7231, 6.4.1
	StatusMovedPermanently = 301 // RFC 7231, 6.4.2
	StatusFound            = 302 // RFC 7231, 6.4.3

	StatusBadRequest   = 400 // RFC 7231, 6.5.1
	StatusUnauthorized = 401 // RFC 7235, 3.1
	StatusForbidden    = 403 // RFC 7231, 6.5.3

	StatusInternalServerError = 500 // RFC 7231, 6.6.1
	StatusNotImplemented      = 501 // RFC 7231, 6.6.2
	StatusBadGateway          = 502 // RFC 7231, 6.6.3
	StatusServiceUnavailable  = 503 // RFC 7231, 6.6.4
	StatusGatewayTimeout      = 504 // RFC 7231, 6.6.5

)
