package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"sync"
)

var wg sync.WaitGroup;	

func postRequest() {

	defer wg.Done();

	// Define the URL to send the POST request to
	url := "http://localhost:5002/place-order"

	// Define the request body as a byte array
	requestBody := []byte(`{
		"owner_id": 4
}`)
	// Create a new HTTP request with the POST method and request body
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Create a new HTTP client
	client := &http.Client{}

	// Send the POST request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Print the response status code and body
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(responseBody))
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go postRequest();		
	}
	wg.Wait();
}
