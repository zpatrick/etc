package main

import "net/http"

const (
	ShowMovie = "..."
)

type MyController struct{}

func (m *MyController) Routes() pkg.Routes {
	return pkg.Routes{
		AroundEach: RequireLogin,
		Routes: []pkg.Route{
			pkg.FooRoute{
				Around:  RequireScope("read", "article"),
				Methods: []string{http.MethodGet},
			},
		},
	}
}
