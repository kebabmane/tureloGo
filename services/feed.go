package dao

import (
	"github.com/kebabmane/tureloGo/app"
	"github.com/kebabmane/tureloGo/models"
)

type feedDAO interface {
	Get(rs app.RequestScope, id int) (*models.Feed, error)
	Count(rs app.RequestScope) (int, error)
	Query(rs app.RequestScope, offset, limit int) ([]models.Feed, error)
	Create(rs app.RequestScope, feed *models.Feed) error
	Update(rs app.RequestScope, id int, feed *models.Feed) error
	Delete(rs app.RequestScope, id int) error
}

type FeedService struct {
	dao feedDAO
}

func NewFeedService(dao feedDAO) *FeedService {
	return &AFeedService{dao}
}

func (s *FeedService) Get(rs app.RequestScope, id int) (*models.Feed, error) {
	return s.dao.Get(rs, id)
}

func (s *FeedService) Create(rs app.RequestScope, model *models.Feed) (*models.Feed, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.Id)
}

func (s *FeedService) Update(rs app.RequestScope, id int, model *models.Feed) (*models.Feed, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, id)
}

func (s *FeedService) Delete(rs app.RequestScope, id int) (*models.Feed, error) {
	feed, err := s.dao.Get(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return feed, err
}

func (s *FeedService) Count(rs app.RequestScope) (int, error) {
	return s.dao.Count(rs)
}

func (s *FeedService) Query(rs app.RequestScope, offset, limit int) ([]models.Feed, error) {
	return s.dao.Query(rs, offset, limit)
}
