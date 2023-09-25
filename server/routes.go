package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/omcmanus1/converter/data"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Get("/api", Home)
	r.Get("/api/convert/get-encode", GetHandlerEncode)
	r.Get("/api/convert/get-marshal", GetHandlerMarshal)
	r.Post("/api/convert/list", PostConversions)
	r.Post("/api/convert/weight-us", PostWeightUS)
	r.Post("/api/convert/volume-us", PostVolumeUS)
	r.Post("/api/convert/weight-metric", PostWeightMetric)
	r.Post("/api/convert/volume-metric", PostVolumeMetric)

	return r
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello, welcome to the recipe converter..."))
}

func GetHandlerMarshal(w http.ResponseWriter, r *http.Request) {
	HandleGetRequestMarshal(w, r, data.Input, Flow)
}

func GetHandlerEncode(w http.ResponseWriter, r *http.Request) {
	HandleGetRequestEncode(w, r, data.SingleInput, Flow)
}

func PostConversions(w http.ResponseWriter, r *http.Request) {
	HandlePostRequest(w, r, Flow)
}

func PostWeightUS(w http.ResponseWriter, r *http.Request) {
	HandlePostRequest(w, r, WeightUS)
}

func PostVolumeUS(w http.ResponseWriter, r *http.Request) {
	HandlePostRequest(w, r, VolumeUS)
}

func PostWeightMetric(w http.ResponseWriter, r *http.Request) {
	HandlePostRequest(w, r, WeightMetric)
}

func PostVolumeMetric(w http.ResponseWriter, r *http.Request) {
	HandlePostRequest(w, r, VolumeMetric)
}
