package activities

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type Activity struct {
	ent.Schema
}

func (Activity) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.String("slug").Unique(),
		field.String("title"),
		field.String("poster_image_url"),
		field.String("short_description"),
		field.Text("full_description"),
		field.Time("happens_at").Comment("Start time of the activity"),

		field.Int("attendants").Comment("Number of activity attendants"),

		field.Time("created_at"),
		field.Time("updated_at"),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

func (Activity) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("happens_at"),
		index.Fields("attendants"),
		index.Fields("created_at"),
		index.Fields("updated_at"),
		index.Fields("deleted_at"),
	}
}

func (Activity) Config() ent.Config {
	return ent.Config{Table: "activities"}
}
