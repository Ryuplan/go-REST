package internal

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RunRouter() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	fmt.Println("Inside from other package")
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello World!"))
		if err != nil {
			return
		}
	})
	http.ListenAndServe(":3000", r)
}
