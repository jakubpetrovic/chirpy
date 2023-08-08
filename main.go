package main

import (
	"log"
	"net/http"
)

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {

	const rootFilepath = "."
	const port = "8080"
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(rootFilepath)))

	corsMux := middlewareCors(mux)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}
	srv.Addr = ":8080"
	srv.Handler = corsMux

	log.Printf("Serving files from %s on port: %s\n", rootFilepath, port)
	log.Fatal(srv.ListenAndServe())

}
