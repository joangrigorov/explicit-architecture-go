package users

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type Verification struct {
	ent.Schema
}

func (Verification) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}).Unique(),
		field.Text("csrf_token"),
		field.Time("expires_at"),
		field.Time("used_at").Nillable().Optional(),

		field.Time("created_at"),
	}
}

func (Verification) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("expires_at"),
		index.Fields("used_at"),
		index.Fields("created_at"),
	}
}

func (Verification) Config() ent.Config {
	return ent.Config{Table: "verifications"}
}
