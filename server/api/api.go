package api

import (
	"client_server/server/database"
	"client_server/server/entities"
	"context"
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
	ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
	defer cancel()

	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
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

	var cotation entities.ServerCotation
	err = cotation.UnmarshalCotation(body, &cotation)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	dbCtx, dbCancel := context.WithTimeout(r.Context(), 10*time.Millisecond)
	defer dbCancel()

	database.InsertCotation(dbCtx, &cotation)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(entities.ServerBid{Bid: cotation.USDBRL.Bid})
}
