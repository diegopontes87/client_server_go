package main

import (
	"client_server/client/entities"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	req, err := http.Get("http://localhost:8080/cotation")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during request: %v\n", err)
	}

	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading response%v\n", err)
	}

	var cotation entities.ClientCotation
	err = json.Unmarshal(res, &cotation)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during parse%v\n", err)
	}

	fmt.Println(cotation)
}
