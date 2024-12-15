package router

import "net/http"

// returns HTTP request multiplexer
func GetHTTPMultiplexer() *http.ServeMux {
	// Setup router (using standard net/http for example)
	mux := http.NewServeMux()

	// ================================================
	// Order routes
	// ================================================
	mux.HandleFunc("/orders")
	mux.HandleFunc("/orders/{id}")
	mux.HandleFunc("/orders/{id}/close")

	// ================================================
	// Menu routes
	// ================================================
	mux.HandleFunc("/menu")
	mux.HandleFunc("/menu/{id}")

	// ================================================
	// Inventory routes
	// ================================================
	mux.HandleFunc("/inventory")
	mux.HandleFunc("/inventory/{id}")
	mux.HandleFunc("/reports/total-sales")
	mux.HandleFunc("/reports/popular-items")

	return mux
}
