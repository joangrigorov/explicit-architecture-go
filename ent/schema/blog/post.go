package blog

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Post struct {
	ent.Schema
}

func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("content"),
	}
}

func (Post) Edges() []ent.Edge {
	return nil
}
