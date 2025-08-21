package ent

// activity schema
//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./schema/activity --target ../../internal/infrastructure/components/activity/persistence/ent/generated
//go:generate atlas migrate diff activity --dir "file://../migrations/activity" --to "ent://schema/activity" --dev-url "docker://postgres/17/test?search_path=public"

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./schema/attendance --target ../../internal/infrastructure/components/attendance/persistence/ent/generated
//go:generate atlas migrate diff attendance --dir "file://../migrations/attendance" --to "ent://schema/attendance" --dev-url "docker://postgres/17/test?search_path=public"

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./schema/user --target ../../internal/infrastructure/components/user/persistence/ent/generated
//go:generate atlas migrate diff user --dir "file://../migrations/user" --to "ent://schema/user" --dev-url "docker://postgres/17/test?search_path=public"
