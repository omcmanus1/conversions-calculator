package converter

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("GetRequest", GetRequest)
}

func GetRequest(w http.ResponseWriter, r *http.Request) {
	GetHandlerEncode(w, r)
	fmt.Fprintf(w, "Function output: %v!", r)
}

func Routes(w http.ResponseWriter, r *http.Request)