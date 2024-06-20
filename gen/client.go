package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	baseURL string
}

func NewClient(baseURL string) *Client {
	return &Client{baseURL: baseURL}
}

type ConfirmPassword struct {
	Password string `json:"password"`
}

type Email struct {
	Email string `json:"email"`
}

type Password struct {
	Password string `json:"password"`
}

type UserLoginReq struct {
	Email    Email    `json:"email"`
	Password Password `json:"password"`
}

type UserSignUpReq struct {
	Email           Email           `json:"email"`
	Password        Password        `json:"password"`
	ConfirmPassword ConfirmPassword `json:"confirmPassword"`
}

type UsersSignupResponse struct {
	Status string `json:"status"`
}

func (c *Client) POST_UsersSignup(req UserSignUpReq) (UsersSignupResponse, error) {
	var resp UsersSignupResponse

	url := fmt.Sprintf("%s/users/signup", c.baseURL)
	body, err := json.Marshal(req)
	if err != nil {
		return resp, err
	}

	httpResp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return resp, err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return resp, fmt.Errorf("unexpected status: %v", httpResp.Status)
	}

	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

type UsersLoginResponse struct {
	Status string `json:"status"`
}

func (c *Client) POST_UsersLogin(req UserLoginReq) (UsersLoginResponse, error) {
	var resp UsersLoginResponse

	url := fmt.Sprintf("%s/users/login", c.baseURL)
	body, err := json.Marshal(req)
	if err != nil {
		return resp, err
	}

	httpResp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return resp, err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return resp, fmt.Errorf("unexpected status: %v", httpResp.Status)
	}

	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
