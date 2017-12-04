package apis

import (
	"strconv"

	"github.com/go-ozzo/ozzo-routing"
	"github.com/kebabmane/tureloGo/app"
	"github.com/kebabmane/tureloGo/models"
)

type (
	// artistService specifies the interface for the artist service needed by artistResource.
	feedEntryService interface {
		Get(rs app.RequestScope, id int) (*models.FeedEntry, error)
		Query(rs app.RequestScope, offset, limit int) ([]models.FeedEntry, error)
		Count(rs app.RequestScope) (int, error)
		Create(rs app.RequestScope, model *models.FeedEntry) (*models.FeedEntry, error)
		Update(rs app.RequestScope, id int, model *models.FeedEntry) (*models.FeedEntry, error)
		Delete(rs app.RequestScope, id int) (*models.FeedEntry, error)
	}

	// artistResource defines the handlers for the CRUD APIs.
	feedEntryResource struct {
		service feedEntryService
	}
)

// ServeArtist sets up the routing of artist endpoints and the corresponding handlers.
func ServeFeedEntryResource(rg *routing.RouteGroup, service feedEntryService) {
	r := &feedEntryResource{service}
	rg.Get("/feed_entries/<id>", r.get)
	rg.Get("/feed_entries", r.query)
	rg.Post("/feed_entries", r.create)
	rg.Put("/feed_entries/<id>", r.update)
	rg.Delete("/feed_entries/<id>", r.delete)
}

func (r *feedEntryResource) get(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	response, err := r.service.Get(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *feedEntryResource) query(c *routing.Context) error {
	rs := app.GetRequestScope(c)
	count, err := r.service.Count(rs)
	if err != nil {
		return err
	}
	paginatedList := getPaginatedListFromRequest(c, count)
	items, err := r.service.Query(app.GetRequestScope(c), paginatedList.Offset(), paginatedList.Limit())
	if err != nil {
		return err
	}
	paginatedList.Items = items
	return c.Write(paginatedList)
}

func (r *feedEntryResource) create(c *routing.Context) error {
	var model models.FeedEntry
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *feedEntryResource) update(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	rs := app.GetRequestScope(c)

	model, err := r.service.Get(rs, id)
	if err != nil {
		return err
	}

	if err := c.Read(model); err != nil {
		return err
	}

	response, err := r.service.Update(rs, id, model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *feedEntryResource) delete(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	response, err := r.service.Delete(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return c.Write(response)
}
