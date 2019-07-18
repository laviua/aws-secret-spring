package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"strings"
)

func main() {
	s := session.Must(session.NewSession())
	sm := secretsmanager.New(s)
	result, err := sm.ListSecrets(&secretsmanager.ListSecretsInput{})
	if err != nil {
		panic(err.Error())
	}
	configtext := "{"
	for _, element := range result.SecretList {
		sn := element.Name
		output, err := sm.GetSecretValue(&secretsmanager.GetSecretValueInput{SecretId: sn})
		if err != nil {
			panic(err.Error())
		}

		var objmap map[string]interface{}
		err = json.Unmarshal([]byte(*output.SecretString), &objmap)

		if err != nil {
			panic(err.Error())
		}

		for key, value := range objmap {
			configtext += "\"" + key + "\"" + ":" + "\"" + value.(string) + "\"" + ","
		}
	}
	configtext = strings.TrimSuffix(configtext, ",")
	configtext += "}"
	fmt.Print(configtext)
}
