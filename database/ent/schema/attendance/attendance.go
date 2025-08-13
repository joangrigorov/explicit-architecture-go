package attendance

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type Attendance struct {
	ent.Schema
}

func (Attendance) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),

		field.UUID("attendee_id", uuid.UUID{}),
		field.UUID("activity_id", uuid.UUID{}),

		field.String("activity_slug"),
		field.String("activity_title"),
		field.String("activity_poster_image_url"),
		field.String("activity_short_description"),
		field.Time("activity_happens_at").Comment("Start time of the activity"),

		field.Time("created_at"),
		field.Time("updated_at"),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

func (Attendance) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("attendee_id", "activity_id").Unique(),
		index.Fields("activity_slug"),
		index.Fields("activity_happens_at"),
		index.Fields("created_at"),
		index.Fields("updated_at"),
		index.Fields("deleted_at"),
	}
}

func (Attendance) Config() ent.Config {
	return ent.Config{Table: "attendances"}
}
