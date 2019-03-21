package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"github.com/go-chi/render"
	"log"
	"github.com/rodion-goritskov/supuesto/backend/apiv1"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.RedirectSlashes,
		)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/", apiv1.Routes())
	})

	return router
}

func main() {
	router := Routes()

	log.Fatal(http.ListenAndServe(":4444", router))
}
