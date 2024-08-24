package octopus

type StatusCode int
type codeMessage string

const (
	StatusContinue                      StatusCode = 100
	StatusSwitchingProtocols            StatusCode = 101
	StatusProcessing                    StatusCode = 102
	StatusEarlyHints                    StatusCode = 103
	StatusOK                            StatusCode = 200
	StatusCreated                       StatusCode = 201
	StatusAccepted                      StatusCode = 202
	StatusNonAuthoritative              StatusCode = 203
	StatusNoContent                     StatusCode = 204
	StatusResetContent                  StatusCode = 205
	StatusPartialContent                StatusCode = 206
	StatusMultiStatus                   StatusCode = 207
	StatusAlreadyReported               StatusCode = 208
	StatusIMUsed                        StatusCode = 226
	StatusMultipleChoices               StatusCode = 300
	StatusMovedPermanently              StatusCode = 301
	StatusFound                         StatusCode = 302
	StatusSeeOther                      StatusCode = 303
	StatusNotModified                   StatusCode = 304
	StatusUseProxy                      StatusCode = 305
	StatusTemporaryRedirect             StatusCode = 307
	StatusPermanentRedirect             StatusCode = 308
	StatusBadRequest                    StatusCode = 400
	StatusUnauthorized                  StatusCode = 401
	StatusPaymentRequired               StatusCode = 402
	StatusForbidden                     StatusCode = 403
	StatusNotFound                      StatusCode = 404
	StatusMethodNotAllowed              StatusCode = 405
	StatusNotAcceptable                 StatusCode = 406
	StatusProxyAuthRequired             StatusCode = 407
	StatusRequestTimeout                StatusCode = 408
	StatusConflict                      StatusCode = 409
	StatusGone                          StatusCode = 410
	StatusLengthRequired                StatusCode = 411
	StatusPreconditionFailed            StatusCode = 412
	StatusPayloadTooLarge               StatusCode = 413
	StatusURITooLong                    StatusCode = 414
	StatusUnsupportedMedia              StatusCode = 415
	StatusRangeNotSatisfiable           StatusCode = 416
	StatusExpectationFailed             StatusCode = 417
	StatusTeapot                        StatusCode = 418
	StatusMisdirectedRequest            StatusCode = 421
	StatusUnprocessableEntity           StatusCode = 422
	StatusLocked                        StatusCode = 423
	StatusFailedDependency              StatusCode = 424
	StatusTooEarly                      StatusCode = 425
	StatusUpgradeRequired               StatusCode = 426
	StatusPreconditionRequired          StatusCode = 428
	StatusTooManyRequests               StatusCode = 429
	StatusRequestHeaderFieldsTooLarge   StatusCode = 431
	StatusUnavailableForLegalReasons    StatusCode = 451
	StatusInternalServerError           StatusCode = 500
	StatusNotImplemented                StatusCode = 501
	StatusBadGateway                    StatusCode = 502
	StatusServiceUnavailable            StatusCode = 503
	StatusGatewayTimeout                StatusCode = 504
	StatusHTTPVersionNotSupported       StatusCode = 505
	StatusVariantAlsoNegotiates         StatusCode = 506
	StatusInsufficientStorage           StatusCode = 507
	StatusLoopDetected                  StatusCode = 508
	StatusNotExtended                   StatusCode = 510
	StatusNetworkAuthenticationRequired StatusCode = 511
)

var statusMessages = map[StatusCode]codeMessage{
	StatusContinue:                      "Continue",
	StatusSwitchingProtocols:            "Switching Protocols",
	StatusProcessing:                    "Processing",
	StatusEarlyHints:                    "Early Hints",
	StatusOK:                            "OK",
	StatusCreated:                       "Created",
	StatusAccepted:                      "Accepted",
	StatusNonAuthoritative:              "Non-Authoritative Information",
	StatusNoContent:                     "No Content",
	StatusResetContent:                  "Reset Content",
	StatusPartialContent:                "Partial Content",
	StatusMultiStatus:                   "Multi-Status",
	StatusAlreadyReported:               "Already Reported",
	StatusIMUsed:                        "IM Used",
	StatusMultipleChoices:               "Multiple Choices",
	StatusMovedPermanently:              "Moved Permanently",
	StatusFound:                         "Found",
	StatusSeeOther:                      "See Other",
	StatusNotModified:                   "Not Modified",
	StatusUseProxy:                      "Use Proxy",
	StatusTemporaryRedirect:             "Temporary Redirect",
	StatusPermanentRedirect:             "Permanent Redirect",
	StatusBadRequest:                    "Bad Request",
	StatusUnauthorized:                  "Unauthorized",
	StatusPaymentRequired:               "Payment Required",
	StatusForbidden:                     "Forbidden",
	StatusNotFound:                      "Not Found",
	StatusMethodNotAllowed:              "Method Not Allowed",
	StatusNotAcceptable:                 "Not Acceptable",
	StatusProxyAuthRequired:             "Proxy Authentication Required",
	StatusRequestTimeout:                "Request Timeout",
	StatusConflict:                      "Conflict",
	StatusGone:                          "Gone",
	StatusLengthRequired:                "Length Required",
	StatusPreconditionFailed:            "Precondition Failed",
	StatusPayloadTooLarge:               "Payload Too Large",
	StatusURITooLong:                    "URI Too Long",
	StatusUnsupportedMedia:              "Unsupported Media Type",
	StatusRangeNotSatisfiable:           "Range Not Satisfiable",
	StatusExpectationFailed:             "Expectation Failed",
	StatusTeapot:                        "I'm a teapot",
	StatusMisdirectedRequest:            "Misdirected Request",
	StatusUnprocessableEntity:           "Unprocessable Entity",
	StatusLocked:                        "Locked",
	StatusFailedDependency:              "Failed Dependency",
	StatusTooEarly:                      "Too Early",
	StatusUpgradeRequired:               "Upgrade Required",
	StatusPreconditionRequired:          "Precondition Required",
	StatusTooManyRequests:               "Too Many Requests",
	StatusRequestHeaderFieldsTooLarge:   "Request Header Fields Too Large",
	StatusUnavailableForLegalReasons:    "Unavailable For Legal Reasons",
	StatusInternalServerError:           "Internal Server Error",
	StatusNotImplemented:                "Not Implemented",
	StatusBadGateway:                    "Bad Gateway",
	StatusServiceUnavailable:            "Service Unavailable",
	StatusGatewayTimeout:                "Gateway Timeout",
	StatusHTTPVersionNotSupported:       "HTTP Version Not Supported",
	StatusVariantAlsoNegotiates:         "Variant Also Negotiates",
	StatusInsufficientStorage:           "Insufficient Storage",
	StatusLoopDetected:                  "Loop Detected",
	StatusNotExtended:                   "Not Extended",
	StatusNetworkAuthenticationRequired: "Network Authentication Required",
}
