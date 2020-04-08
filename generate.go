package sdk

//go:generate java -jar swagger/swagger-codegen-cli-2.2.1.jar generate --config swagger/config.json -Dmodels -l go -i swagger/databricks-2.0.yaml -o models
//go:generate gofmt -s -w models


