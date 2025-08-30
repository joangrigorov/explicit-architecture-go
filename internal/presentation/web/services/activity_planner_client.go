package services

import (
	"app/config/web"
	"app/internal/presentation/web/pages/identity/forms"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

func (s *ActivityPlannerClient) SignUp(req *forms.SignUp) error {
	jsonBody, err := json.Marshal(req)

	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s/user/v1/registration", s.host)
	request, err := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(jsonBody))

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

type VerificationPreflightResponse struct {
	Error       string `json:"error,omitempty"`
	ValidCSRF   bool   `json:"valid_csrf"`
	MaskedEmail string `json:"masked_email"`
	Expired     bool   `json:"expired"`
}

func (s *ActivityPlannerClient) PreflightVerification(verificationID, token string) (*VerificationPreflightResponse, error) {
	uri := fmt.Sprintf("%s/user/v1/verifications/%s/preflight?token=%s", s.host, verificationID, token)
	request, err := http.NewRequest(http.MethodGet, uri, nil)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if !strings.Contains(response.Header.Get("Content-Type"), "application/json") {
		return nil, errors.New("unexpected content type: " + string(response.Header.Get("Content-Type")))
	}

	var resp VerificationPreflightResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

type PasswordSetupResponse struct {
	Error string `json:"error,omitempty"`
}

func (s *ActivityPlannerClient) PasswordSetup(req *forms.PasswordSetup) error {
	jsonBody, err := json.Marshal(req)

	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s/user/v1/verifications/%s/password-setup", s.host, req.VerificationID)
	request, err := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(jsonBody))

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

	if response.StatusCode == http.StatusNoContent {
		return nil
	}

	if response.StatusCode == http.StatusGone {
		return errors.New("CSRF token expired")
	}

	if response.StatusCode == http.StatusBadRequest {
		return errors.New("CSRF token is invalid")
	}

	return errors.New("Backend responded with unexpected status code: " + string(rune(response.StatusCode)))
}
