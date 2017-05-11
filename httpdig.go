package httpdig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const apiURL = "https://dns.google.com/resolve"
const ednsSubnet = "0.0.0.0/0"
const defaultTimeout = 30 * time.Second

type Response struct {
	Status int  `json:"Status"`
	TC     bool `json:"TC"`
	RD     bool `json:"RD"`
	RA     bool `json:"RA"`
	AD     bool `json:"AD"`
	CD     bool `json:"CD"`

	Question []struct {
		Name string `json:"name"`
		Type int    `json:"type"`
	} `json:"Question"`

	Answer []struct {
		Name string `json:"name"`
		Type int    `json:"type"`
		TTL  int    `json:"TTL"`
		Data string `json:"data"`
	} `json:"Answer"`

	Authority []struct {
		Name string `json:"name"`
		Type int    `json:"type"`
		TTL  int    `json:"TTL"`
		Data string `json:"data"`
	} `json:"Authority"`

	Additional       []interface{} `json:"Additional"`
	EdnsClientSubnet string        `json:"edns_client_subnet"`
	Comment          string        `json:"Comment"`
}

func dig(host, recordType string, timeout time.Duration) ([]byte, error) {
	client := &http.Client{Timeout: timeout}

	req, _ := http.NewRequest("GET", apiURL, nil)

	query := req.URL.Query()
	query.Add("name", host)
	query.Add("type", recordType)
	query.Add("edns_client_subnet", ednsSubnet)

	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, fmt.Errorf("http error: %s", err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("error reading response: %s", err.Error())
	}

	return body, nil
}

// Query sends request to Google dns service and parses response.
// e.g: httpdig.Query("google.com", "NS")
func Query(host string, t string) (Response, error) {
	resp, err := QueryWithTimeout(host, t, defaultTimeout)
	return resp, err
}

// QueryWithTimeout allows a timeout on the request.
func QueryWithTimeout(host, t string, timeout time.Duration) (Response, error) {
	resp, err := dig(host, t, timeout)
	if err != nil {
		return Response{}, err
	}

	response := Response{}
	err = json.Unmarshal(resp, &response)

	if err != nil {
		return Response{}, err
	}

	return response, nil
}
