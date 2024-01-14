package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"
)

// stackoverflow.com/questions/30681054/what-is-the-usage-of-backtick-in-golang-structs-definition
type AuthResponse struct {
	Token string `json:"token"`
}

type ImageLayer struct {
	Size int `json:"size"`
	Digest string `json:"digest"`
	MediaType string `json:"mediaType"`
}

func authenticate_image(image_name string) (*AuthResponse, error) {
	auth_url_query := url.PathEscape(
		fmt.Sprintf("service=registry.docker.io&scope=repository:library/%s:pull", image_name),
	) // format string with sprintf and escape

	res, err := http.Get(
		fmt.Sprintf("https://auth.docker.io/token?%s", auth_url_query)
	)
	if err != nil { return nil, err }

	var auth AuthResponse

	err = json.NewDecoder(res.Body).Decode(&auth)
	if err != nil { return nil, err }

	return &auth, nil
}

func download_layers(image_name string, layers []ImageLayer, tmpDir string, auth *AuthResponse) ([]string, error) {

	registry_address := "http://localhost:5000/v2/library/blobs/%s/%s"
	url := fmt.Sprintf(registry_address, image_name, layer.Digest)

	var layer_files []string

	for _, layer := range layers {
		req, _ := http.NewRequest("GET", url, nil)

		if auth != nil {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", auth.Token))
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil { return nil, err }

		if res.StatusCode != 200 {
			return nil, fmt.Errorf("Layer download failed with status code: %d", res.StatusCode)
		}
		defer res.Body.Close()

		content, err := ioutil.ReadAll(res.Body)
		if err != nil { return nil, err }
	}

	return layer_files, nil
}

func download_image(image_name string, tag_name string
						jailed_command_path string) error {
	layers_dir, err := os.MkdirTemp(".", "layers")
	if err != nil { return err }
}