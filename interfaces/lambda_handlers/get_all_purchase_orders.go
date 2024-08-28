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

func getAllPurchaseOrdersHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	idCompany, _ := strconv.Atoi(request.QueryStringParameters["idCompany"])
	database := db.GetDatabaseInstance()
	purchaseRepo := repository_impl.NewPurchaseOrderRepository(database)
	purchaseOrderUseCase := use_cases.NewPurchaseOrderUseCase(purchaseRepo)

	orders, err := purchaseOrderUseCase.GetOrdersWithDetailsAndSupliers(ctx, idCompany)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
			}}, err
	}

	responseBody, err := json.Marshal(orders)
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
	lambda.Start(getAllPurchaseOrdersHandler)
}
