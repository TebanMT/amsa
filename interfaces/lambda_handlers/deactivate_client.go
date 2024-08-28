package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/TebanMT/amsa/app/use_cases"
	"github.com/TebanMT/amsa/infrastructure/db"
	"github.com/TebanMT/amsa/infrastructure/db/repository_impl"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func deactivateClientHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	idClient, _ := strconv.Atoi(request.PathParameters["idClient"])
	database := db.GetDatabaseInstance()
	clientRepo := repository_impl.NewClientRepository(database)
	deactivateClient := use_cases.NewClientUseCases(clientRepo)

	client, err := deactivateClient.DeactivateClient(ctx, idClient)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
			}}, err
	}

	response := map[string]interface{}{
		"deactivate": client,
	}

	responseBody, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
			}}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
		Body: string(responseBody),
	}, nil
}

func main() {
	lambda.Start(deactivateClientHandler)
}
