package tests

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func testGetRequest(baseUrl string, urlParams string, expectedStatus int) {
	url := fmt.Sprintf("%s/%s", baseUrl, urlParams)
	r, _ := http.Get(url)
	if r.StatusCode != expectedStatus {
		fmt.Printf("expected %d, got %d", expectedStatus, r.StatusCode)
		panic(errors.New(""))
	}
}

func postRequest(baseUrl string, urlParams string, expectedStatus int, body []byte) {
	url := fmt.Sprintf("%s/%s", baseUrl, urlParams)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
        fmt.Println("Error creating request:", err)
        return
    }
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
        fmt.Println("Error sending request:", err)
        return
    }
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
        fmt.Println("Error reading response body:", err)
        return
    }
	_ = respBody
	if resp.StatusCode != expectedStatus {
		fmt.Printf("expected %d, got %d", expectedStatus, resp.StatusCode)
		panic(errors.New(""))
	}
}