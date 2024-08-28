package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TebanMT/amsa/app/use_cases"
	"github.com/TebanMT/amsa/domain/entities"
	"github.com/TebanMT/amsa/infrastructure/db"
	"github.com/TebanMT/amsa/infrastructure/db/repository_impl"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func createSupplierHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Procesando solicitud: %s\n", request)
	var supplier entities.Supplier
	if err := json.Unmarshal([]byte(request.Body), &supplier); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
			}}, err
	}

	// Inicializar repositorio y caso de uso
	database := db.GetDatabaseInstance()
	supplierRepo := repository_impl.NewSupplierRepository(database)
	createSupplierUseCase := use_cases.NewSupplierUseCase(supplierRepo)

	// Crear el Proveedor
	if err := createSupplierUseCase.RegisterSupplierIntoDB(ctx, &supplier); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
	}

	response := map[string]interface{}{
		"created": true,
		"message": "Proveedor Creado Correctamente",
	}

	responseBody, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
			},
		}, err
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
	lambda.Start(createSupplierHandler)
}
