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

func VincularFacturaHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Parsear el cuerpo de la solicitud para obtener los datos del cliente
	fmt.Printf("Procesando solicitud: %s\n", request)
	idCompany, _ := strconv.Atoi(request.QueryStringParameters["idCompany"])
	id, _ := strconv.Atoi(request.QueryStringParameters["id"])

	// Inicializar repositorio y caso de uso
	database := db.GetDatabaseInstance()
	facturaRepo := repository_impl.NewBillRepository(database)
	saleRepo := repository_impl.NewSaleRepository(database)
	createFacturaUseCase := use_cases.NewBillUseCase(facturaRepo, use_cases.WithBillRepoAndSaleRepo(saleRepo))

	// update
	res, err := createFacturaUseCase.Vincular(ctx, idCompany, id)
	if err != nil || !res {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
	}

	response := map[string]interface{}{
		"updated": true,
		"message": "Vinculado Correctamente",
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
	lambda.Start(VincularFacturaHandler)
}
