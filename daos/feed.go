package daos

import (
	"github.com/kebabmane/tureloGo/app"
	"github.com/kebabmane/tureloGo/models"
)

// ArtistDAO persists artist data in database
type FeedDAO struct{}

// NewArtistDAO creates a new ArtistDAO
func NewFeedDAO() *FeedDAO {
	return &FeedDAO{}
}

// Get reads the artist with the specified ID from the database.
func (dao *FeedDAO) Get(rs app.RequestScope, id int) (*models.Feed, error) {
	var feed models.Feed
	err := rs.Tx().Select().Model(id, &feed)
	return &feed, err
}

// Create saves a new artist record in the database.
// The Artist.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *FeedDAO) Create(rs app.RequestScope, feed *models.Feed) error {
	feed.ID = 0
	return rs.Tx().Model(feed).Insert()
}

// Update saves the changes to an artist in the database.
func (dao *FeedDAO) Update(rs app.RequestScope, id int, feed *models.Feed) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	feed.ID = id
	return rs.Tx().Model(feed).Exclude("Id").Update()
}

// Delete deletes an artist with the specified ID from the database.
func (dao *FeedDAO) Delete(rs app.RequestScope, id int) error {
	feed, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(feed).Delete()
}

// Count returns the number of the artist records in the database.
func (dao *FeedDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").From("feed").Row(&count)
	return count, err
}

// Query retrieves the artist records with the specified offset and limit from the database.
func (dao *FeedDAO) Query(rs app.RequestScope, offset, limit int) ([]models.Feed, error) {
	feeds := []models.Feed{}
	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&feeds)
	return feeds, err
}
