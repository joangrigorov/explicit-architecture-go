package domain

import (
	"time"
)

type ActivityID string

type Activity struct {
	ID               ActivityID
	Slug             string
	Title            string
	PosterImageUrl   string
	ShortDescription string
	FullDescription  string
	HappensAt        time.Time
	Attendants       int
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func NewActivity(
	id ActivityID,
	slug string,
	title string,
	posterImageUrl string,
	shortDescription string,
	fullDescription string,
	happensAt time.Time,
	attendants int,
) *Activity {
	return &Activity{
		ID:               id,
		Slug:             slug,
		Title:            title,
		PosterImageUrl:   posterImageUrl,
		ShortDescription: shortDescription,
		FullDescription:  fullDescription,
		HappensAt:        happensAt,
		Attendants:       attendants,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func ReconstituteActivity(
	id ActivityID,
	slug string,
	title string,
	posterImageUrl string,
	shortDescription string,
	fullDescription string,
	happensAt time.Time,
	attendants int,
	createdAt time.Time,
	updatedAt time.Time,
) *Activity {
	return &Activity{
		ID:               id,
		Slug:             slug,
		Title:            title,
		PosterImageUrl:   posterImageUrl,
		ShortDescription: shortDescription,
		FullDescription:  fullDescription,
		HappensAt:        happensAt,
		Attendants:       attendants,

		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
