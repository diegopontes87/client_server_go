package main

import (
	"client_server/client/entities"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotation", nil)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("Timeout occurred during client request")
		} else {
			log.Println("Error during request:", err)
		}
	} else {
		log.Println("Success getting data")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error during request: %v", err)
		panic(err)
	}

	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response: %v:", err)
	}

	var bid entities.ClientBid
	err = json.Unmarshal(res, &bid)

	if err != nil {
		log.Printf("Error during parse: %v:", err)
	}

	fmt.Println(bid)
	createFileAndSave(bid)
}

func createFileAndSave(bid entities.ClientBid) {
	var file *os.File
	filePath := "file.txt"

	if _, err := os.Stat(filePath); err == nil {
		log.Println("Opening file")
		file, err = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Failed to open file: %v", err)
		}
	} else if errors.Is(err, os.ErrNotExist) {
		log.Println("File doesnt exist")
		file, err = os.Create(filePath)
		log.Println("Creating file")
		if err != nil {
			log.Fatalf("Failed to create file: %v", err)
		}
	}
	defer file.Close()
	content := fmt.Sprintf("\nDolar: %s", bid)
	_, err := file.Write([]byte(content))
	if err != nil {
		log.Fatalf("Error trying to write in file: %v", err)
	} else {
		log.Println("Successfully wrote to file.")
	}
}
