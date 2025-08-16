package domain

import "time"

type AttendanceID string
type AttendeeID string
type ActivityID string

type Attendance struct {
	ID AttendanceID

	Attendee *Attendee
	Activity *Activity

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAttendance(
	id AttendanceID,
	attendee *Attendee,
	activity *Activity,
) *Attendance {
	return &Attendance{
		ID:       id,
		Attendee: attendee,
		Activity: activity,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type Attendee struct {
	ID AttendeeID
}

type Activity struct {
	ID               ActivityID
	Slug             string
	Title            string
	PosterImageUrl   string
	ShortDescription string
	HappensAt        time.Time
}

func ReconstituteAttendance(
	id AttendanceID,
	attendeeId AttendeeID,
	activityId ActivityID,
	activitySlug string,
	activityTitle string,
	activityPosterImageUrl string,
	activityShortDescription string,
	activityHappensAt time.Time,
	createdAt time.Time,
	updatedAt time.Time,
) *Attendance {
	return &Attendance{
		ID: id,

		Attendee: &Attendee{ID: attendeeId},
		Activity: &Activity{
			ID:               activityId,
			Slug:             activitySlug,
			Title:            activityTitle,
			PosterImageUrl:   activityPosterImageUrl,
			ShortDescription: activityShortDescription,
			HappensAt:        activityHappensAt,
		},

		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
