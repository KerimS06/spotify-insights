package artistApi

import (
	"encoding/json"
	"fmt"

	"github.com/We-Code-at-Nights/spotify-insights/src/collection"
	"github.com/We-Code-at-Nights/spotify-insights/src/spotify/util"
	"github.com/go-resty/resty/v2"
)

func FindByID(id string) (collection.Artist, error) {
	token, err := util.GetAccessToken()
	if err != nil {
		fmt.Println(err)
		return collection.Artist{}, err
	}

	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+token.AccessToken).
		SetPathParam("id", id).Get(util.ArtistEndpoint)
	if err != nil {
		return collection.Artist{}, err
	} else if resp.IsError() {
		return collection.Artist{}, fmt.Errorf("artist get request failed")
	}

	var response ArtistResponse
	if err := json.Unmarshal([]byte(resp.Body()), &response); err != nil {
		return collection.Artist{}, err
	}

	a := collection.Artist{
		ID:         response.ID,
		Name:       response.Name,
		Popularity: response.Popularity,
		Genres:     response.Genres,
		Followers:  response.Followers.Total,
	}

	return a, nil

}
