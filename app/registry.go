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

func download_image_layer(image_name string, tag_name string
						jailed_command_path string) error {
	
}