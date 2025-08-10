package domain

import (
	"testing"
	"time"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
)

var f = faker.New()

func TestNewAttendance(t *testing.T) {
	id := AttendanceId(f.UUID().V4())

	attendeeId := AttendeeId(f.UUID().V4())

	activityId := ActivityId(f.UUID().V4())
	activitySlug := f.Lorem().Word() + "-" + f.Lorem().Word()
	activityTitle := f.Lorem().Sentence(3)
	activityPosterImageUrl := f.ProfileImage().Image().Name()
	activityShortDescription := f.Lorem().Sentence(10)
	activityHappensAt := f.Time().Time(time.Now())

	attendance := NewAttendance(
		id,
		&Attendee{Id: attendeeId},
		&Activity{
			Id:               activityId,
			Slug:             activitySlug,
			Title:            activityTitle,
			PosterImageUrl:   activityPosterImageUrl,
			ShortDescription: activityShortDescription,
			HappensAt:        activityHappensAt,
		},
	)

	assert.NotNil(t, attendance)

	assert.Equal(t, id, attendance.Id)

	assert.Equal(t, attendeeId, attendance.Attendee.Id)

	assert.Equal(t, activityId, attendance.Activity.Id)
	assert.Equal(t, activitySlug, attendance.Activity.Slug)
	assert.Equal(t, activityPosterImageUrl, attendance.Activity.PosterImageUrl)
	assert.Equal(t, activityTitle, attendance.Activity.Title)
	assert.Equal(t, activityShortDescription, attendance.Activity.ShortDescription)
	assert.Equal(t, activityHappensAt, attendance.Activity.HappensAt)
	assert.NotNil(t, attendance.CreatedAt)
	assert.NotNil(t, attendance.UpdatedAt)

	assert.Len(t, attendance.events, 1)
	assert.IsType(t, &AttendanceCreated{}, attendance.events[0])
}

func TestReconstituteAttendance(t *testing.T) {
	id := AttendanceId(f.UUID().V4())

	attendeeId := AttendeeId(f.UUID().V4())

	activityId := ActivityId(f.UUID().V4())
	activitySlug := f.Lorem().Word() + "-" + f.Lorem().Word()
	activityTitle := f.Lorem().Sentence(3)
	activityPosterImageUrl := f.ProfileImage().Image().Name()
	activityShortDescription := f.Lorem().Sentence(10)
	activityHappensAt := f.Time().Time(time.Now())

	createdAt := f.Time().Time(time.Now())
	updatedAt := f.Time().Time(time.Now())

	attendance := ReconstituteAttendance(
		id,
		attendeeId,
		activityId,
		activitySlug,
		activityTitle,
		activityPosterImageUrl,
		activityShortDescription,
		activityHappensAt,
		createdAt,
		updatedAt,
	)

	assert.NotNil(t, attendance)

	assert.Equal(t, id, attendance.Id)
	assert.Equal(t, attendeeId, attendance.Attendee.Id)
	assert.Equal(t, activityId, attendance.Activity.Id)
	assert.Equal(t, activitySlug, attendance.Activity.Slug)
	assert.Equal(t, activityTitle, attendance.Activity.Title)
	assert.Equal(t, activityPosterImageUrl, attendance.Activity.PosterImageUrl)
	assert.Equal(t, activityShortDescription, attendance.Activity.ShortDescription)
	assert.Equal(t, activityHappensAt, attendance.Activity.HappensAt)

	assert.Len(t, attendance.events, 0)
}
