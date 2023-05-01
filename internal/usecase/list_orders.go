package usecase

import (
	"github.com/ffsales/20-CleanArch/internal/entity"
	"github.com/ffsales/20-CleanArch/pkg/events"
)

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderListed     events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrderListed events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
		OrderListed:     OrderListed,
		EventDispatcher: EventDispatcher,
	}
}

func (l *ListOrdersUseCase) Execute() ([]OrderOutputDTO, error) {

	orders, err := l.OrderRepository.ListOrders()
	if err != nil {
		return []OrderOutputDTO{}, err
	}

	var listDtos []OrderOutputDTO

	for _, order := range orders {
		dto := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Price,
			FinalPrice: order.FinalPrice,
		}
		listDtos = append(listDtos, dto)
	}
	return listDtos, nil
}
