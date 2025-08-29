package ent

// activity schema
//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./schema/activities --target ../../internal/infrastructure/component/activity/persistence/ent/generated
//go:generate atlas migrate diff activity --dir "file://../migrations/activity" --to "ent://schema/activities" --dev-url "docker://postgres/17/test?search_path=public" --auto-approve

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./schema/attendances --target ../../internal/infrastructure/component/attendance/persistence/ent/generated
//go:generate atlas migrate diff attendance --dir "file://../migrations/attendance" --to "ent://schema/attendances" --dev-url "docker://postgres/17/test?search_path=public" --auto-approve

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./schema/users --target ../../internal/infrastructure/component/user/persistence/ent/generated
//go:generate atlas migrate diff user --dir "file://../migrations/user" --to "ent://schema/users" --dev-url "docker://postgres/17/test?search_path=public" --auto-approve
