package utils

import (
	// "encoding/json"
	"net/http"
)

// Fungsi ini menerima parameter http.ResponseWriter dan data interface{}
// kemudian mengirim response dalam format JSON
func RespondJSON(w http.ResponseWriter, data interface{}) {
	// TODO: Implementasi pengiriman response dalam format JSON
	// 1. Set header Content-Type menjadi application/json
	// 2. Encode data ke format JSON dan kirim ke writer
}