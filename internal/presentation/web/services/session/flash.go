package session

import (
	"app/internal/infrastructure/framework/support"
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
)

func init() {
	gob.Register(map[string]string{})
	gob.Register(map[string]interface{}{})
	gob.Register(Alert{})
}

type Flash struct {
	tr ut.Translator
}

type AlertKind string

func (k AlertKind) String() string {
	return string(k)
}

const (
	AlertError   AlertKind = "error"
	AlertWarning AlertKind = "warning"
	AlertInfo    AlertKind = "info"
	AlertSuccess AlertKind = "success"
)

type Alert struct {
	Kind AlertKind
	Msg  string
}

func NewFlash(tr ut.Translator) *Flash {
	return &Flash{tr: tr}
}

func (f *Flash) AddAlert(session *sessions.Session, container string, kind AlertKind, msg string) {
	session.AddFlash(Alert{Kind: kind, Msg: msg}, container+":alert")
}

func (f *Flash) AddFormErrors(session *sessions.Session, form interface{}, err error) {
	container := fmt.Sprintf("%T", form)
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		bag := map[string]string{}
		for _, e := range ve {
			bag[support.TagFieldName(e, form, "form")] = e.Translate(f.tr)
		}
		session.AddFlash(bag, container+":errors")
	}
}

func (f *Flash) AddFormValues(session *sessions.Session, form interface{}) {
	container := fmt.Sprintf("%T", form)
	values := map[string]interface{}{}
	rv := reflect.ValueOf(form)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	rt := rv.Type()

	for i := 0; i < rv.NumField(); i++ {
		field := rt.Field(i)
		value := rv.Field(i)

		if !value.CanInterface() {
			continue
		}

		name := field.Tag.Get("form")
		if name == "" || name == "-" {
			name = field.Name
		} else {
			if idx := strings.Index(name, ","); idx != -1 {
				name = name[:idx]
			}
		}

		values[name] = value.Interface()
	}

	session.AddFlash(values, container+":values")
}

func (f *Flash) GetAlerts(session *sessions.Session, container string) []Alert {
	flashes := session.Flashes(container + ":alert")
	if len(flashes) == 0 {
		return nil
	}

	var alerts []Alert

	for _, flash := range flashes {
		alerts = append(alerts, flash.(Alert))
	}

	return alerts
}

// GetFormErrors retrieves validation errors for the given form type from session flashes.
func (f *Flash) GetFormErrors(session *sessions.Session, form interface{}) map[string]string {
	container := fmt.Sprintf("%T", form)
	key := container + ":errors"

	flashes := session.Flashes(key)
	if len(flashes) == 0 {
		log.Println("No flashes")
		return nil
	}

	// last one wins if multiple
	if bag, ok := flashes[len(flashes)-1].(map[string]string); ok {
		return bag
	}

	return nil
}

// GetFormValues retrieves submitted form values for the given form type from session flashes.
func (f *Flash) GetFormValues(session *sessions.Session, form interface{}) map[string]interface{} {
	container := fmt.Sprintf("%T", form)
	key := container + ":values"

	flashes := session.Flashes(key)
	if len(flashes) == 0 {
		return nil
	}

	if bag, ok := flashes[len(flashes)-1].(map[string]interface{}); ok {
		return bag
	}
	return nil
}
