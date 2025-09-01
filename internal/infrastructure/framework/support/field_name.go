package support

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// TagFieldName returns the JSON/Form tag-based field path for the given validator.FieldError
func TagFieldName(fe validator.FieldError, root interface{}, tag string) string {
	ns := fe.StructNamespace()
	parts := strings.Split(ns, ".")

	t := reflect.TypeOf(root)
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	rootName := t.Name()
	if len(parts) == 0 || parts[0] != rootName {
		panic(fmt.Sprintf("validation error namespace root %q does not match provided type %q", parts[0], rootName))
	}

	if len(parts) > 0 {
		parts = parts[1:]
	}

	var jsonPath []string

	for _, part := range parts {
		if t.Kind() != reflect.Struct {
			panic("nesting mismatch")
		}

		var field reflect.StructField
		found := false

		for i := 0; i < t.NumField(); i++ {
			field = t.Field(i)
			if field.Name == part {
				found = true
				break
			}
		}

		if !found {
			panic(fmt.Sprintf("field %q not found in struct %s", part, t.Name()))
		}

		jsonTag := field.Tag.Get(tag)
		jsonName := strings.Split(jsonTag, ",")[0]
		if jsonName == "-" || jsonName == "" {
			jsonName = field.Name
		}

		jsonPath = append(jsonPath, jsonName)
		t = field.Type
		for t.Kind() == reflect.Pointer {
			t = t.Elem()
		}
	}

	return strings.Join(jsonPath, ".")
}
