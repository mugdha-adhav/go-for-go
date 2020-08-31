package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Cortex struct {
	URL      string
	Query    string
	Request  Request
	Response Response
}

type Request struct {
	Endpoint string
	Body     string
	Method   string
	Auth     Auth
}

type Auth struct {
	User   string
	Passwd string
}

type Response struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
	Error  string `json:"error"`
}

type Data struct {
	ResultType string   `json:"resultType"`
	Result     []Result `json:"result"`
}

type Result struct {
	Metric Metric        `json:"metric"`
	Value  []interface{} `json:"value"`
}

// Depends on the metrics to be exported
type Metric struct {
}

const urlPath = "/api/prom/api/v1/query"

func parseInputs(c *Cortex) error {
	flag.StringVar(&c.URL, "url", "http://cortex.com", "Cortex URL")
	flag.StringVar(&c.Request.Auth.User, "user", "default", "Cortex URL")
	flag.StringVar(&c.Request.Auth.Passwd, "password", "password", "Cortex URL")

	// Run prometheus query **up** by default
	flag.StringVar(&c.Query, "query", "up", "Prometheus query to be executed on Cortex")

	flag.Parse()

	return nil
}

func createRequest(c *Cortex) error {
	// Validate Cortex URL
	u, err := url.ParseRequestURI(c.URL)
	if err != nil {
		return err
	}
	c.URL = u.String()

	// Set request endpoint with Cortex URL and path
	c.Request.Endpoint = fmt.Sprintf("%s%s", c.URL, urlPath)

	// Set request body with Cortex query
	c.Request.Body = fmt.Sprintf("query=%s", c.Query)

	// Set request type as POST by default
	c.Request.Method = "POST"

	return nil
}

func sendBasicAuthRequest(in Request) ([]byte, error) {
	// Create new http client and request with basic auth header
	client := &http.Client{}

	req, err := http.NewRequest(in.Method, in.Endpoint, strings.NewReader(in.Body))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(in.Auth.User, in.Auth.Passwd)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send request to Cortex
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Read response body into byte array
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func parseResponse(res *Response, d []byte) error {
	// Unmarshalling json response to struct
	if err := json.Unmarshal(d, &res); err != nil {
		return err
	}

	if res.Status == "error" {
		return fmt.Errorf(res.Error)
	}

	log.Printf("Response receieved from Cortex successfully\n%+v", res.Data)

	return nil
}

func main() {
	var obj Cortex

	// Parsing all the inputs from command line
	if err := parseInputs(&obj); err != nil {
		log.Fatalf("Error occured while parsing input parameters. Error is: %s", err.Error())
	}

	// Populate request object for Cortex
	if err := createRequest(&obj); err != nil {
		log.Fatalf("Error occured while parsing input parameters. Error is: %s", err.Error())
	}

	// Send API request to Cortex with basic authentication
	data, err := sendBasicAuthRequest(obj.Request)
	if err != nil {
		log.Fatalf("Error occured while sending request to Cortex. Error is: %s", err.Error())
	}

	// Unmarshalling and printing response from Cortex
	if err := parseResponse(&obj.Response, data); err != nil {
		log.Printf("Error occured while parsing Cortex response. Error is: %s", err.Error())
	}
}
