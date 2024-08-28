package main

import (
	"context"
	"fmt"

	"github.com/TebanMT/amsa/app/use_cases"
	"github.com/TebanMT/amsa/domain/entities"
	"github.com/TebanMT/amsa/infrastructure/db"
	"github.com/TebanMT/amsa/infrastructure/db/repository_impl"
)

func main() {
	fmt.Println("Hola, mundo!")

	//LOGIN

	/*username := "johnn"
	password := "Mendiola2015*"
	fmt.Printf("Procesando solicitud: %s ..... %s\n", username, password)

	// Inicializar la conexión a la base de datos y crear el repositorio
	database := db.GetDatabaseInstance()
	userRepository := repository_impl.NewUserRepository(database)

	//
	authService := authentication.NewJWTTokenService("TEST")

	// Crear el caso de uso y ejecutarlo
	ctx := context.Background()
	authenticateUserUseCase := use_cases.NewAuthenticateUser(userRepository, authService)
	user, err := authenticateUserUseCase.Authenticate(ctx, username, password)

	//REGISTER


		username := entities.User{
			Name: "John Doe", LastName: "TEST",
			SecondLastName: "TEST", Username: "johnn",
			Password: "Mendiola2015*", PhoneNumber: "+524151266863"}
		database := db.GetDatabaseInstance()
		userRepository := repository_impl.NewUserRepository(database)
		authService := cognito.NewCognitoAuthService()
		ctx := context.Background()
		authenticateUserUseCase := use_cases.NewUserCases(userRepository, authService)
		user, err := authenticateUserUseCase.RegisterUserIntoDB(ctx, username)*/

	// verify token
	/*
		authService := authentication.NewJWTTokenService("TEST")
		ctx := context.Background()
		user, err := authService.VerifyToken(ctx, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDUzNjE1NzYsInN1YiI6ImpvaG5uIn0.zzBwOPVYABcLXcJnVi5LXF7sPOAWzikCVY3e3PP-KH0")*/

	//DEACTIVATE CLIENT
	/*
		database := db.GetDatabaseInstance()
		clientRepo := repository_impl.NewClientRepository(database)
		ctx := context.Background()
		authenticateUserUseCase := use_cases.NewClientUseCases(clientRepo)
		user, err := authenticateUserUseCase.DeactivateClient(ctx, 2)*/

	//UPDATE CLIENT

	username := entities.Client{
		ID:   2,
		Name: "John Doe", Phone: "TEST",
		Address: "TEST", Email: "johnn",
		Description: "DESC*", Activated: true}
	database := db.GetDatabaseInstance()
	userRepository := repository_impl.NewClientRepository(database)
	ctx := context.Background()
	authenticateUserUseCase := use_cases.NewClientUseCases(userRepository)
	err := authenticateUserUseCase.UpdateClient(ctx, &username)

	fmt.Printf("RESULT === %s .... \n", err)
}
