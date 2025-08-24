package users

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.String("email").Unique(),
		field.String("username").Unique(),
		field.String("first_name"),
		field.String("last_name"),
		field.Enum("role").Values("admin", "member"),
		field.Time("confirmed_at").Optional().Nillable(),
		field.String("idp_user_id").Unique().Optional().Nillable(),

		field.Time("created_at"),
		field.Time("updated_at"),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("role"),
		index.Fields("confirmed_at"),
		index.Fields("created_at"),
		index.Fields("updated_at"),
		index.Fields("deleted_at"),
	}
}

func (User) Config() ent.Config {
	return ent.Config{Table: "users"}
}
