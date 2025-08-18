package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"github.com/mahendrankrishnan/app1/models"
)

// You would move the appointments map and nextID into a service layer later
var appointments = make(map[int]models.Appointment)
var nextID = 1

// Add a global DB variable (or inject via function/struct for better design)
var DB *sql.DB

func SetDB(db *sql.DB) {
	DB = db
}

// CreateAppointment godoc
// @Summary Create a new appointment
// @Accept json
// @Produce json
// @Param appointment body models.Appointment true "Appointment data"
// @Success 201 {object} models.Appointment
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router / [post]
func CreateAppointment(w http.ResponseWriter, r *http.Request) {
	var appt models.Appointment
	if err := json.NewDecoder(r.Body).Decode(&appt); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert into the database
	query := `INSERT INTO appointments (appt_name,appt_type, appt_desc, appt_time) VALUES ($1, $2, $3,$4) RETURNING id`
	err := DB.QueryRow(query, appt.ApptName, appt.ApptType, appt.ApptDesc, appt.ApptTime).Scan(&appt.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the created appointment (with ID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(appt)
}
func UpdateAppointment(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var appt models.Appointment
	if err := json.NewDecoder(r.Body).Decode(&appt); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	appt.ID = id
	appointments[id] = appt
	w.WriteHeader(http.StatusOK)
}

// GetAppointments godoc
// @Summary List appointments
// @Produce json
// @Success 200 {array} models.Appointment
// @Router / [get]
func GetAppointments(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT id, appt_name, appt_desc, appt_time FROM appointments")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var list []models.Appointment
	for rows.Next() {
		var appt models.Appointment
		if err := rows.Scan(&appt.ID, &appt.ApptName, &appt.ApptDesc, &appt.ApptTime); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		list = append(list, appt)
	}
	json.NewEncoder(w).Encode(list)
}

//	func GetAppointment(w http.ResponseWriter, r *http.Request) {
//		idParam := chi.URLParam(r, "id")
//		id, err := strconv.Atoi(idParam)
//		if err != nil {
//			http.Error(w, "invalid id", http.StatusBadRequest)
//			return
//		}
//		appt, exists := appointments[id]
//		if !exists {
//			http.NotFound(w, r)
//			return
//		}
//		json.NewEncoder(w).Encode(appt)
//	}

// GetAppointment godoc
// @Summary Get appointment by ID
// @Produce json
// @Param id path int true "Appointment ID"
// @Success 200 {object} models.Appointment
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /{id} [get]

func GetAppointment(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var appt models.Appointment
	query := `SELECT id, appt_name, appt_type, appt_desc, appt_time FROM appointments WHERE id = $1`
	err = DB.QueryRow(query, id).Scan(&appt.ID, &appt.ApptName, &appt.ApptType, &appt.ApptDesc, &appt.ApptTime)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(appt)
}

//	func DeleteAppointment(w http.ResponseWriter, r *http.Request) {
//		idParam := chi.URLParam(r, "id")
//		id, err := strconv.Atoi(idParam)
//		if err != nil {
//			http.Error(w, "invalid id", http.StatusBadRequest)
//			return
//		}
//		if _, exists := appointments[id]; !exists {
//			http.NotFound(w, r)
//			return
//		}
//		delete(appointments, id)
//		w.WriteHeader(http.StatusNoContent)
//	}

// DeleteAppointment godoc
// @Summary delete appointment by ID
// @Produce json
// @Param id path int true "Appointment ID"
// @Success 200 {object} models.Appointment
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /{id} [delete]

func DeleteAppointment(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	result, err := DB.Exec("DELETE FROM appointments WHERE id = $1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
