package client

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Johnnycyan/elevenlabs/client/types"
)

func (c Client) GetUserInfo(ctx context.Context) (types.UserResponseModel, error) {
	url := c.endpoint + "/v1/user"
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return types.UserResponseModel{}, err
	}
	req.Header.Set("xi-api-key", c.apiKey)
	req.Header.Set("User-Agent", "github.com/Johnnycyan/elevenlabs")
	req.Header.Set("accept", "application/json")
	res, err := client.Do(req)

	switch res.StatusCode {
	case 401:
		return types.UserResponseModel{}, ErrUnauthorized
	case 200:
		if err != nil {
			return types.UserResponseModel{}, err
		}

		var user types.UserResponseModel
		defer res.Body.Close()
		jerr := json.NewDecoder(res.Body).Decode(&user)
		if jerr != nil {
			return types.UserResponseModel{}, jerr
		}
		return user, err
	case 422:
		fallthrough
	default:
		ve := types.ValidationError{}
		defer res.Body.Close()
		_ = json.NewDecoder(res.Body).Decode(&ve)
		return types.UserResponseModel{}, ve
	}
}

func (c Client) GetSubscriptionInfo(ctx context.Context) (types.Subscription, error) {
	info, err := c.GetUserInfo(ctx)
	return info.Subscription, err
}
