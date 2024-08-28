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

type CreateBillBody struct {
	ID             uint
	Date           string
	Folio          string
	Status         string
	Client         int
	TotalOperation float64
	Company_id     int
	Products       []ProductsBody
}

func createBillHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Procesando solicitud: %s\n", request)

	var body CreateBillBody
	var bill entities.Bill

	if err := json.Unmarshal([]byte(request.Body), &body); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
			}}, err
	}

	bill = entities.Bill{
		ID:         body.ID,
		Client_id:  body.Client,
		Folio:      body.Folio,
		Status:     body.Status,
		Date:       body.Date,
		Total:      body.TotalOperation,
		Company_id: body.Company_id,
	}

	// Preparar detalles de la venta
	details := make([]*entities.BillDetail, 0, len(body.Products))
	for _, prod := range body.Products {
		detail := &entities.BillDetail{
			ProductID:    prod.IdProduct,
			ProductName:  prod.ProductName,
			ProductClave: prod.ProductClave,
			Quantity:     prod.Amount,
			UnitaryPrice: prod.UnitPrice,
			Iva:          prod.Iva,
			TotalPrice:   prod.Total,
		}
		details = append(details, detail)
	}

	// Inicializar repositorio y caso de uso
	database := db.GetDatabaseInstance()
	billRepo := repository_impl.NewBillRepository(database)
	createSaleUseCase := use_cases.NewBillUseCase(billRepo)

	// Crear el Producto
	if err := createSaleUseCase.RegisterBill(ctx, &bill, details); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
	}

	response := map[string]interface{}{
		"created": true,
		"message": "Factura Creada Correctamente",
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
	lambda.Start(createBillHandler)
}
