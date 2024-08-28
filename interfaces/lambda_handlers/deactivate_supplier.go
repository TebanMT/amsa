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

func deactivateSupplierHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	idSupplier, _ := strconv.Atoi(request.PathParameters["idSupplier"])
	database := db.GetDatabaseInstance()
	supplierRepo := repository_impl.NewSupplierRepository(database)
	createSupplierUseCase := use_cases.NewSupplierUseCase(supplierRepo)

	err := createSupplierUseCase.DeactivateSupplier(ctx, idSupplier)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
			}}, err
	}

	response := map[string]interface{}{
		"deactivate": true,
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
	lambda.Start(deactivateSupplierHandler)
}
