package services

import (
    "encoding/json"
    "net/http"
	"os"
    "github.com/mateopolci/AmbulanciaYa/src/models"
)

// Servicio de Los Pinos
func GetDatosLosPinos() models.DatosLosPinos {
	apiURL := os.Getenv("LOS_PINOS_API_URL")
	resp, err := http.Get(apiURL)
    if err != nil {
        return models.DatosLosPinos{
            Msg: "Error al conectarse con el servicio de Los Pinos",
        }
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return models.DatosLosPinos{
            Msg: "Error: No se pudo obtener los datos de Los Pinos",
        }
    }

    var data models.DatosLosPinos
    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        return models.DatosLosPinos{
            Msg: "Error en decodificar la data del servicio de Los Pinos",
        }
    }

    return data
}
