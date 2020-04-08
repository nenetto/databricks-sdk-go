package instanceprofile

import (
	"encoding/json"
	"github.com/nenetto/databricks-sdk-go/client"
	"github.com/nenetto/databricks-sdk-go/models"
)

type Endpoint struct {
	Client *client.Client
}

func (c *Endpoint) List() (*models.InstanceprofileListResponse, error) {
	bytes, err := c.Client.Query("GET", "instance-profiles/list", nil)
	if err != nil {
		return nil, err
	}

	resp := models.InstanceprofileListResponse{}
	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}


func (c *Endpoint) Add(request *models.InstanceprofileAddRequest) error {
	_, err := c.Client.Query("POST", "instance-profiles/add", request)
	return err
}

func (c *Endpoint) Remove(request *models.InstanceprofileRemoveRequest) error {
	_, err := c.Client.Query("POST", "instance-profiles/remove", request)
	return err
}