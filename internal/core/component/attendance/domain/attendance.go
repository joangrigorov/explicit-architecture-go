package domain

import "time"

type AttendanceId string
type AttendeeId string
type ActivityId string

type Attendance struct {
	events []AttendanceEvent

	Id AttendanceId

	Attendee *Attendee
	Activity *Activity

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAttendance(
	id AttendanceId,
	attendee *Attendee,
	activity *Activity,
) *Attendance {
	return &Attendance{
		events: []AttendanceEvent{
			NewAttendanceCreated(),
		},

		Id:       id,
		Attendee: attendee,
		Activity: activity,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type Attendee struct {
	Id AttendeeId
}

type Activity struct {
	Id               ActivityId
	Slug             string
	Title            string
	PosterImageUrl   string
	ShortDescription string
	HappensAt        time.Time
}

func ReconstituteAttendance(
	id AttendanceId,
	attendeeId AttendeeId,
	activityId ActivityId,
	activitySlug string,
	activityTitle string,
	activityPosterImageUrl string,
	activityShortDescription string,
	activityHappensAt time.Time,
	createdAt time.Time,
	updatedAt time.Time,
) *Attendance {
	return &Attendance{
		events: []AttendanceEvent{},

		Id: id,

		Attendee: &Attendee{Id: attendeeId},
		Activity: &Activity{
			Id:               activityId,
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
