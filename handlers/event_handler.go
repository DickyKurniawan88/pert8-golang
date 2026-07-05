package handlers

import (
	// "database/sql"
	// "encoding/json"
	// "log"
	"net/http"
	// "strconv"
	// "strings"
	// "time"
	//
	// "namafolder/configs"
	// "namafolder/models"
	// "namafolder/utils"
)

// HandleEvents mengarahkan request ke handler yang sesuai berdasarkan method HTTP
func HandleEvents(w http.ResponseWriter, r *http.Request) {
	// TODO: Implementasikan routing berdasarkan HTTP method
	// 1. Ekstrak ID dari URL jika ada
	// 2. Routing ke fungsi yang sesuai (GET, POST, PUT, DELETE)
	// 3. Return error jika method tidak didukung
}

// TODO: Implementasikan fungsi handleGetEvents untuk mengambil data event (semua atau spesifik)
// Fungsi menerima parameter w http.ResponseWriter dan id int
func handleGetEvents(w http.ResponseWriter, id int) {
	// TODO: Implementasi GET untuk semua events atau event dengan ID tertentu
	// Jika id == 0, ambil semua event, jika tidak, ambil event dengan ID tersebut
}

// TODO: Implementasikan fungsi handlePostEvent untuk menambahkan event baru
// Fungsi menerima parameter w http.ResponseWriter dan r *http.Request
func handlePostEvent(w http.ResponseWriter, r *http.Request) {
	// TODO: Implementasi POST untuk menambahkan event baru
	// 1. Decode JSON dari request body ke struct Event
	// 2. Validasi data event
	// 3. Set nilai default dan timestamp
	// 4. Simpan ke database
	// 5. Return response sukses atau error
}

// TODO: Implementasikan fungsi handleUpdateEvent untuk memperbarui data event
// Fungsi menerima parameter w http.ResponseWriter, r *http.Request, dan id int
func handleUpdateEvent(w http.ResponseWriter, r *http.Request, id int) {
	// TODO: Implementasi PUT untuk memperbarui event
	// 1. Decode JSON dari request body ke struct Event
	// 2. Validasi data event
	// 3. Update data di database
	// 4. Return response sukses atau error
}

// TODO: Implementasikan fungsi handleDeleteEvent untuk menghapus event
// Fungsi menerima parameter w http.ResponseWriter dan id int
func handleDeleteEvent(w http.ResponseWriter, id int) {
	// TODO: Implementasi DELETE untuk menghapus event
	// 1. Hapus event dari database berdasarkan ID
	// 2. Return response sukses atau error
}