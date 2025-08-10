package domain

import (
	"time"
)

type ActivityId string

type Activity struct {
	events []ActivityEvent

	Id               ActivityId
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
	id ActivityId,
	slug string,
	title string,
	posterImageUrl string,
	shortDescription string,
	fullDescription string,
	happensAt time.Time,
	attendants int,
) *Activity {
	return &Activity{
		events: []ActivityEvent{
			NewActivityCreated(),
		},

		Id:               id,
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
	id ActivityId,
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
		events: []ActivityEvent{},

		Id:               id,
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
