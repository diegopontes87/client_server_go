package api

import (
	"client_server/server/entities"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func CreateServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cotation", getCotation)
	http.ListenAndServe(":8080", mux)
}

func getCotation(w http.ResponseWriter, r *http.Request) {

	client := &http.Client{
		Timeout: 200 * time.Millisecond,
	}

	req, err := http.NewRequest("GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error during request execution: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error trying to read body: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var cotation entities.Cotation
	err = cotation.UnmarshalCotation(body, &cotation)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(cotation.USDBRL)
}