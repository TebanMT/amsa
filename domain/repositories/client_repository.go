package repositories

import (
	"context"

	"github.com/TebanMT/amsa/domain/entities"
)

type ClientRepository interface {
	CreateClient(ctx context.Context, client *entities.Client) error
	GetAllClients(ctx context.Context, idCompany int) ([]entities.Client, error)
	Deactivate(ctx context.Context, idUser int) (bool, error)
	UpdateClient(ctx context.Context, client *entities.Client) error
	GetCountClients(ctx context.Context, idCompany int) (int64, error)
	GetAllClientsWithBills(ctx context.Context, idCompany int) ([]entities.Client, error)
}
