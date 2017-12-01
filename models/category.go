package models

import "github.com/go-ozzo/ozzo-validation"

// Category represents an category record.
type Category struct {
	ID                  int    `json:"id" db:"id"`
	CategoryName        string `json:"category_name" db:"category_name"`
	CategoryImageURL    string `json:"category_image_url" db:"category_image_url"`
	CategoryDescription string `json:"category_description" db:"category_description"`
	FeedsCount          string `json:"feeds_count" db:"feeds_count"`
}

// Validate validates the category fields.
func (m Category) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.CategoryName, validation.Required, validation.Length(0, 120)),
	)
}
