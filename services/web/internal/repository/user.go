package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type userRepository struct {
	baseURL string
}

func NewUserRepository(baseURL string) UserRepository {
	return &userRepository{
		baseURL: baseURL,
	}
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
	Error   string `json:"error"`
}

type registerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type registerResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func (r *userRepository) Login(email, password string) (string, error) {
	loginData := loginRequest{
		Email:    email,
		Password: password,
	}

	jsonData, err := json.Marshal(loginData)
	if err != nil {
		return "", fmt.Errorf("failed to marshal login data: %w", err)
	}

	resp, err := http.Post(
		r.baseURL+"/users/login",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return "", fmt.Errorf("failed to send login request: %w", err)
	}
	defer resp.Body.Close()

	var result loginResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("login failed: %s", result.Error)
	}

	return result.Token, nil
}

func (r *userRepository) Register(email, password, name string) error {
	registerData := registerRequest{
		Email:    email,
		Password: password,
		Name:     name,
	}

	jsonData, err := json.Marshal(registerData)
	if err != nil {
		return fmt.Errorf("failed to marshal register data: %w", err)
	}

	resp, err := http.Post(
		r.baseURL+"/users/register",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return fmt.Errorf("failed to send register request: %w", err)
	}
	defer resp.Body.Close()

	var result registerResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("registration failed: %s", result.Error)
	}

	return nil
}
