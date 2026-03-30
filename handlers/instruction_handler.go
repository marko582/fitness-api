package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"fitness-api/db"
	"fitness-api/models"
)

func GetInstructionsByExerciseID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	exerciseID, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid exercise ID", http.StatusBadRequest)
		return
	}

	rows, err := db.DB.Query(`
		SELECT id, exercise_id, step_number, description
		FROM public.instructions
		WHERE exercise_id = $1
		ORDER BY step_number ASC
	`, exerciseID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var instructions []models.Instruction

	for rows.Next() {
		var ins models.Instruction

		err := rows.Scan(
			&ins.ID,
			&ins.ExerciseID,
			&ins.StepNumber,
			&ins.Description,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		instructions = append(instructions, ins)
	}

	if instructions == nil {
		instructions = []models.Instruction{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(instructions)
}
