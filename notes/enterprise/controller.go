package enterprise

import "net/http"

type Controller struct {
	Name     string
	Before   http.Handler
	After    http.Handler
	Handlers []Handler
}
