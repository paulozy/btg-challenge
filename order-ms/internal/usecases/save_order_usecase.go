package usecases

import (
	"log"

	"github.com/paulozy/btg-challenge/order-ms/internal/entity"
	"github.com/paulozy/btg-challenge/order-ms/internal/infra/database"
	"github.com/paulozy/btg-challenge/order-ms/pkg"
)

type SaveOrderUseCase struct {
	OrderRepository database.OrderRepositoryInterface
}

type SaveOrderInput struct {
	OrderCode  int           `json:"orderCode"`
	ClientCode int           `json:"clientCode"`
	Items      []entity.Item `json:"items"`
}

func NewSaveOrderUseCase(repo database.OrderRepositoryInterface) *SaveOrderUseCase {
	return &SaveOrderUseCase{
		OrderRepository: repo,
	}
}

func (useCase *SaveOrderUseCase) Execute(data SaveOrderInput) (*entity.Order, pkg.Error) {
	order := entity.NewOrder(
		data.OrderCode,
		data.ClientCode,
		data.Items,
	)

	err := useCase.OrderRepository.Save(order)
	if err != nil {
		log.Fatalf("error: %s", err)
		return nil, pkg.NewInternalServerError(err)
	}

	return order, pkg.Error{}
}
