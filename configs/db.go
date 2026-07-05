package configs

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Variabel global untuk instance database
var DB *sql.DB

// ConnectDB melakukan koneksi ke database MySQL
func ConnectDB() {
    // TODO: Ubah string koneksi di bawah sesuai konfigurasi di lab masing - masing
    var err error

    DB, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/event_realm?parseTime=true")
    if err != nil {
        log.Fatal("Gagal membuka koneksi database:", err)
    }

    // Konfigurasi pool koneksi (max open conns, max idle conns, conn max lifetime)
    DB.SetMaxOpenConns(10)
    DB.SetMaxIdleConns(5)
    DB.SetConnMaxLifetime(time.Minute * 5)

    // Verifikasi koneksi dengan ping
    err = DB.Ping()
    if err != nil {
        log.Fatal("Gagal melakukan ping ke database:", err)
    }

    // Verifikasi tabel events dengan query sederhana
    _, err = DB.Query("SELECT 1 FROM events LIMIT 1")
    if err != nil {
        log.Fatal("Gagal menjalankan query pada tabel events:", err)
    }

    log.Println("Berhasil terhubung ke database")
}