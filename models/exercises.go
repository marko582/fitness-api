package models

type Exercise struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Image    string `json:"image_url"`
    Gender   string `json:"gender"`
    Type     string `json:"exercise_type"`
    Overview string `json:"overview"`
}