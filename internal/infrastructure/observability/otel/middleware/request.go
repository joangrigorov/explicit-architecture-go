package middleware

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func RecordRequestData(c *gin.Context) {
	span := trace.SpanFromContext(c.Request.Context())
	if !span.SpanContext().IsValid() {
		log.Fatalln(errors.New("invalid span context"))
		c.Next()
		return
	}

	var attrs []attribute.KeyValue

	// Query params
	for k, vals := range c.Request.URL.Query() {
		if len(vals) == 1 {
			attrs = append(attrs, attribute.String("http.query."+k, vals[0]))
		} else {
			attrs = append(attrs, attribute.StringSlice("http.query."+k, vals))
		}
	}

	// Path params
	for _, p := range c.Params {
		attrs = append(attrs, attribute.String("http.path_param."+p.Key, p.Value))
	}

	if len(attrs) > 0 {
		span.AddEvent("http.request.params", trace.WithAttributes(attrs...))
	}

	c.Next()
}

func mapToAttributes(m map[string]string) []attribute.KeyValue {
	attrs := make([]attribute.KeyValue, 0, len(m))
	for k, v := range m {
		attrs = append(attrs, attribute.String(k, v))
	}
	return attrs
}
