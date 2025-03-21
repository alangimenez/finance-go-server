package clients

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/alangimenez/finance-go-server/model"
)

// Cliente HTTP con timeout
var client = &http.Client{Timeout: 15 * time.Second}

// Función para hacer una solicitud GET
func GetRealTimeQuotes() model.Quotes {
	url := "https://finance-back-server.onrender.com/public/lastvalue/actualquotes"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Agregar headers si es necesario
	req.Header.Set("Accept", "application/json")

	// Hacer la petición
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Leer respuesta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Parsear JSON
	var apiResponse model.Quotes
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		log.Fatal(err)
	}

	return apiResponse
}
