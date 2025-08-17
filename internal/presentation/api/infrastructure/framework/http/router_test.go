package http

import (
	port "app/internal/presentation/api/port/http"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TODO rewrite all of the tests below to use mocked gin interfaces instead of httptest.
// TODO httptest seems better suited for integration tests

func TestRouter_Group(t *testing.T) {
	gin.SetMode(gin.TestMode)

	engine := gin.New()
	r := NewRouter(engine)
	group := r.Group("/test")

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	engine.ServeHTTP(w, req)

	assert.NotNil(t, group)
	assert.Implements(t, (*port.Router)(nil), group)
}

func TestRouter_Use(t *testing.T) {
	gin.SetMode(gin.TestMode)

	engine := gin.New()
	r := NewRouter(engine)
	r.Use(func(c port.Context) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{})
	})
	r.GET("/test", func(c port.Context) {
		c.NoContent()
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestRouter_POST(t *testing.T) {
	gin.SetMode(gin.TestMode)

	engine := gin.New()
	r := NewRouter(engine)

	r.POST("/test", func(c port.Context) {
		c.NoContent()
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/test", nil)
	engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestRouter_GET(t *testing.T) {
	gin.SetMode(gin.TestMode)

	engine := gin.New()
	r := NewRouter(engine)

	r.GET("/test", func(c port.Context) {
		c.NoContent()
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestRouter_PATCH(t *testing.T) {
	gin.SetMode(gin.TestMode)

	engine := gin.New()
	r := NewRouter(engine)

	r.PATCH("/test", func(c port.Context) {
		c.NoContent()
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPatch, "/test", nil)
	engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestRouter_PUT(t *testing.T) {
	gin.SetMode(gin.TestMode)

	engine := gin.New()
	r := NewRouter(engine)

	r.PUT("/test", func(c port.Context) {
		c.NoContent()
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/test", nil)
	engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestRouter_DELETE(t *testing.T) {
	gin.SetMode(gin.TestMode)

	engine := gin.New()
	r := NewRouter(engine)

	r.DELETE("/test", func(c port.Context) {
		c.NoContent()
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/test", nil)
	engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestRouter_NewGinEngine(t *testing.T) {
	gin.SetMode(gin.TestMode)
	engine := NewGinEngine()
	assert.NotNil(t, engine)
}
