package lotr

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetCharacters() ([]Character, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/character", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	characters := ChacarterList{}
	err = json.Unmarshal(body, &characters)
	if err != nil {
		return nil, err
	}

	return characters.Docs, nil
}
