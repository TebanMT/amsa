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
	IvaPrice     float64
	Total        float64
}

type CreatePurchaseBody struct {
	ID             uint
	Date           string
	Folio          string
	Status         string
	Supplier       int
	TotalOperation float64
	Company_id     int
	Products       []ProductsBody
}

func createPurchaseOrderHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Procesando solicitud: %s\n", request)

	var body CreatePurchaseBody
	var purchase entities.PurchaseOrder

	if err := json.Unmarshal([]byte(request.Body), &body); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
			}}, err
	}

	purchase = entities.PurchaseOrder{
		ID:          body.ID,
		Supplier_id: body.Supplier,
		Folio:       body.Folio,
		Status:      body.Status,
		OrderDate:   body.Date,
		Total:       body.TotalOperation,
		Company_id:  body.Company_id,
	}

	// Preparar detalles de la venta
	details := make([]*entities.PurchaseOrderDetail, 0, len(body.Products))
	for _, prod := range body.Products {
		detail := &entities.PurchaseOrderDetail{
			ProductID:    prod.IdProduct,
			ProductName:  prod.ProductName,
			ProductClave: prod.ProductClave,
			Quantity:     prod.Amount,
			UnitaryPrice: prod.UnitPrice,
			Iva:          prod.IvaPrice,
			TotalPrice:   prod.Total,
		}
		details = append(details, detail)
	}

	// Inicializar repositorio y caso de uso
	database := db.GetDatabaseInstance()
	purchaseRepo := repository_impl.NewPurchaseOrderRepository(database)
	purchaseOrderUseCase := use_cases.NewPurchaseOrderUseCase(purchaseRepo)

	// Crear el Producto
	if err := purchaseOrderUseCase.RegisterPurchaseOrder(ctx, &purchase, details); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
	}

	response := map[string]interface{}{
		"created": true,
		"message": "Orden de Compra Creada Correctamente",
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
	lambda.Start(createPurchaseOrderHandler)
}
