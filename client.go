package zonar

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	Customer string
	Username string
	Password string
}

func New(customer, username, password string) *Client {
	return &Client{
		Customer: customer,
		Username: username,
		Password: password,
	}
}

func (c *Client) request(form url.Values) ([]byte, error) {
	client := &http.Client{}

	form.Add("username", c.Username)
	form.Add("password", c.Password)
	form.Add("customer", c.Customer)

	req, err := http.NewRequest("POST", host, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Kamion Zonar-Client "+ClientVersion)

	// dump, err := httputil.DumpRequest(req, true)
	// if err != nil {
	// 	return nil, err
	// }

	// log.Println(string(dump))

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}
