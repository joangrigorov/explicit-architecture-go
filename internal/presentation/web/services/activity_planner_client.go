package services

import (
	"app/config/web"
	"app/internal/presentation/web/pages/identity/forms"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
)

type ActivityPlannerClient struct {
	host string
}

func NewActivityPlannerClient(cfg *web.Config) *ActivityPlannerClient {
	return &ActivityPlannerClient{
		host: cfg.Api.Host,
	}
}

func (s ActivityPlannerClient) SignUp(req *forms.SignUp) error {
	jsonBody, err := json.Marshal(req)

	if err != nil {
		return err
	}

	log.Println(string(jsonBody), s.host)

	request, err := http.NewRequest(http.MethodPost, s.host+"/user/v1/registration", bytes.NewBuffer(jsonBody))

	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	if !strings.Contains(response.Header.Get("Content-Type"), "application/json") {
		return errors.New("unexpected content type: " + string(response.Header.Get("Content-Type")))
	}

	if response.StatusCode != http.StatusOK {
		return errors.New("unexpected status code: " + string(body))
	}

	return nil
}
