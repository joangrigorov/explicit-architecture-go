package mock

// 3rd party library mocks
//go:generate mockgen -destination=ext/go-playground/universal-translator/mock_translator.go -package=ut github.com/go-playground/universal-translator Translator
//go:generate mockgen -destination=ext/go-playground/validator/mock_errors.go -package=validator github.com/go-playground/validator/v10 FieldError
//go:generate mockgen -destination=ext/go-playground/validator/mock_field_level.go -package=validator github.com/go-playground/validator/v10 FieldLevel

// Repository mocks
//go:generate mockgen -source=../internal/core/component/activity/application/repositories/activity_repository.go -destination=core/component/activity/application/repositories/mock_activity_repository.go -package=repositories
//go:generate mockgen -source=../internal/core/component/attendance/application/repositories/attendance_repository.go -destination=core/component/attendance/application/repositories/mock_attendance_repository.go -package=repositories

//go:generate mockgen -source=../internal/presentation/web/infrastructure/framework/validation/rules.go -destination=presentation/web/infrastructure/framework/validation/mock_rules.go -package=validation
//go:generate mockgen -source=../internal/presentation/web/port/http/router.go -destination=presentation/web/port/http/mock_router.go -package=http
//go:generate mockgen -source=../internal/presentation/web/port/http/context.go -destination=presentation/web/port/http/mock_context.go -package=http
