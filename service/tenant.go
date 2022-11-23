package service

import "github.com/thebluefowl/zion/model"

type TenantService struct {
	tenantRepository model.TenantRepository
}

func NewTenantService(tenantRepository model.TenantRepository) *TenantService {
	return &TenantService{tenantRepository: tenantRepository}
}

func (s *TenantService) Create(t *model.Tenant) (*model.Tenant, error) {
	if err := s.tenantRepository.Create(t); err != nil {
		return nil, err
	}
	return t, nil
}

func (s *TenantService) Get(id string) (*model.Tenant, error) {
	return s.tenantRepository.Get(id)
}
