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
	form.Add("username", c.Username)
	form.Add("password", c.Password)
	form.Add("customer", c.Customer)

	req, err := http.NewRequest("POST", host, strings.NewReader(form.Encode()))
	defer req.Body.Close()

	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(req.Body)
}
