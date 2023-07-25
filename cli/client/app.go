package client

import (
	"fmt"
	"io/ioutil"
)

func (c *APIClient) AppStatus(appID string) error {
	res, err := c.request("https://api.squarecloud.app/v2/apps/" + appID + "/status")
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// body to string
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	return nil
}
