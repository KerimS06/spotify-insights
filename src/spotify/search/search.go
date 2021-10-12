package search

import (
	"encoding/json"
	"fmt"

	"github.com/We-Code-at-Nights/spotify-insights/src/collection"
	"github.com/We-Code-at-Nights/spotify-insights/src/spotify/util"
	"github.com/go-resty/resty/v2"
)

func TrackByNameAndArtist(trackName, artistName string) (collection.Track, error) {
	token, err := util.GetAccessToken()
	if err != nil {
		fmt.Println(err)
		return collection.Track{}, err
	}

	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+token.AccessToken).
		SetQueryParams(map[string]string{
			"limit": "1",
			"q":     fmt.Sprintf(`track:"%s" artist:"%s"`, trackName, artistName),
			"type":  "track",
		}).Get(util.SearchEndpoint)
	if err != nil {
		return collection.Track{}, err
	} else if resp.IsError() {
		return collection.Track{}, fmt.Errorf("track search request failed")
	}

	var response TrackResponse
	if err := json.Unmarshal([]byte(resp.Body()), &response); err != nil {
		return collection.Track{}, err
	}

	if len(response.Tracks.Items) == 0 {
		return collection.Track{}, fmt.Errorf("track search request empty result")
	}

	item := response.Tracks.Items[0]
	artists := make([]collection.Artist, 0)

	for _, artist := range item.Artists {
		a := collection.Artist{
			ID:   artist.ID,
			Name: artist.Name,
		}

		artists = append(artists, a)
	}

	t := collection.Track{
		ID:         item.ID,
		Name:       item.Name,
		Popularity: item.Popularity,
		Album: collection.Album{
			ID:          item.Album.ID,
			Name:        item.Album.Name,
			ReleaseDate: item.Album.ReleaseDate,
		},
		Artists: artists,
	}

	return t, nil
}
