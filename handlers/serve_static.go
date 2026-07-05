package handlers

import (
	"net/http"
	// "os"
	// "path/filepath"
	// "strings"
)

// TODO: Implementasikan fungsi ServeStaticFile untuk melayani file statis dari folder catalog
// Fungsi ini harus:
// 1. Menerima parameter http.ResponseWriter, *http.Request, baseDir string, dan fileServer http.Handler
// 2. Tidak melayani permintaan ke path /api/
// 3. Mengecek keberadaan file
// 4. Mengarahkan ke halaman utama jika file tidak ditemukan
// 5. Melayani file statis menggunakan fileServer

// ServeStaticFile melayani file statis dari folder catalog
func ServeStaticFile(w http.ResponseWriter, r *http.Request, baseDir string, fileServer http.Handler) {
	// TODO: Implementasi melayani file statis
	// 1. Hapus slash di awal path jika ada
	// 2. Gabungkan baseDir dengan path untuk mendapatkan full path
	// 3. Periksa apakah request mengarah ke API, jika ya, return NotFound
	// 4. Periksa apakah file ada, jika tidak, redirect ke halaman utama
	// 5. Jika file ada, gunakan fileServer untuk melayani file
}