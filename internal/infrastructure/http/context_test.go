package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGinContext_NoContent(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Setup router and handler
	r := gin.New()
	r.GET("/test", func(c *gin.Context) {
		gctx := NewGinContext(c)
		gctx.NoContent()
	})

	// Perform request
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestGinContext_Context(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a known value to inject into the context
	key := "user"
	val := "alice"

	r := gin.New()
	r.GET("/test", func(c *gin.Context) {
		// Replace context with one that has our value
		ctxWithVal := context.WithValue(c.Request.Context(), key, val)
		c.Request = c.Request.WithContext(ctxWithVal)

		gctx := NewGinContext(c)
		got := gctx.Context().Value(key)

		assert.Equal(t, val, got)
	})

	// Trigger the handler
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	r.ServeHTTP(w, req)
}

func TestGinContext_JSON(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.GET("/test", func(c *gin.Context) {
		gctx := NewGinContext(c)
		gctx.JSON(http.StatusOK, "ok")
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "\"ok\"", w.Body.String())
}

func TestGinContext_ShouldBindJSON(t *testing.T) {
	type example struct {
		Name string `json:"name"`
	}

	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.GET("/test", func(c *gin.Context) {
		gctx := NewGinContext(c)
		obj := &example{}
		if err := gctx.ShouldBindJSON(obj); err != nil {
			panic(err)
		}

		assert.Equal(t, "John Doe", obj.Name)
	})

	jsonBytes, err := json.Marshal(gin.H{"name": "John Doe"})
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test", bytes.NewReader(jsonBytes))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)
}

func TestGinContext_ParamString(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/test/:id", func(c *gin.Context) {
		gctx := NewGinContext(c)

		assert.Equal(t, "123abc", gctx.ParamString("id"))
	})
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test/123abc", nil)
	r.ServeHTTP(w, req)
}

func TestGinContext_ParamInt(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		r := gin.New()
		r.GET("/test/:id", func(c *gin.Context) {
			gctx := NewGinContext(c)
			paramInt, _ := gctx.ParamInt("id")
			assert.Equal(t, 123, paramInt)
		})
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/test/123", nil)
		r.ServeHTTP(w, req)
	})

	t.Run("cast failure", func(t *testing.T) {
		r := gin.New()
		r.GET("/test/:id", func(c *gin.Context) {
			gctx := NewGinContext(c)
			_, err := gctx.ParamInt("id")
			assert.Equal(t, "param error: id: abc cannot be converted to integer", err.Error())
		})
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/test/abc", nil)
		r.ServeHTTP(w, req)
	})
}

func TestGinContext_IsMethod(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("is GET", func(t *testing.T) {
		r := gin.New()
		r.GET("/test", func(c *gin.Context) {
			gctx := NewGinContext(c)
			assert.True(t, gctx.IsGet())
			assert.False(t, gctx.IsPost())
			assert.False(t, gctx.IsPatch())
			assert.False(t, gctx.IsPut())
			assert.False(t, gctx.IsDelete())
		})
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		r.ServeHTTP(w, req)
	})

	t.Run("is POST", func(t *testing.T) {
		r := gin.New()
		r.POST("/test", func(c *gin.Context) {
			gctx := NewGinContext(c)
			assert.True(t, gctx.IsPost())
			assert.False(t, gctx.IsGet())
			assert.False(t, gctx.IsPatch())
			assert.False(t, gctx.IsPut())
			assert.False(t, gctx.IsDelete())
		})
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/test", nil)
		r.ServeHTTP(w, req)
	})

	t.Run("is PATCH", func(t *testing.T) {
		r := gin.New()
		r.PATCH("/test", func(c *gin.Context) {
			gctx := NewGinContext(c)
			assert.True(t, gctx.IsPatch())
			assert.False(t, gctx.IsPost())
			assert.False(t, gctx.IsGet())
			assert.False(t, gctx.IsPut())
			assert.False(t, gctx.IsDelete())
		})
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPatch, "/test", nil)
		r.ServeHTTP(w, req)
	})

	t.Run("is PUT", func(t *testing.T) {
		r := gin.New()
		r.PUT("/test", func(c *gin.Context) {
			gctx := NewGinContext(c)
			assert.True(t, gctx.IsPut())
			assert.False(t, gctx.IsPatch())
			assert.False(t, gctx.IsPost())
			assert.False(t, gctx.IsGet())
			assert.False(t, gctx.IsDelete())
		})
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPut, "/test", nil)
		r.ServeHTTP(w, req)
	})

	t.Run("is DELETE", func(t *testing.T) {
		r := gin.New()
		r.DELETE("/test", func(c *gin.Context) {
			gctx := NewGinContext(c)
			assert.True(t, gctx.IsDelete())
			assert.False(t, gctx.IsPut())
			assert.False(t, gctx.IsPatch())
			assert.False(t, gctx.IsPost())
			assert.False(t, gctx.IsGet())
		})
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodDelete, "/test", nil)
		r.ServeHTTP(w, req)
	})
}

