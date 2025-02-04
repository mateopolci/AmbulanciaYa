package services

import(
	"encoding/json"
	"os"
	"github.com/mateopolci/AmbulanciaYa/src/models"
)

// Simulacion de servicio de Los Pinos
func GetDatosLosPinos() models.DatosLosPinos {
	file, err := os.Open("src/utils/jsons/LosPinos.json")
	if err != nil {
        return models.DatosLosPinos{
            Msg: "Error loading Los Pinos JSON data",
        }
    }
	defer file.Close()
	decoder := json.NewDecoder(file)
	data := models.DatosLosPinos{}
    if err := decoder.Decode(&data); err != nil {
        return models.DatosLosPinos{
            Msg: "Error parsing weather data",
        }
    }
	return data
}