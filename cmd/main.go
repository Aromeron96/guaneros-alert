package main

import (
	"io"
	"log"
	"net/http"
)






func main() {

	client := &http.Client{}
//Este request permite añadir headers
	
	const apiKey = ""

	req, err := http.NewRequest ("GET", "https://api.coingecko.com/api/v3/simple/price?vs_currencies=usd&ids=bitcoin&names=Bitcoin&symbols=btc", nil)

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("x-cg-demo-api-key", apiKey)
	resp, err := client.Do(req)
	
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error en la respuesta %d ", resp.StatusCode)
	}

	body ,err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
	}
	
  // parseamos el contenido del body a string, esto es un parche, lo suy seria parsear el json y manejar los datos con un struct
	log.Println(string(body))





}