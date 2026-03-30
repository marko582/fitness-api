package models

import "database/sql"

type Exercise struct {
	ID         int            `json:"id"`
	Name       string         `json:"name"`
	Image      string         `json:"image_url"`
	VideoURL   sql.NullString `json:"video_url"`
	Gender     string         `json:"gender"`
	Type       string         `json:"exercise_type"`
	Difficulty string         `json:"difficulty"`
	Overview   string         `json:"overview"`
}
