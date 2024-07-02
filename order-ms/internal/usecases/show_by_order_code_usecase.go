package usecases

import (
	"github.com/paulozy/btg-challenge/order-ms/internal/entity"
	"github.com/paulozy/btg-challenge/order-ms/internal/infra/database"
	"github.com/paulozy/btg-challenge/order-ms/pkg"
)

type ShowOrderByOrderCodeUseCase struct {
	OrderRepository database.OrderRepositoryInterface
}

type ShowOrderByOrderCodeInput struct {
	OrderCode int `json:"orderCode"`
}

func NewShowOrderByOrderCodeUseCase(repo database.OrderRepositoryInterface) *ShowOrderByOrderCodeUseCase {
	return &ShowOrderByOrderCodeUseCase{
		OrderRepository: repo,
	}
}

func (useCase *ShowOrderByOrderCodeUseCase) Execute(data ShowOrderByOrderCodeInput) (*entity.Order, pkg.Error) {
	order, err := useCase.OrderRepository.FindByOrderCode(data.OrderCode)
	if err != nil {
		return nil, pkg.NewInternalServerError(err)
	}

	return order, pkg.Error{}
}
