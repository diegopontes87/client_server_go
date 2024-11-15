package main

import (
	"client_server/client/entities"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	SetTimeOut(ctx)
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotation", nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error during request:", err)
		panic(err)
	}

	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response%v\n:", err)
	}

	var bid entities.ClientBid
	err = json.Unmarshal(res, &bid)

	if err != nil {
		log.Printf("Error during parse%v\n:", err)
	}

	fmt.Println(bid)
}

func SetTimeOut(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Hotel book cancelled. Timeout reached.")
		return
	case <-time.After(300 * time.Millisecond):
		fmt.Println("Hotel Booked!")
	}
}
