package mock

// 3rd party library mocks
//go:generate mockgen -destination=ext/go-playground/universal-translator/mock_translator.go -package=ut github.com/go-playground/universal-translator Translator
//go:generate mockgen -destination=ext/go-playground/validator/mock_errors.go -package=validator github.com/go-playground/validator/v10 FieldError
//go:generate mockgen -destination=ext/go-playground/validator/mock_field_level.go -package=validator github.com/go-playground/validator/v10 FieldLevel

// Repository mocks
//go:generate mockgen -source=../internal/core/component/activity/application/repositories/activity_repository.go -destination=core/component/activity/application/repositories/mock_activity_repository.go -package=repositories
//go:generate mockgen -source=../internal/core/component/attendance/application/repositories/attendance_repository.go -destination=core/component/attendance/application/repositories/mock_attendance_repository.go -package=repositories
//go:generate mockgen -source=../internal/core/component/user/application/repositories/user_repository.go -destination=core/component/user/application/repositories/mock_user_repository.go -package=repositories

// Application port mocks
//go:generate mockgen -source=../internal/core/port/events/bus.go -destination=core/port/events/mock_bus.go -package=events
//go:generate mockgen -source=../internal/core/port/logging/logger.go -destination=core/port/logging/mock_logger.go -package=logging
//go:generate mockgen -source=../internal/core/port/uuid/generator.go -destination=core/port/uuid/mock_generator.go -package=uuid

// Shared kernel mocks
//go:generate mockgen -source=../internal/core/shared_kernel/events/event.go -destination=core/shared_kernel/events/mock_event.go -package=events

// Presentation port mocks
//go:generate mockgen -source=../internal/infrastructure/framework/http/router.go -destination=infrastructure/framework/http/mock_router.go -package=http
//go:generate mockgen -source=../internal/infrastructure/framework/http/context.go -destination=infrastructure/framework/http/mock_context.go -package=http

//go:generate mockgen -source=../internal/infrastructure/framework/validation/rules.go -destination=infrastructure/framework/validation/mock_rules.go -package=validation

// Infrastructure mocks
//go:generate mockgen -source=../internal/infrastructure/framework/events/subscriber.go -destination=infrastructure/framework/events/mock_subscriber.go -package=events
