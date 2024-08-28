package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/TebanMT/amsa/app/use_cases"
	"github.com/TebanMT/amsa/infrastructure/db"
	"github.com/TebanMT/amsa/infrastructure/db/repository_impl"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func getSuppliersHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Printf("Procesando solicitud: %s\n", request)
	idSupplier, _ := strconv.Atoi(request.QueryStringParameters["idCompany"])
	database := db.GetDatabaseInstance()
	supplierRepo := repository_impl.NewSupplierRepository(database)
	createSupplierUseCase := use_cases.NewSupplierUseCase(supplierRepo)

	suppliers, err := createSupplierUseCase.GetAllSuppliersByCompany(ctx, idSupplier)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
			}}, err
	}

	responseBody, err := json.Marshal(suppliers)
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
	lambda.Start(getSuppliersHandler)
}
