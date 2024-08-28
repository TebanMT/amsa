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

func createUserHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Parsear el cuerpo de la solicitud para obtener los datos del cliente
	fmt.Printf("Procesando solicitud: %s\n", request)
	var user entities.User
	if err := json.Unmarshal([]byte(request.Body), &user); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
			}}, err
	}

	// Inicializar repositorio y caso de uso
	database := db.GetDatabaseInstance()
	userRepo := repository_impl.NewUserRepository(database)
	createUserUseCase := use_cases.NewUserCases(userRepo)

	// Crear el cliente
	if err := createUserUseCase.RegisterUserIntoDB(ctx, &user); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
	}

	response := map[string]interface{}{
		"created": true,
		"message": "Cliente Creado Correctamente",
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
	lambda.Start(createUserHandler)
}
