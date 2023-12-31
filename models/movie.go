package models

import "time"

type (
	// Movie
	Movie struct {
		ID                  uint              `json:"id" gorm:"primary_key"`
		Title               string            `json:"title"`
		Year                int               `json:"year"`
		AgeRatingCategoryID uint              `json:"AgeRatingCategoryID"`
		CreatedAt           time.Time         `json:"created_at"`
		UpdateAt            time.Time         `json:"updated_at"`
		AgeRatingCategory   AgeRatingCategory `json:"-"`
	}
)
