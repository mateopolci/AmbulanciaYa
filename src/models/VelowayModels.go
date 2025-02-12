package models

type DatosVeloway struct {
	Altura                 int16   `json:"altura"`
	Peso                   float32 `json:"peso"`
	EnfermedadCardiaca     *string `json:"enfermedadCardiaca"`
	EnfermedadRespiratoria *string `json:"enfermedadRespiratoria"`
	Alergias               *string `json:"alergias"`
	Epilepsia              bool    `json:"epilepsia"`
	Diabetes               bool    `json:"diabetes"`
}
