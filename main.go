package main

import (
	"fmt"
	"log"
	"net/http"
	// TODO: Import package yang diperlukan (uncomment dan modifikasi line 8 - 11)
	// "namafolder/configs"
	// "namafolder/handlers"
	// "namafolder/middlewares"
	// _ "github.com/go-sql-driver/mysql"
)

func main() {
    // TODO: Ganti dengan 5 digit terakhir NPM masing - masing.
    // Contoh: npm 56419080, maka ubah angka port: 19080
    PORT := 28081 

    // TODO: Inisialisasi koneksi database


    // Inisialisasi server dan route
    // Buat mux untuk routing
    mux := http.NewServeMux()
    
    // Buat file server untuk konten statis
    fileServer := http.FileServer(http.Dir("catalog"))
    
    // Tambahkan route untuk file statis
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        handlers.ServeStaticFile(w, r, "catalog", fileServer)
    })
    
    // TODO: Tambahkan route untuk API events, lihat contoh di atas untuk menambahkan route. contoh lain: mux.HandleFunc("/api/endpoint1/", handlers.HandleEvents)
    

    // Terapkan middleware logging
    loggedMux := middlewares.LogRequestHandler(mux)

    // Jalankan server dengan port yang sudah ditentukan
    fmt.Printf("Server berjalan di http://localhost:%d\n", PORT)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), loggedMux))
}