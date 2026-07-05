package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"eventrealm_release/configs" // UBAH 12345 DENGAN NPM KALIAN
	"eventrealm_release/models"  // UBAH 12345 DENGAN NPM KALIAN
	"eventrealm_release/utils"   // UBAH 12345 DENGAN NPM KALIAN
)

func HandleEvents(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/events/")
	var id int
	if idStr != "" && idStr != "/api/events" {
		id, _ = strconv.Atoi(idStr)
	}

	switch r.Method {
	case http.MethodGet:
		handleGetEvents(w, id)
	case http.MethodPost:
		handlePostEvent(w, r)
	case http.MethodPut:
		handleUpdateEvent(w, r, id)
	case http.MethodDelete:
		handleDeleteEvent(w, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGetEvents(w http.ResponseWriter, id int) {
	var events []models.Event

	if id == 0 {
		rows, err := configs.DB.Query("SELECT * FROM events")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var e models.Event
			rows.Scan(&e.ID, &e.Title, &e.Organizer, &e.Description, &e.Date, &e.Location, &e.Price, &e.Capacity, &e.ImageURL, &e.Status, &e.CreatedAt, &e.UpdatedAt)
			events = append(events, e)
		}
	} else {
		var e models.Event
		err := configs.DB.QueryRow("SELECT * FROM events WHERE id_event = ?", id).Scan(
			&e.ID, &e.Title, &e.Organizer, &e.Description, &e.Date, &e.Location, &e.Price, &e.Capacity, &e.ImageURL, &e.Status, &e.CreatedAt, &e.UpdatedAt)
		if err != nil {
			http.Error(w, "Event tidak ditemukan", http.StatusNotFound)
			return
		}
		events = append(events, e)
	}

	utils.RespondJSON(w, events)
}

func handlePostEvent(w http.ResponseWriter, r *http.Request) {
	var e models.Event
	json.NewDecoder(r.Body).Decode(&e)
	e.Status = "upcoming"
	e.Date = time.Now() // default jika kosong

	_, err := configs.DB.Exec("INSERT INTO events (title, organizer, description, date, location, price, capacity, image_url, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		e.Title, e.Organizer, e.Description, e.Date, e.Location, e.Price, e.Capacity, e.ImageURL, e.Status)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.RespondJSON(w, map[string]string{"message": "Event berhasil ditambahkan"})
}

func handleUpdateEvent(w http.ResponseWriter, r *http.Request, id int) {
	var e models.Event
	json.NewDecoder(r.Body).Decode(&e)

	_, err := configs.DB.Exec("UPDATE events SET title=?, organizer=?, description=?, location=?, price=?, capacity=?, image_url=?, status=? WHERE id_event=?",
		e.Title, e.Organizer, e.Description, e.Location, e.Price, e.Capacity, e.ImageURL, e.Status, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.RespondJSON(w, map[string]string{"message": "Event berhasil diupdate"})
}

func handleDeleteEvent(w http.ResponseWriter, id int) {
	_, err := configs.DB.Exec("DELETE FROM events WHERE id_event = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.RespondJSON(w, map[string]string{"message": "Event berhasil dihapus"})
}
