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
        //Debug
        fmt.Println("Error cargando archivo .env con godotenv.Load()")
		return models.DatosVeloway{}
	}

    velowayApiUrl := os.Getenv("VELOWAY_API_URL") + "/" + telefono
	apiKey := os.Getenv("API_KEY")
    
	if velowayApiUrl == "" || apiKey == "" {
        //Debug
        fmt.Println("Alguna de las variables de entorno es string vacio")
		return models.DatosVeloway{}
	}

	req, err := http.NewRequest("GET", velowayApiUrl, nil)
	if err != nil {
        //Debug
        fmt.Println("Error desesctructurando la request")
		return models.DatosVeloway{}
	}
    
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))
    
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
        //Debug
        fmt.Println("Error en el client.Do(req)")
        return models.DatosVeloway{}
	}
	defer resp.Body.Close()
    
	if resp.StatusCode != http.StatusOK {
        //Debug
        fmt.Println("La api no devuelve 200")
		return models.DatosVeloway{}
	}
    
	var data models.DatosVeloway
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        //Debug
        fmt.Println("Error decodificando el JSON de veloway")
		return models.DatosVeloway{}
	}

    //Debug
    fmt.Println("Este es el JSON que responde Veloway en velowayServices", data)
    
	return data
}
