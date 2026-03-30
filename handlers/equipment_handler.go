package handlers

import (
	"encoding/json"
	"net/http"

	"fitness-api/db"
	"fitness-api/models"
)

func GetAllEquipments(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(`
		SELECT id, name
		FROM public.equipments;
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var equipments []models.Equipment

	for rows.Next() {
		var eq models.Equipment

		err := rows.Scan(
			&eq.ID,
			&eq.Name,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		equipments = append(equipments, eq)
	}

	if equipments == nil {
		equipments = []models.Equipment{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(equipments)
}
