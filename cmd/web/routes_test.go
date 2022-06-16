package main

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi"
	"github.com/martino/bookings/internal/config"
)

func TestRoutes(t *testing.T) {

	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing; test pass
	default:
		t.Error(fmt.Sprintf("type is not *chi.Mux, but type is %T", v))
	}
}
