package gcf

import (
	"fmt"
	"net/http"

	"github.com/Fatwaff/be_ksi"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("KSI", KSI_Register)
}

func KSI_Register(w http.ResponseWriter, r *http.Request) {
	allowedOrigins := []string{"https://ksi-billboard.github.io", "http://127.0.0.1:5500", "http://127.0.0.1:5501"}
	origin := r.Header.Get("Origin")

	for _, allowedOrigin := range allowedOrigins {
		if allowedOrigin == origin {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			break
		}
	}
	// w.Header().Set("Access-Control-Allow-Origin", "https://intern-monitoring.github.io")
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	fmt.Fprintf(w, be_ksi.SignUpHandler("MONGOSTRING", "db_ksi", r))

}