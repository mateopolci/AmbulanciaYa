package services

import (
	"encoding/json"
	"net/http"
	"os"
	"github.com/mateopolci/AmbulanciaYa/src/models"
)

// Servicio de Veloway
func GetDatosVeloway() models.DatosVeloway {
	velowayApiUrl := os.Getenv("VELOWAY_API_URL")
	resp, err := http.Get(velowayApiUrl)
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
