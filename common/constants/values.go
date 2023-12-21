package constants

import (
	"time"
)

var HttpMethodsE = struct {
	GET    string
	POST   string
	DELETE string
}{
	GET:    "GET",
	POST:   "POST",
	DELETE: "DELETE",
}

var AUTH_TTL = 10 * time.Minute
