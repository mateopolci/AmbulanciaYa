package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/mateopolci/AmbulanciaYa/src/models"
)

func GetDatosVeloway(telefono string) models.DatosVeloway {

    velowayApiUrl := os.Getenv("VELOWAY_API_URL") + "/" + telefono
	apiKey := os.Getenv("API_KEY")
    
	if velowayApiUrl == "" || apiKey == "" {
        fmt.Println("Alguna de las variables de entorno se esta recuperando como string vacio")
		return models.DatosVeloway{}
	}

	req, err := http.NewRequest("GET", velowayApiUrl, nil)
	if err != nil {
        fmt.Println("Error desesctructurando la request")
		return models.DatosVeloway{}	
	}
    
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))
    
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
        fmt.Println("Error en la request a la api de veloway")
        return models.DatosVeloway{}
	}
	defer resp.Body.Close()
    
	if resp.StatusCode != http.StatusOK {
        fmt.Println("La api no devuelve 200")
		return models.DatosVeloway{}
	}
    
	var data models.DatosVeloway
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        fmt.Println("Error decodificando el JSON de veloway")
		return models.DatosVeloway{}
	}
    
	return data
}
