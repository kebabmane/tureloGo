package services

import (
	"github.com/kebabmane/tureloGo/app"
	"github.com/kebabmane/tureloGo/models"
)

type feedEntryDAO interface {
	Get(rs app.RequestScope, id int) (*models.FeedEntry, error)
	Count(rs app.RequestScope) (int, error)
	Query(rs app.RequestScope, offset, limit int) ([]models.FeedEntry, error)
	Create(rs app.RequestScope, feed *models.FeedEntry) error
	Update(rs app.RequestScope, id int, feed *models.FeedEntry) error
	Delete(rs app.RequestScope, id int) error
}

type FeedEntryService struct {
	dao feedEntryDAO
}

func NewFeedEntryService(dao feedEntryDAO) *FeedEntryService {
	return &FeedEntryService{dao}
}

func (s *FeedEntryService) Get(rs app.RequestScope, id int) (*models.FeedEntry, error) {
	return s.dao.Get(rs, id)
}

func (s *FeedEntryService) Create(rs app.RequestScope, model *models.FeedEntry) (*models.FeedEntry, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.ID)
}

func (s *FeedEntryService) Update(rs app.RequestScope, id int, model *models.FeedEntry) (*models.FeedEntry, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, id)
}

func (s *FeedEntryService) Delete(rs app.RequestScope, id int) (*models.FeedEntry, error) {
	feedEntry, err := s.dao.Get(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return feedEntry, err
}

func (s *FeedEntryService) Count(rs app.RequestScope) (int, error) {
	return s.dao.Count(rs)
}

func (s *FeedEntryService) Query(rs app.RequestScope, offset, limit int) ([]models.FeedEntry, error) {
	return s.dao.Query(rs, offset, limit)
}
