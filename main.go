package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nic-jennings/go-serverless-lambda/config"
)

var SetupError error

var dbCon string

func init() {
	dbCon = os.Getenv("databaseConnection")
	SetupError = config.ConnectToDb(dbCon)
}

func GenerateResponse(Body string, Code int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{Body: Body, StatusCode: Code, Headers: map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "POST,OPTIONS",
		"Access-Control-Allow-Headers": "X-Amz-Date,X-Api-Key,X-Amz-Security-Token,X-Requested-With,X-Auth-Token,Referer,User-Agent,Origin,Content-Type,Authorization,Accept,Access-Control-Allow-Methods,Access-Control-Allow-Origin,Access-Control-Allow-Headers",
	}}
}

func HandleRequest(_ context.Context, request events.LambdaFunctionURLRequest) (events.APIGatewayProxyResponse, error) {
	if SetupError != nil {
		// We can query our database as follows:
		// defer db.Close()
		// results, err := db.Query("SELECT id, name FROM tags")
		return GenerateResponse("Hello World", 200), nil
	} else {
		// throw error if our database connection fails
		return GenerateResponse("Error connecting to DB", 500), nil
	}
}

func main() {
	lambda.Start(HandleRequest)
}
