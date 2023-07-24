package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"herralmayoruser/awsgo"
	"herralmayoruser/db"
	"herralmayoruser/models"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InitAWS()

	if !ValidParams() {
		fmt.Println("Error en los parametros. Debe enviar 'SecretName")
		err := errors.New("Error en los parametros. Debe enviar 'SecretName")
		return event, err
	}

	var data models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			data.UserEmail = att
			fmt.Println("Email: " + data.UserEmail)
		case "sub":
			data.UserUUID = att
			fmt.Println("UUID: " + data.UserUUID)
		}
	}

	err := db.ReadSecret()

	if err != nil {
		fmt.Println("Error al leer el Secret: " + err.Error())
		return event, err
	}

	err = db.SignUp(data)
	return event, err
}

func ValidParams() bool {
	var getParams bool

	_, getParams = os.LookupEnv("SECRET_NAME")

	return getParams
}
