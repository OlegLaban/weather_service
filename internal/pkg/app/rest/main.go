package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/OlegLaban/weather_service/internal/contracts"
	"github.com/go-chi/chi/v5"
)

type app struct {
	h contracts.Handlers
	l contracts.Logger
	c contracts.RestConfig
}

func New(h contracts.Handlers, l contracts.Logger, c contracts.RestConfig) *app {
	return &app{
		h: h,
		l: l,
		c: c,
	}
}

func (a *app) Run() error {
	r := chi.NewRouter()
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		answer, err := a.h.Ping()
		if err != nil {
			msg := "can`t get answer from Ping method of handlers"
			a.l.Error(msg, err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{"error": %s}`, msg)
		}
		jsonBody, err := json.Marshal(answer)
		if err != nil {
			msg := "can`t marshal response for Ping handler"
			a.l.Error(msg, err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{"error": %s}`, msg)
		}
		w.Write(jsonBody)
	})

	return http.ListenAndServe(a.c.Port(), r)
}
