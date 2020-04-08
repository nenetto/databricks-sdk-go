package main

import (
	"encoding/json"
	"flag"
	"github.com/nenetto/databricks-sdk-go/api/instanceprofile"
	"github.com/nenetto/databricks-sdk-go/client"
	"github.com/nenetto/databricks-sdk-go/models"
	"io/ioutil"
)

func main() {
	flag.Parse() // required to suppress warnings from glog

	// TODO: use a better logging library
	//flag.Lookup("logtostderr").Value.Set("true")
	//flag.Lookup("stderrthreshold").Value.Set("INFO")

	secrets := loadSecrets()

	cl, err := client.NewClient(client.Options{
		Domain: &secrets.Domain, Token: &secrets.Token,
	})
	if err != nil {
		panic(err)
	}

	endpoint := instanceprofile.Endpoint{
		Client: cl,
	}

	request := models.InstanceprofileAddRequest{
		InstanceProfileArn: secrets.InstanceprofileArn,
	}

	errAdd := endpoint.Add(&request)
	if errAdd != nil {
		panic(errAdd)
	}


}

type secrets struct {
	Domain      string `json:"domain"`
	Token       string `json:"token"`
	InstanceprofileArn string `json:"instance_profile_arn"`
}

func loadSecrets() *secrets {
	content, err := ioutil.ReadFile("../secrets.json")
	if err != nil {
		panic(err)
	}

	var sc secrets
	err = json.Unmarshal(content, &sc)
	if err != nil {
		panic(err)
	}

	return &sc
}
