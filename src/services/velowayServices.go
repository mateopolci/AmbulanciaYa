package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"github.com/joho/godotenv"

	"github.com/mateopolci/AmbulanciaYa/src/models"
)

func GetDatosVeloway(telefono string) models.DatosVeloway {
	err := godotenv.Load()
	if err != nil {
		return models.DatosVeloway{}
	}

    velowayApiUrl := os.Getenv("VELOWAY_API_URL") + "/" + telefono
	apiKey := os.Getenv("API_KEY")

	if velowayApiUrl == "" || apiKey == "" {
		return models.DatosVeloway{}
	}

	req, err := http.NewRequest("GET", velowayApiUrl, nil)
	if err != nil {
		return models.DatosVeloway{}
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.DatosVeloway{}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.DatosVeloway{}
	}

	var data models.DatosVeloway
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return models.DatosVeloway{}
	}

	return data
}
