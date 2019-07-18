package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"os"
	"strings"
)

func main() {
	appname := os.Args[1]

	s := session.Must(session.NewSession())
	sm := secretsmanager.New(s)
	result, err := sm.ListSecrets(&secretsmanager.ListSecretsInput{})
	if err != nil {
		panic(err.Error())
	}
	configtext := "{"
	for _, secretElement := range result.SecretList {

		// if app tag matched - go ahead
		if matchApplication(secretElement, appname) {

			output, err := sm.GetSecretValue(&secretsmanager.GetSecretValueInput{SecretId: secretElement.Name})
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
	}
	configtext = strings.TrimSuffix(configtext, ",")
	configtext += "}"
	fmt.Print(configtext)
}

func matchApplication(element *secretsmanager.SecretListEntry, appname string) bool {
	for _, tag := range element.Tags {
		if *tag.Key == "app" && *tag.Value != appname {
			return false
		}
	}
	return true
}