func TestGinContext_IsJsonRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("it is json request", func(t *testing.T) {
		r := gin.New()
		r.POST("/test", func(c *gin.Context) {
			gctx := NewGinContext(c)
			assert.True(t, gctx.IsJsonRequest())
		})
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/test", nil)
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)
	})

	t.Run("it is not a json request", func(t *testing.T) {
		r := gin.New()
		r.POST("/test", func(c *gin.Context) {
			gctx := NewGinContext(c)
			assert.False(t, gctx.IsJsonRequest())
		})
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/test", nil)
		r.ServeHTTP(w, req)
	})
}

// brokenReader I define my own reader that will trigger the json validation error on wrong byte read
type brokenReader struct{}

func (brokenReader) Read(p []byte) (int, error) {
	return 0, fmt.Errorf("read failure")
}

func (brokenReader) Close() error {
	return nil
}

func TestGinContext_IsJsonBodyValid(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("valid json body", func(t *testing.T) {
		r := gin.New()
		r.POST("/test", func(c *gin.Context) {
			gctx := NewGinContext(c)
			assert.True(t, gctx.IsJsonBodyValid())
		})

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/test", strings.NewReader(`{"this": "valid"}`))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)
	})

	t.Run("missing body is not a valid json body", func(t *testing.T) {
		r := gin.New()
		r.POST("/test", func(c *gin.Context) {
			gctx := NewGinContext(c)
			assert.False(t, gctx.IsJsonBodyValid())
		})

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/test", nil)
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)
	})

	t.Run("malformed json body", func(t *testing.T) {
		r := gin.New()
		r.POST("/test", func(c *gin.Context) {
			gctx := NewGinContext(c)
			assert.False(t, gctx.IsJsonBodyValid())
		})

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/test", strings.NewReader(`{"this": "is invalid}`))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)
	})

	t.Run("missing json content type", func(t *testing.T) {
		r := gin.New()
		r.POST("/test", func(c *gin.Context) {
			gctx := NewGinContext(c)
			assert.False(t, gctx.IsJsonBodyValid())
		})

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/test", strings.NewReader(`{"this": "is valid"}`))
		r.ServeHTTP(w, req)
	})

	t.Run("byte reader issue", func(t *testing.T) {
		r := gin.New()
		r.POST("/test", func(c *gin.Context) {
			gctx := NewGinContext(c)
			assert.False(t, gctx.IsJsonBodyValid())
		})

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/test", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Body = brokenReader{}
		r.ServeHTTP(w, req)
	})
}

func TestGinContext_Next(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/test", func(c *gin.Context) {
		NewGinContext(c).Next()
	}, func(c *gin.Context) {
		c.Status(http.StatusAccepted)
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/test", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusAccepted, w.Code)
}
