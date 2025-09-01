package activity_planner

import (
	"app/config/web"
	"app/internal/presentation/web/pages/identity/forms"
	"app/internal/presentation/web/services/identity"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	host string
}

type Profile struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (s *Client) SignUp(req *forms.SignUp) error {
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

	if response.StatusCode == http.StatusCreated || response.StatusCode == http.StatusNoContent || response.StatusCode == http.StatusOK {
		return nil
	}

	if response.StatusCode == http.StatusUnprocessableEntity || response.StatusCode == http.StatusBadRequest || response.StatusCode == http.StatusConflict {
		var errResponse ErrorResponse
		if err = json.Unmarshal(body, &errResponse); err != nil {
			return err
		}
		return errors.New(errResponse.Error)
	}

	return errors.New("unexpected status code")
}

type VerificationPreflightResponse struct {
	Error       string `json:"error,omitempty"`
	ValidCSRF   bool   `json:"valid_csrf"`
	MaskedEmail string `json:"masked_email"`
	Expired     bool   `json:"expired"`
}

func (s *Client) PreflightVerification(verificationID, token string) (*VerificationPreflightResponse, error) {
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

func (s *Client) PasswordSetup(req *forms.PasswordSetup) error {
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

func (s *Client) Me(session identity.AuthenticationSession) (Profile, error) {
	var profile Profile

	uri := fmt.Sprintf("%s/user/v1/me", s.host)
	request, err := http.NewRequest(http.MethodGet, uri, nil)

	if err != nil {
		return profile, err
	}

	request.Header.Set("Authorization", "Bearer "+session.AccessToken)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return profile, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if !strings.Contains(response.Header.Get("Content-Type"), "application/json") {
		return profile, errors.New("unexpected content type: " + string(response.Header.Get("Content-Type")))
	}

	if err = json.Unmarshal(body, &profile); err != nil {
		return profile, err
	}

	return profile, nil
}

func (p Profile) Initials() string {
	if len(p.FirstName) == 0 || len(p.LastName) == 0 {
		return ""
	}

	// Take first rune of each, to support Unicode properly
	f := []rune(p.FirstName)[0:1]
	l := []rune(p.LastName)[0:1]

	return strings.ToUpper(string(f)) + strings.ToUpper(string(l))
}

func (p Profile) FullName() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

func NewClient(cfg web.Config) *Client {
	return &Client{
		host: cfg.Api.Host,
	}
}
