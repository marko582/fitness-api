package models

type Instruction struct {
	ID          int    `json:"id"`
	ExerciseID  int    `json:"exercise_id"`
	StepNumber  int    `json:"step_number"`
	Description string `json:"description"`
}
