package daos

import (
	"github.com/kebabmane/tureloGo/app"
	"github.com/kebabmane/tureloGo/models"
)

// ArtistDAO persists artist data in database
type FeedEntryDAO struct{}

// NewArtistDAO creates a new ArtistDAO
func NewFeedEntryDAO() *FeedEntryDAO {
	return &FeedEntryDAO{}
}

// Get reads the artist with the specified ID from the database.
func (dao *FeedEntryDAO) Get(rs app.RequestScope, id int) (*models.FeedEntry, error) {
	var feedEntry models.FeedEntry
	err := rs.Tx().Select().Model(id, &feedEntry)
	return &feedEntry, err
}

// Create saves a new artist record in the database.
// The Artist.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *FeedEntryDAO) Create(rs app.RequestScope, feedEntry *models.FeedEntry) error {
	feedEntry.ID = 0
	return rs.Tx().Model(feedEntry).Insert()
}

// Update saves the changes to an artist in the database.
func (dao *FeedEntryDAO) Update(rs app.RequestScope, id int, feedEntry *models.FeedEntry) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	feedEntry.ID = id
	return rs.Tx().Model(feedEntry).Exclude("Id").Update()
}

// Delete deletes an artist with the specified ID from the database.
func (dao *FeedEntryDAO) Delete(rs app.RequestScope, id int) error {
	feedEntry, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(feedEntry).Delete()
}

// Count returns the number of the artist records in the database.
func (dao *FeedEntryDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").From("feed_entry").Row(&count)
	return count, err
}

// Query retrieves the artist records with the specified offset and limit from the database.
func (dao *FeedEntryDAO) Query(rs app.RequestScope, offset, limit int) ([]models.FeedEntry, error) {
	feedEntries := []models.FeedEntry{}
	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&feedEntries)
	return feedEntries, err
}
