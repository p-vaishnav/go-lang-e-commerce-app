package backenderrors

var errorType = struct {
	validation         string
	server             string
	Unauthorized       string
	conflict           string
	ServiceUnavailable string
	NotFound           string
	Downstream         string
	Forbidden          string
}{
	validation:         "validation",
	server:             "server",
	Unauthorized:       "unauthorized",
	conflict:           "conflict",
	ServiceUnavailable: "service unavailable",
	NotFound:           "not found",
	Downstream:         "downstream",
	Forbidden:          "forbidden",
}
