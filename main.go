package main

import (
	"fmt"
	"log"
	"net/http"

	"eventrealm_release/configs"     // UBAH 12345 DENGAN NPM KALIAN
	"eventrealm_release/handlers"    // UBAH 12345 DENGAN NPM KALIAN
	"eventrealm_release/middlewares" // UBAH 12345 DENGAN NPM KALIAN

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// UBAH DENGAN 5 DIGIT TERAKHIR NPM KALIAN
	PORT := 12345

	configs.ConnectDB()

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("catalog"))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeStaticFile(w, r, "catalog", fileServer)
	})

	// Route API
	mux.HandleFunc("/api/events/", handlers.HandleEvents)
	mux.HandleFunc("/api/events", handlers.HandleEvents)

	loggedMux := middlewares.LogRequestHandler(mux)

	fmt.Printf("Server berjalan di http://localhost:%d\n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), loggedMux))
}
