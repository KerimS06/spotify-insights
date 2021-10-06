package util

import (
	"encoding/json"

	"github.com/We-Code-at-Nights/spotify-insights/src/config"
	"github.com/go-resty/resty/v2"
)

func GetAccessToken() (TokenResponse, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", "Basic "+config.TokenEndpoint).
		SetFormData(map[string]string{
			"grant_type": "client_credentials",
		}).Post(config.TokenEndpoint)
	if err != nil {
		return TokenResponse{}, err
	}

	var response TokenResponse
	if err := json.Unmarshal([]byte(resp.Body()), &response); err != nil {
		return TokenResponse{}, nil
	}

	return response, nil
}
