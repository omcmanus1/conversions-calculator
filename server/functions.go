package converter

import (
	"net/http"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

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

func requestHandler(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		setHeaders(w, r)
		handlerFunc(w, r)
	}
}

func init() {
	functions.HTTP("PostConversions", requestHandler(PostConversions))
	functions.HTTP("PostWeightUS", requestHandler(PostWeightUS))
	functions.HTTP("PostVolumeUS", requestHandler(PostVolumeUS))
	functions.HTTP("PostWeightMetric", requestHandler(PostWeightMetric))
	functions.HTTP("PostVolumeMetric", requestHandler(PostVolumeMetric))
	functions.HTTP("PostHeightMetric", requestHandler(PostHeightMetric))
	functions.HTTP("PostHeightFeet", requestHandler(PostHeightFeet))
}
