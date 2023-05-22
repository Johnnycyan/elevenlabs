package client

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/getcohesive/elevenlabs/client/types"
)

func (c Client) TTSWriter(ctx context.Context, w io.Writer, text, modelID, voiceID string, options types.SynthesisOptions) error {
	options.Clamp()
	url := fmt.Sprintf(c.endpoint+"/v1/text-to-speech/%s", voiceID)
	opts := types.TTS{
		Text:          text,
		ModelID:       modelID,
		VoiceSettings: options,
	}
	b, _ := json.Marshal(opts)
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	req.Header.Set("xi-api-key", c.apiKey)
	req.Header.Set("User-Agent", "github.com/getcohesive/elevenlabs")
	req.Header.Set("accept", "audio/mpeg")
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	switch res.StatusCode {
	case 401:
		return ErrUnauthorized
	case 200:
		if err != nil {
			return err
		}
		defer res.Body.Close()
		_, err = io.Copy(w, res.Body)
		return nil
	case 422:
		fallthrough
	default:
		ve := types.ValidationError{}
		defer res.Body.Close()
		_ = json.NewDecoder(res.Body).Decode(&ve)
		return ve
	}
}

func (c Client) TTS(ctx context.Context, text, voiceID, modelID string, options types.SynthesisOptions) ([]byte, error) {
	options.Clamp()
	url := fmt.Sprintf(c.endpoint+"/v1/text-to-speech/%s", voiceID)
	client := &http.Client{}
	opts := types.TTS{
		Text:          text,
		ModelID:       modelID,
		VoiceSettings: options,
	}
	b, _ := json.Marshal(opts)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("xi-api-key", c.apiKey)
	req.Header.Set("User-Agent", "github.com/getcohesive/elevenlabs")
	req.Header.Set("accept", "audio/mpeg")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 401:
		return nil, ErrUnauthorized
	case 200:
		if err != nil {
			return nil, err
		}
		b := bytes.Buffer{}
		w := bufio.NewWriter(&b)

		defer res.Body.Close()
		_, err = io.Copy(w, res.Body)
		if err != nil {
			return nil, err
		}
		return b.Bytes(), nil
	default:
		ve := types.ValidationError{}
		defer res.Body.Close()
		_ = json.NewDecoder(res.Body).Decode(&ve)
		return nil, ve
	}
}

func (c Client) TTSStream(ctx context.Context, w io.Writer, text, modelID, voiceID string, options types.SynthesisOptions) error {
	options.Clamp()
	url := fmt.Sprintf(c.endpoint+"/v1/text-to-speech/%s/stream", voiceID)
	opts := types.TTS{
		Text:          text,
		ModelID:       modelID,
		VoiceSettings: options,
	}
	b, _ := json.Marshal(opts)
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	req.Header.Set("xi-api-key", c.apiKey)
	req.Header.Set("User-Agent", "github.com/getcohesive/elevenlabs")
	req.Header.Set("accept", "audio/mpeg")
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	switch res.StatusCode {
	case 401:
		return ErrUnauthorized
	case 200:
		if err != nil {
			return err
		}
		defer res.Body.Close()
		_, err = io.Copy(w, res.Body)
		return nil
	default:
		ve := types.ValidationError{}
		defer res.Body.Close()
		_ = json.NewDecoder(res.Body).Decode(&ve)
		return ve
	}
}
