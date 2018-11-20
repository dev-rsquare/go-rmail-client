package rmail

import "net/http"
import "encoding/json"
import "bytes"
import "fmt"
import "io/ioutil"

// Client struct store metadata and functions for executing queries
type Client struct {
	Env           Env
	Host          string
	Authorization string
}

// InitClient : Initiaize Client
// envStr is environment string ("development" or "production" or "prod")
func InitClient(envStr string, token string) *Client {

	env := findRMailEnv(envStr)

	return &Client{
		Env:           env,
		Host:          findRMailHost(env),
		Authorization: token,
	}
}

func (client *Client) sendMail(query string, vars interface{}) {

	httpClient := http.DefaultClient

	body := map[string]interface{}{
		"query":     query,
		"variables": vars,
	}

	bodyByte, err := json.Marshal(body)
	if err != nil {
		panic("RMail Request is invalid.")
	}

	req, err := http.NewRequest("POST", client.Host, bytes.NewReader(bodyByte))
	if err != nil {
		panic("RMail Request is invalid")
	}

	req.Header.Add("Authorization", client.Authorization)
	req.Header.Add("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		panic("Rmail Response is errored")
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)

	fmt.Printf("%v", string(b))
	fmt.Printf("%v", err)
}

func (client *Client) template(query string, vars interface{}) ([]byte, error) {

	httpClient := http.DefaultClient

	body := map[string]interface{}{
		"query":     query,
		"variables": vars,
	}

	bodyByte, err := json.Marshal(body)
	if err != nil {
		panic("RMail Request is invalid.")
	}

	req, err := http.NewRequest("POST", client.Host, bytes.NewReader(bodyByte))
	if err != nil {
		panic("RMail Request is invalid")
	}

	req.Header.Add("Authorization", client.Authorization)
	req.Header.Add("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		panic("Rmail Response is errored")
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
