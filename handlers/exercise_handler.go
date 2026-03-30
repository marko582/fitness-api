package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"fitness-api/db"
	"fitness-api/models"

	"github.com/go-chi/chi/v5"
)

func GetAllExercises(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(`
        SELECT id, name, image_url, gender, exercise_type, overview 
        FROM public.exercises;
    `)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var exercises []models.Exercise

	for rows.Next() {
		var ex models.Exercise

		err := rows.Scan(
			&ex.ID,
			&ex.Name,
			&ex.Image,
			&ex.Gender,
			&ex.Type,
			&ex.Overview,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		exercises = append(exercises, ex)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exercises)
}

func GetExercises(w http.ResponseWriter, r *http.Request) {
	bodyPart := r.URL.Query().Get("bodyPart")
	equipments := r.URL.Query().Get("equipments")

	query := `
		SELECT DISTINCT e.id, e.name, e.image_url, e.gender, e.exercise_type, e.overview
		FROM public.exercises e
	`

	args := []interface{}{}
	argIndex := 1

	if bodyPart != "" {
		query += `
			JOIN public.exercise_body_parts ebp ON e.id = ebp.exercise_id
			JOIN public.body_parts bp ON bp.id = ebp.body_part_id
		`
	}

	if equipments != "" {
		query += `
			JOIN public.exercise_equipments ee ON e.id = ee.exercise_id
			JOIN public.equipments eq ON eq.id = ee.equipment_id
		`
	}

	query += " WHERE 1=1 "

	if bodyPart != "" {
		query += " AND LOWER(bp.name) = LOWER($" + strconv.Itoa(argIndex) + ")"
		args = append(args, bodyPart)
		argIndex++
	}

	if equipments != "" {
		query += " AND LOWER(eq.name) = LOWER($" + strconv.Itoa(argIndex) + ")"
		args = append(args, equipments)
		argIndex++
	}

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var exercises []models.Exercise

	for rows.Next() {
		var ex models.Exercise

		err := rows.Scan(
			&ex.ID,
			&ex.Name,
			&ex.Image,
			&ex.Gender,
			&ex.Type,
			&ex.Overview,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		exercises = append(exercises, ex)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exercises)
}

func GetExerciseByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var ex models.Exercise

	err = db.DB.QueryRow(`
		SELECT id, name, image_url, gender, exercise_type, overview
		FROM public.exercises
		WHERE id = $1
	`, id).Scan(
		&ex.ID,
		&ex.Name,
		&ex.Image,
		&ex.Gender,
		&ex.Type,
		&ex.Overview,
	)

	if err == sql.ErrNoRows {
		http.Error(w, "Exercise not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ex)
}

func GetRelatedExercises(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	exerciseID, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid exercise ID", http.StatusBadRequest)
		return
	}

	rows, err := db.DB.Query(`
		SELECT e.id, e.name, e.image_url, e.gender, e.exercise_type, e.overview
		FROM public.exercise_relations er
		JOIN public.exercises e ON e.id = er.related_exercise_id
		WHERE er.exercise_id = $1
	`, exerciseID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var exercises []models.Exercise

	for rows.Next() {
		var ex models.Exercise

		err := rows.Scan(
			&ex.ID,
			&ex.Name,
			&ex.Image,
			&ex.Gender,
			&ex.Type,
			&ex.Overview,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		exercises = append(exercises, ex)
	}

	// da ne vrati null
	if exercises == nil {
		exercises = []models.Exercise{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exercises)
}
