package services

import (
	"context"
	"time"
	"errors"

	"MyProject/internal/dto"
	"MyProject/internal/models"
	"MyProject/internal/repository"
	"MyProject/internal/constant"
)

var (
	ErrOrderNotFound = errors.New("order not found")
	ErrFacilityNotFound = errors.New("facility not found")
)

func IsOrderNotFound(err error) bool {
	return errors.Is(err, ErrOrderNotFound)
}

func IsFacilityNotFound(err error) bool {
	return errors.Is(err, ErrFacilityNotFound)
}

type OrderService interface {
	GetAll(ctx context.Context) ([]dto.OrderResponse, error)
	GetByID(ctx context.Context, id string) (*dto.OrderResponse, error)
	Create(ctx context.Context, order dto.CreateOrderRequest) error
}

type orderService struct {
	orderRepo    repository.OrderRepo
	facilityRepo repository.FacilityRepo
}

func NewOrderService(orderRepo repository.OrderRepo, facilityRepo repository.FacilityRepo) OrderService {
	return &orderService{
		orderRepo: orderRepo,
		facilityRepo: facilityRepo,
	}
}

func (s *orderService) GetAll(ctx context.Context) ([]dto.OrderResponse, error) {
	res, err := s.orderRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	
	orders := make([]dto.OrderResponse, 0, len(res))
	
	for _, o := range res {
		orders = append(orders, dto.OrderResponse{
			ID: o.ID,
			FacilityCode: o.FacilityCode,
			Status: o.Status,
			CreatedAt: o.CreatedAt,
		})
	}

	return orders, nil
}
func (s *orderService) GetByID(ctx context.Context, id string) (*dto.OrderResponse, error) {
	o, err := s.orderRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if o == nil {
		return nil, ErrOrderNotFound
	}
	return &dto.OrderResponse{
		ID: o.ID,
		FacilityCode: o.FacilityCode,
		Status: o.Status,
		CreatedAt: o.CreatedAt,
	}, nil
}
func (s *orderService) Create(ctx context.Context, order dto.CreateOrderRequest) error {
	if _, err := s.facilityRepo.GetByCode(ctx, order.FacilityCode); err != nil {
		return ErrFacilityNotFound
	}

	new_order := models.Order{
		ID: order.ID,
		FacilityCode: order.FacilityCode,
		Status: constant.CreatedOrderStatus,
		CreatedAt: time.Now().UTC(),
	}

	return s.orderRepo.Create(ctx, new_order)
}