package secretm

import (
	"encoding/json"
	"fmt"
	"herralmayoruser/awsgo"
	"herralmayoruser/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(nameSecret string) (models.SecretRDSJson, error) {
	var dataSecret models.SecretRDSJson

	fmt.Println(" > Pido Screto: ", nameSecret)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)

	key, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nameSecret),
	})

	if err != nil {
		fmt.Println(err.Error())
		return dataSecret, err
	}

	json.Unmarshal([]byte(*key.SecretString), &dataSecret)

	fmt.Println(" > Lectura Secret OK" + nameSecret)

	return dataSecret, nil
}
