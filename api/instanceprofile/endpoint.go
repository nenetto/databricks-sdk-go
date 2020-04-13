package instanceprofile

import (
	"encoding/json"
	"fmt"
	"github.com/nenetto/databricks-sdk-go/client"
	"github.com/nenetto/databricks-sdk-go/models"
)

type Endpoint struct {
	Client *client.Client
}

func (c *Endpoint) List() (*models.InstanceprofileListResponse, error) {
	resp := models.InstanceprofileListResponse{}
	bytes, err := c.Client.Query("GET", "instance-profiles/list", nil)
	if err != nil {
		return &resp, err
	}

	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		fmt.Println(err)
		return &resp, err
	}

	return &resp, nil
}

func (c *Endpoint) Exist(request *models.InstanceprofileInfo) bool {
	found := false

	resp, _ := c.List()


	for _, ip := range resp.InstanceProfiles {
		if ip.InstanceProfileArn == request.InstanceProfileArn {
			found = true
		}
	}

	return found
}


func (c *Endpoint) Add(request *models.InstanceprofileAddRequest) error {
	_, err := c.Client.Query("POST", "instance-profiles/add", request)
	return err
}

func (c *Endpoint) Remove(request *models.InstanceprofileRemoveRequest) error {
	_, err := c.Client.Query("POST", "instance-profiles/remove", request)
	return err
}