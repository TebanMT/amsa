package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/TebanMT/amsa/app/use_cases"
	"github.com/TebanMT/amsa/domain/entities"
	"github.com/TebanMT/amsa/infrastructure/db"
	"github.com/TebanMT/amsa/infrastructure/db/repository_impl"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func getClientsHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var clients []entities.Client
	var err error
	idCompany, _ := strconv.Atoi(request.QueryStringParameters["idCompany"])
	include := request.QueryStringParameters["include"]
	database := db.GetDatabaseInstance()
	clientRepo := repository_impl.NewClientRepository(database)
	getAllClientsUseCase := use_cases.NewGetAllClientsUseCase(clientRepo)

	if include == "bills" {
		clients, err = getAllClientsUseCase.GetAllClientsWithBillsAndDetailsByCompany(ctx, idCompany)
	} else {
		clients, err = getAllClientsUseCase.Execute(ctx, idCompany)
	}

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
			}}, err
	}

	responseBody, err := json.Marshal(clients)
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
	lambda.Start(getClientsHandler)
}
