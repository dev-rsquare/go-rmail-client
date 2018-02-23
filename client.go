package rmail

import "net/http"
import "encoding/json"
import "bytes"
import "fmt"
import "io/ioutil"

type Client struct {
	Env           RMailEnv
	Host          string
	Authorization string
}

func InitClient(envStr string) *Client {

	env := findRMailEnv(envStr)

	return &Client{
		Env:           env,
		Host:          findRMailHost(env),
		Authorization: findRMailAuthorization(env),
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
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)

	fmt.Printf("%v", string(b))
	fmt.Printf("%v", err)
}
