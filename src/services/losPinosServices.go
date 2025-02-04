package services

import(
	"encoding/json"
	"os"
	"github.com/mateopolci/AmbulanciaYa/src/models"
)

// Simulacion de servicio de Los Pinos
func GetDatosLosPinos() models.DatosLosPinos {
	file, _ := os.Open("datosLosPinos.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	data := models.DatosLosPinos{}
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	return data
}