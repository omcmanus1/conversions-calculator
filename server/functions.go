package converter

import (
	"net/http"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("GetRequest", GetRequest)
	functions.HTTP("PostListHTTP", PostListHTTP)
	functions.HTTP("PostWeightUsHTTP", PostWeightUsHTTP)
	functions.HTTP("PostVolumeUsHTTP", PostVolumeUsHTTP)
	functions.HTTP("PostWeightMetricHTTP", PostWeightMetricHTTP)
	functions.HTTP("PostVolumeMetricHTTP", PostVolumeMetricHTTP)
	functions.HTTP("PostHeightMetricHTTP", PostHeightMetricHTTP)
	functions.HTTP("PostHeightFeetHTTP", PostHeightFeetHTTP)
}

func GetRequest(w http.ResponseWriter, r *http.Request) {
	GetHandlerEncode(w, r)
}

func PostListHTTP(w http.ResponseWriter, r *http.Request) {
	setHeaders(w, r)
	PostConversions(w, r)
}

func PostWeightUsHTTP(w http.ResponseWriter, r *http.Request) {
	setHeaders(w, r)
	PostWeightUS(w, r)
}

func PostVolumeUsHTTP(w http.ResponseWriter, r *http.Request) {
	setHeaders(w, r)
	PostVolumeUS(w, r)
}

func PostWeightMetricHTTP(w http.ResponseWriter, r *http.Request) {
	setHeaders(w, r)
	PostWeightMetric(w, r)
}

func PostVolumeMetricHTTP(w http.ResponseWriter, r *http.Request) {
	setHeaders(w, r)
	PostVolumeMetric(w, r)
}

func PostHeightMetricHTTP(w http.ResponseWriter, r *http.Request) {
	setHeaders(w, r)
	PostHeightMetric(w, r)
}

func PostHeightFeetHTTP(w http.ResponseWriter, r *http.Request) {
	setHeaders(w, r)
	PostHeightFeet(w, r)
}

func setHeaders(w http.ResponseWriter, r *http.Request) {
	corsOrigin := os.Getenv("CORS_ORIGIN")
	w.Header().Set("Access-Control-Allow-Origin", corsOrigin)
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Max-Age", "3600")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}
}
