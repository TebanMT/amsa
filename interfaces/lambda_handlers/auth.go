package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TebanMT/amsa/app/use_cases"
	"github.com/TebanMT/amsa/infrastructure/authentication"
	"github.com/TebanMT/amsa/infrastructure/db"
	"github.com/TebanMT/amsa/infrastructure/db/repository_impl"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Request es una estructura para manejar la entrada de la función Lambda.
type Request struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Response es una estructura para enviar respuestas desde la función Lambda.
type Response struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

func auth(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Leer el nombre de usuario y la contraseña de los query params
	fmt.Printf("Procesando solicitud: %s\n", request)

	username := request.QueryStringParameters["username"]
	password := request.QueryStringParameters["password"]

	// Inicializar la conexión a la base de datos y crear el repositorio
	database := db.GetDatabaseInstance()
	userRepository := repository_impl.NewUserRepository(database)

	//
	authService := authentication.NewJWTTokenService("TEST")

	// Crear el caso de uso y ejecutarlo
	authenticateUserUseCase := use_cases.NewAuthenticateUser(userRepository, authService)
	user, err := authenticateUserUseCase.Authenticate(ctx, username, password)
	if err != nil {
		// Manejar errores, como credenciales incorrectas o problemas de conexión a la base de datos
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
	}

	// Manejo de la autenticación fallida
	if err != nil {
		fmt.Println("Error de autenticación:", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusUnauthorized,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
			},
			Body: "Error de ..",
		}, nil
	}

	// Si la autenticación es exitosa, user no será nil
	response := map[string]interface{}{
		"authenticated": user != nil,
		"user":          user,
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
	lambda.Start(auth)
}
