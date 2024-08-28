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

type ProductsBody struct {
	IdProduct    uint
	ProductName  string
	ProductClave string
	Amount       int
	UnitPrice    float64
	Iva          float64
	Total        float64
}

type CreateSaleBody struct {
	ID             uint
	Date           string
	Client         int
	TotalOperation float64
	Company_id     int
	Products       []ProductsBody
}

func createSaleHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Procesando solicitud: %s\n", request)

	var body CreateSaleBody
	var sale entities.Sale

	if err := json.Unmarshal([]byte(request.Body), &body); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
			}}, err
	}

	sale = entities.Sale{
		ID:         body.ID,
		ClientID:   body.Client,
		Date:       body.Date,
		Total:      body.TotalOperation,
		Company_id: body.Company_id,
	}

	// Preparar detalles de la venta
	details := make([]*entities.SaleDetail, 0, len(body.Products))
	for _, prod := range body.Products {
		detail := &entities.SaleDetail{
			ProductID:    prod.IdProduct,
			ProductName:  prod.ProductName,
			ProductClave: prod.ProductClave,
			Amount:       prod.Amount,
			Precio:       prod.UnitPrice,
			Iva:          prod.Iva,
			Total:        prod.Total,
		}
		details = append(details, detail)
	}

	// Inicializar repositorio y caso de uso
	database := db.GetDatabaseInstance()
	saleRepo := repository_impl.NewSaleRepository(database)
	createSaleUseCase := use_cases.NewSaleUseCase(saleRepo)

	// Crear el Producto
	if err := createSaleUseCase.RegisterSale(ctx, &sale, details); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
	}

	response := map[string]interface{}{
		"created": true,
		"message": "Venta Creada Correctamente",
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
	lambda.Start(createSaleHandler)
}
