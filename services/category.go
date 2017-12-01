package services

import (
	"github.com/kebabmane/tureloGo/app"
	"github.com/kebabmane/tureloGo/models"
)

type categoryDAO interface {
	Get(rs app.RequestScope, id int) (*models.Category, error)
	Count(rs app.RequestScope) (int, error)
	Query(rs app.RequestScope, offset, limit int) ([]models.Category, error)
	Create(rs app.RequestScope, category *models.Category) error
	Update(rs app.RequestScope, id int, category *models.Category) error
	Delete(rs app.RequestScope, id int) error
}

type CategoryService struct {
	dao categoryDAO
}

func NewCategoryService(dao categoryDAO) *CategoryService {
	return &CategoryService{dao}
}

func (s *CategoryService) Get(rs app.RequestScope, id int) (*models.Category, error) {
	return s.dao.Get(rs, id)
}

func (s *CategoryService) Create(rs app.RequestScope, model *models.Category) (*models.Category, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.ID)
}

func (s *CategoryService) Update(rs app.RequestScope, id int, model *models.Category) (*models.Category, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, id)
}

func (s *CategoryService) Delete(rs app.RequestScope, id int) (*models.Category, error) {
	category, err := s.dao.Get(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return category, err
}

func (s *CategoryService) Count(rs app.RequestScope) (int, error) {
	return s.dao.Count(rs)
}

func (s *CategoryService) Query(rs app.RequestScope, offset, limit int) ([]models.Category, error) {
	return s.dao.Query(rs, offset, limit)
}
