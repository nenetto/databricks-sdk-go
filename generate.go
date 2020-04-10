package sdk

//go:generate java -jar swagger/swagger-codegen-cli-2.2.1.jar generate --config swagger/config.json -Dmodels -l go -i swagger/databricks-2.0.yaml -o models
//go:generate rm models/api_client.go models/api_response.go models/configuration.go models/.gitignore models/README.md
//go:generate gofmt -s -w models
