package main

import (
	"fmt"
)

func main() {
	fileURL := "https://go.dev/dl/go1.22.3.src.tar.gz"
	filepath := "go1.22.3.src.tar.gz"
	jsonURL := "https://go.dev/dl/?mode=json"
	if err := downloadfile(fileURL, filepath); err != nil {
		panic(err)
	}
	println("File downloaded successfully")
	expectedHash, err := getSha256Hash(jsonURL, filepath)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Expected SHA256 hash: %s\n", expectedHash)
	calculatedHash, err := calculateFileSha256(filepath)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Calculated SHA256 hash: %s\n", calculatedHash)
	if calculatedHash == expectedHash {
		println("The hashes match! True")
	} else {
		println("The hashes do not match! False")
	}
}
