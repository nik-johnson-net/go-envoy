package envoy

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	// ErrNotOK is returned if any of the Envoy APIs does not return a 200
	ErrNotOK = errors.New("server did not return 200")
)

// Client provides the API for interacting with the Envoy APIs
type Client struct {
	address string
	client  *http.Client
}

// NewClient creates a new Client that will talk to an Envoy unit at *address*, creating its own http.Client underneath.
func NewClient(address string) *Client {
	client := &http.Client{}

	return &Client{
		address: address,
		client:  client,
	}
}

// NewClientWithHTTP creates a new Client that will talk to an Envoy unit at *address* using the provided http.Client.
func NewClientWithHTTP(address string, client *http.Client) *Client {
	return &Client{
		address: address,
		client:  client,
	}
}

func (c *Client) get(url string, response interface{}) error {
	resp, err := c.client.Get(fmt.Sprintf("http://%s%s", c.address, url))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return ErrNotOK
	}

	return json.NewDecoder(resp.Body).Decode(response)
}

// Inventory returns the list of parts installed in the system and registered with the Envoy unit
func (c *Client) Inventory() ([]Inventory, error) {
	var inventory []Inventory
	err := c.get("/inventory.json?deleted=1", &inventory)
	return inventory, err
}

// Production returns the current data for Production and Consumption sensors, if equipped.
func (c *Client) Production() (Production, error) {
	var production Production
	err := c.get("/production.json?details=1", &production)
	return production, err
}
