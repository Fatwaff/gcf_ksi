package gcf

import (
	"fmt"
	"net/http"

	"github.com/Fatwaff/be_ksi"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("KSI", KSI_Sewa)
}

func KSI_Sewa(w http.ResponseWriter, r *http.Request) {
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
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if r.Method == http.MethodPost {
		fmt.Fprintf(w, be_ksi.SewaHandler("PASETOPUBLICKEY", "MONGOSTRING", "db_ksi", r))
		return
	}
	if r.Method == http.MethodPut {
		fmt.Fprintf(w, be_ksi.EditSewaHandler("PASETOPUBLICKEY", "MONGOSTRING", "db_ksi", r))
		return
	}
	if r.Method == http.MethodDelete {
		fmt.Fprintf(w, be_ksi.HapusSewaHandler("PASETOPUBLICKEY", "MONGOSTRING", "db_ksi", r))
		return
	}
	// Set CORS headers for the main request.
	fmt.Fprintf(w, be_ksi.GetSewaHandler("PASETOPUBLICKEY", "MONGOSTRING", "db_ksi", r))

}
