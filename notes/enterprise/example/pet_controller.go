package controllers

import (
	"net/http"

	p "google.golang.org/genproto"
)

func (m *MovieController) Routes() []enterprise.Route {
	return []*enterprise.Route{
		{
			Name: "list_movies",
			Route: func(r *http.Request) bool {
				r.Method == http.MethodGet && r.URL.Path.String() == "/movies"
			},
			Middleware: []enterprise.Middleware{
				p.auth.EnsureLogin,
				p.auth.CanListMovies,
				p.metrics.Duration("list_movies"),
				p.logger.Log("list_movies"),
			},
			Handler: http.HandlerFunc(m.listMovies),
		},
		{
                        Name: "get_movie",
                        Route: func(r *http.Request) bool {
                                r.Method == http.MethodGet && glob.Match(r.URL.Path.String(), "/movies/*/")
                        },
                        Middleware: []enterprise.Middleware{
                                p.auth.EnsureLogin,
                                p.auth.CanListMovies,
                                p.metrics.Duration("list_movies"),
                                p.logger.Log("list_movies"),
                        },
                        Handler: http.HandlerFunc(m.listMovies),
                },
	}
}
