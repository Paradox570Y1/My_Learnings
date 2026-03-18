package services

import (
	"context"

	"MyProject/internal/dto"
	"MyProject/internal/models"
	"MyProject/internal/repository"
)

type FacilityService interface {
	GetAll(ctx context.Context) ([]dto.FacilityResponse, error)
	GetByCode(ctx context.Context, code string) (*dto.FacilityResponse, error)
	Create(ctx context.Context, facility dto.CreateFacilityRequest) error
}

type facilityService struct {
	facilityRepo repository.FacilityRepo
}

func NewFacilityService(repo repository.FacilityRepo) FacilityService {
	return &facilityService{
		facilityRepo: repo,
	}
}

func (s *facilityService) GetAll(ctx context.Context) ([]dto.FacilityResponse, error) {
	facilities, err := s.facilityRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]dto.FacilityResponse, 0, len(facilities))

	for _,f := range facilities {
		res = append(res, dto.FacilityResponse{
			Code:    f.Code,
			Name:    f.Name,
			Address: f.Address,
		})
	}
	
	return res, nil
}
func (s *facilityService) GetByCode(ctx context.Context, code string) (*dto.FacilityResponse, error) {
	f, err := s.facilityRepo.GetByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	return &dto.FacilityResponse{
		Code:    f.Code,
		Name:    f.Name,
		Address: f.Address,
	}, nil
}
func (s *facilityService) Create(ctx context.Context, facility dto.CreateFacilityRequest) error {
	return s.facilityRepo.Create(ctx, models.Facility{
		Code:    facility.Code,
		Name:    facility.Name,
		Address: facility.Address,
	})
}