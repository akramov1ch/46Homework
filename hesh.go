package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Go struct {
	Version string `json:"version"`
	Files   []struct {
		Filename string `json:"filename"`
		Sha256   string `json:"sha256"`
	} `json:"files"`
}

func GetSha256Hash(url string, filename string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var releases []Go
	if err := json.NewDecoder(resp.Body).Decode(&releases); err != nil {
		return "", err
	}
	for _, release := range releases {
		for _, file := range release.Files {
			if file.Filename == filename {
				return file.Sha256, nil
			}
		}
	}
	return "", fmt.Errorf("file not found")
}

func CalculateFileSha256(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash), nil
}
