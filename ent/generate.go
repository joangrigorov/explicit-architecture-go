package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./schema/activity --target ../internal/infrastructure/persistence/ent/generated/activity
//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./schema/attendance --target ../internal/infrastructure/persistence/ent/generated/attendance
