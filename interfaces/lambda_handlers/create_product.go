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
	Type_id        uint
}

func createProductHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Procesando solicitud: %s\n", request)

	var body CreateProductBody
	var product entities.Product
	var category entities.Category
	var unit entities.MeasurementUnit
	var typeP entities.ProductType

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
		Type_id:        int(body.Type_id),
	}

	category = entities.Category{
		ID:          body.Category_id.ID,
		Name:        body.Category_id.Name,
		Description: body.Category_id.Description,
		Activated:   body.Category_id.Activated,
		Company_id:  body.Category_id.Company_id,
	}

	unit = entities.MeasurementUnit{
		ID:          body.Unit_id.ID,
		Name:        body.Unit_id.Name,
		Description: body.Unit_id.Description,
		Activated:   body.Unit_id.Activated,
		Company_id:  body.Unit_id.Company_id,
	}

	typeP = entities.ProductType{
		ID:          body.Type_id,
		Name:        "",
		Description: "",
		Activated:   true,
		Company_id:  body.Company_id,
	}

	// Inicializar repositorio y caso de uso
	database := db.GetDatabaseInstance()
	productRepo := repository_impl.NewProductRepository(database)
	categoryRepo := repository_impl.NewCategoryRepository(database)
	unitRepo := repository_impl.NewUnitRepository(database)
	typeRepo := repository_impl.NewTypeProductRepository(database)
	createProductUseCase := use_cases.NewProductUseCase(productRepo, use_cases.WithCategoryAndUnitRepoAndTypeRepo(categoryRepo, unitRepo, typeRepo))

	// Crear el Producto
	if err := createProductUseCase.RegisterProductIntoDB(ctx, &product, &category, &unit, &typeP); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
	}

	response := map[string]interface{}{
		"created": true,
		"message": "Producto Creado Correctamente",
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
	lambda.Start(createProductHandler)
}
