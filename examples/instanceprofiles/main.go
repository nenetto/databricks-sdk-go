package main

import (
	"encoding/json"
	"flag"
	"fmt"
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

	fmt.Println("Creating Client")
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
	fmt.Println("----------------------------------------")

	//request := models.InstanceprofileAddRequest{
	//	InstanceProfileArn: secrets.InstanceprofileArn,
	//}
	//
	//errAdd := endpoint.Add(&request)
	//if errAdd != nil {
	//	panic(errAdd)
	//}

	resp, err := endpoint.List()
	if err == nil {
		for _, ip := range resp.InstanceProfiles {
			fmt.Println(ip.InstanceProfileArn)
		}
	}else {
		fmt.Println(err)
	}
	fmt.Println("----------------------------------------")

	request := models.InstanceprofileInfo{
		 InstanceProfileArn: "arn:aws:iam::685319344187:instance-profile/tf.read-ip.databricks.canberrametro.cafrds.pro",
	}
	err = endpoint.Get(&request)
	if err == nil {
		fmt.Printf("Arn found: %s", request.InstanceProfileArn)
	}else {
		fmt.Println(err)
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
