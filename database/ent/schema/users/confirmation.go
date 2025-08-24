package users

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type Confirmation struct {
	ent.Schema
}

func (Confirmation) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}).Unique(),
		field.String("hmac_secret").Default("secret"),

		field.Time("created_at"),
	}
}

func (Confirmation) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at"),
	}
}

func (Confirmation) Config() ent.Config {
	return ent.Config{Table: "confirmations"}
}
