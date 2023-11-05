package auth

import (
	"WebBlog/internal/model"
	"gopkg.in/yaml.v3"
	"os"
)

func GetCreds(path string) (*model.Creds, error) {
	var creds model.Creds

	credsFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(credsFile, &creds)
	if err != nil {
		return nil, err
	}

	return &creds, nil
}
