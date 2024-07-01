package usecases

import (
	"log"

	"github.com/paulozy/btg-challenge/order-ms/internal/entity"
	"github.com/paulozy/btg-challenge/order-ms/internal/infra/database"
	"github.com/paulozy/btg-challenge/order-ms/pkg"
)

type ListOrdersByClientCodeUseCase struct {
	OrderRepository database.OrderRepositoryInterface
}

type ListOrdersByClientCodeInput struct {
	ClientCode int `json:"clientCode"`
}

func NewListOrdersByClientCodeUseCase(repo database.OrderRepositoryInterface) *ListOrdersByClientCodeUseCase {
	return &ListOrdersByClientCodeUseCase{
		OrderRepository: repo,
	}
}

func (useCase *ListOrdersByClientCodeUseCase) Execute(data ListOrdersByClientCodeInput) ([]entity.Order, pkg.Error) {

	println("clientCode", data.ClientCode)

	orders, err := useCase.OrderRepository.GetByClientCode(data.ClientCode)
	if err != nil {
		log.Fatalf("error: %s", err)
		return nil, pkg.NewInternalServerError(err)
	}

	return orders, pkg.Error{}
}
