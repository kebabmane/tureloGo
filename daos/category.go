package daos

import (
	"github.com/kebabmane/tureloGo/app"
	"github.com/kebabmane/tureloGo/models"
)

// ArtistDAO persists artist data in database
type CategoryDAO struct{}

// NewArtistDAO creates a new ArtistDAO
func NewCategoryDAO() *CategoryDAO {
	return &CategoryDAO{}
}

// Get reads the artist with the specified ID from the database.
func (dao *CategoryDAO) Get(rs app.RequestScope, id int) (*models.Category, error) {
	var category models.Category
	err := rs.Tx().Select().Model(id, &category)
	return &category, err
}

// Create saves a new artist record in the database.
// The Artist.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *CategoryDAO) Create(rs app.RequestScope, category *models.Category) error {
	category.ID = 0
	return rs.Tx().Model(category).Insert()
}

// Update saves the changes to an artist in the database.
func (dao *CategoryDAO) Update(rs app.RequestScope, id int, category *models.Category) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	category.ID = id
	return rs.Tx().Model(category).Exclude("Id").Update()
}

// Delete deletes an artist with the specified ID from the database.
func (dao *CategoryDAO) Delete(rs app.RequestScope, id int) error {
	category, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(category).Delete()
}

// Count returns the number of the artist records in the database.
func (dao *CategoryDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").From("category").Row(&count)
	return count, err
}

// Query retrieves the artist records with the specified offset and limit from the database.
func (dao *CategoryDAO) Query(rs app.RequestScope, offset, limit int) ([]models.Category, error) {
	categories := []models.Category{}
	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&categories)
	return categories, err
}
