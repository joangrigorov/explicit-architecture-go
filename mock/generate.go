package mock

// 3rd party library mocks
//go:generate mockgen -destination=ext/go-playground/universal-translator/mock_translator.go -package=ut github.com/go-playground/universal-translator Translator
//go:generate mockgen -destination=ext/go-playground/validator/mock_errors.go -package=validator github.com/go-playground/validator/v10 FieldError
//go:generate mockgen -destination=ext/go-playground/validator/mock_field_level.go -package=validator github.com/go-playground/validator/v10 FieldLevel

// Repository mocks
//go:generate mockgen -source=../internal/core/component/blog/application/repositories/post_repository.go -destination=core/component/blog/application/repositories/mock_post_repository.go -package=repositories

//go:generate mockgen -source=../internal/infrastructure/framework/validation/rules.go -destination=infrastructure/framework/validation/mock_rules.go -package=validation
//go:generate mockgen -source=../internal/presentation/web/port/http/router.go -destination=presentation/web/port/http/mock_router.go -package=http
//go:generate mockgen -source=../internal/presentation/web/port/http/context.go -destination=presentation/web/port/http/mock_context.go -package=http
