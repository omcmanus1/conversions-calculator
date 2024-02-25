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
}

func GetRequest(w http.ResponseWriter, r *http.Request) {
	GetHandlerEncode(w, r)
}

func PostListHTTP(w http.ResponseWriter, r *http.Request) {
	setHeaders(w, r)
	PostConversions(w, r)
}

func PostWeightUsHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")

	PostWeightUS(w, r)
}

func PostVolumeUsHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")

	PostVolumeUS(w, r)
}

func PostWeightMetricHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")

	PostWeightMetric(w, r)
}

func PostVolumeMetricHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")

	PostVolumeMetric(w, r)
}

func setHeaders(w http.ResponseWriter, r *http.Request) http.ResponseWriter {
	corsOrigin := os.Getenv("CORS_ORIGIN")
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", corsOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
	}
	return w
}
