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

type CategoryBody struct {
	ID          uint
	Name        string
	Description string
	Activated   bool
	Company_id  int
}

type UnitBody struct {
	ID          uint
	Name        string
	Description string
	Activated   bool
	Company_id  int
}

type CreateProductBody struct {
	ID             uint
	Name           string
	Clave          string
	Unit           string
	Category       string
	Costo          float64
	Precio         float64
	Iva_percentage int
	Note           string
	Activated      bool
	Company_id     int
	Category_id    CategoryBody
	Unit_id        UnitBody
}

func updateProductHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Parsear el cuerpo de la solicitud para obtener los datos del cliente
	fmt.Printf("Procesando solicitud: %s\n", request)
	var body CreateProductBody
	var product entities.Product

	if err := json.Unmarshal([]byte(request.Body), &body); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
			}}, err
	}

	product = entities.Product{
		ID:             body.ID,
		Name:           body.Name,
		Clave:          body.Clave,
		Unit:           body.Unit,
		Category:       body.Category,
		Costo:          body.Costo,
		Precio:         body.Precio,
		Iva_percentage: body.Iva_percentage,
		Note:           body.Note,
		Activated:      body.Activated,
		Company_id:     body.Company_id,
		Id_category:    int(body.Category_id.ID),
		Unit_id:        int(body.Unit_id.ID),
	}

	// Inicializar repositorio y caso de uso
	database := db.GetDatabaseInstance()
	productRepo := repository_impl.NewProductRepository(database)
	createProductUseCase := use_cases.NewProductUseCase(productRepo)

	// update proveedor
	if err := createProductUseCase.UpdateProduct(ctx, &product); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
	}

	response := map[string]interface{}{
		"updated": true,
		"message": "Producto Actualizado Correctamente",
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
	lambda.Start(updateProductHandler)
}
