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

	url := "http://localhost:5002/place-order"

	requestBody := []byte(`{
		"owner_id": 4
}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

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
