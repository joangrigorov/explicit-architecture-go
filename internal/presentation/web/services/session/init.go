package session

import (
	"app/config/web"

	"github.com/boj/redistore"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"go.uber.org/zap"
)

func NewSessionStore(cfg web.Config) sessions.Store {
	sessionCfg := cfg.Session
	sessionSecret := []byte(sessionCfg.SessionSecret)

	var (
		store sessions.Store
		err   error
	)

	switch sessionCfg.Driver {
	case "redis":
		store, err = redistore.NewRediStore(
			sessionCfg.RedisPoolSize,
			sessionCfg.RedisNetwork,
			sessionCfg.RedisAddr,
			sessionCfg.RedisUsername,
			sessionCfg.RedisPassword,
			sessionSecret,
		)
		store.(*redistore.RediStore).SetMaxLength(sessionCfg.RedisMaxLengthKB * 1024)
	case "file":
		store = sessions.NewFilesystemStore(sessionCfg.FilesystemPath, sessionSecret)
	case "cookie":
	default:
		store = sessions.NewCookieStore(sessionSecret)
	}

	if err != nil {
		panic(err)
	}

	return store
}

func Handler(logger *zap.SugaredLogger, store sessions.Store, sessionName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := store.Get(c.Request, sessionName)
		if err != nil {
			session, _ = store.New(c.Request, sessionName)
		}

		c.Set("session", session)

		c.Next()

		if err := sessions.Save(c.Request, c.Writer); err != nil {
			logger.Error("failed to save session", zap.Error(err))
		}
	}
}

func GetSession(c *gin.Context) *sessions.Session {
	return c.MustGet("session").(*sessions.Session)
}
