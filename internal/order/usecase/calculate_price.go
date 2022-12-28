package usecase

import (
	"github.com/moisesPompilio/calculadora_taxa/internal/order/entity"
	"github.com/moisesPompilio/calculadora_taxa/internal/order/infra/database"
)

type OrderInputDTO struct {
	Id    string
	Price float64
	Tax   float64
}

type OrderOutputDTO struct {
	Id         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

type CalculateFinalPriceUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewCalculateFinalPriceUseCase(orderRepository database.OrderRepository) *CalculateFinalPriceUseCase {
	return &CalculateFinalPriceUseCase{
		OrderRepository: &orderRepository,
	}
}

func (uc *CalculateFinalPriceUseCase) Execute(input OrderInputDTO) (*OrderOutputDTO, error) {
	order, err := entity.NewOrder(input.Id, input.Price, input.Tax)
	if err != nil {
		return nil, err
	}

	err = order.CalculateFinalPrice()
	if err != nil {
		return nil, err
	}

	err = uc.OrderRepository.Save(order)
	if err != nil {
		return nil, err
	}
	return &OrderOutputDTO{
		Id:         order.Id,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
