package util

import (
	"encoding/json"
	"fmt"

	"github.com/We-Code-at-Nights/spotify-insights/src/config"
	"github.com/go-resty/resty/v2"
)

func GetAccessToken() (TokenResponse, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", "Basic "+config.EncodedCredentials).
		SetFormData(map[string]string{
			"grant_type": "client_credentials",
		}).Post(TokenEndpoint)
	if err != nil {
		return TokenResponse{}, err
	} else if resp.IsError() {
		return TokenResponse{}, fmt.Errorf("token request failed")
	}

	var response TokenResponse
	if err := json.Unmarshal([]byte(resp.Body()), &response); err != nil {
		return TokenResponse{}, err
	}

	return response, nil
}
